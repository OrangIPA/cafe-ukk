package controller

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/OrangIPA/ukekehfrozekakhyr/db"
	"github.com/OrangIPA/ukekehfrozekakhyr/helper"
	"github.com/OrangIPA/ukekehfrozekakhyr/model"
	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

type CreateMenuParams struct {
	NamaMenu  string `form:"namaMenu"`
	Jenis     string `form:"jenis"`
	Deskripsi string `form:"deskripsi"`
	Harga     int    `form:"harga"`
}

type UpdateMenuParams struct {
	MenuId uint	`form:"menuId"`
	NamaMenu string `form:"namaMenu"`
	Jenis	string `form:"jenis"`
	Deskripsi string `form:"deskripsi"`
	Harga int	`form:"harga"`
}

func CreateMenu(c *fiber.Ctx) error {
	// Get token claims
	claims := helper.TokenClaims(c)
	role := claims["role"].(string)

	// Get gambar file
	gambar, err := c.FormFile("gambar")
	if err != nil {
		return err
	}

	// Preventing duplicate file name
	title := gambar.Filename
	var menus []model.Menu
	for {
		db.DB.Find(&menus, "gambar = ?", title)
		if len(menus) == 0 { break }
		t := strings.Split(title, ".")
		title = t[0] + "0." + t[1]
	}

	// Return if role is not admin
	if role != "admin" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Parse body
	m := new(CreateMenuParams)
	if err := c.BodyParser(m); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Return if jenis is neither makanan or minuman
	if m.Jenis != "makanan" && m.Jenis != "minuman" {
		return c.Status(fiber.StatusBadRequest).SendString("jenis is invalid")
	}

	// Return if the params is empty
	if m.NamaMenu == "" || m.Harga == 0 {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Create entry model
	newMenu := model.Menu{NamaMenu: m.NamaMenu, Jenis: m.Jenis, Deskripsi: m.Deskripsi, Gambar: title, Harga: m.Harga}

	// Write gambar go file system
	err = c.SaveFile(gambar, fmt.Sprintf("./public/gambarmenu/%s", title))
	if err != nil {
		return err
	}

	// Crete the entry and return the error if any
	err = db.DB.Create(&newMenu).Error
	if err != nil {
		return err
	}

	// Return 201 Created
	return c.SendStatus(fiber.StatusCreated)
}

func GetAllMenu(c *fiber.Ctx) error {
	// Query to database
	var menus []model.Menu
	if err := db.DB.Find(&menus).Error; err != nil {
		return err
	}

	// Return the menus
	return c.JSON(menus)
}

func GetMenuById(c *fiber.Ctx) error {
	// Query to database
	var menu model.Menu
	if err := db.DB.First(&menu, c.Params("id")).Error; err != nil {
		return err
	}
	
	// Return the menu
	return c.JSON(menu)
}

func UpdateMenu(c *fiber.Ctx) error {
	// Get token claims
	claims := helper.TokenClaims(c)
	role := claims["role"].(string)

	// Return if role is not admin
	if role != "admin" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Get gammbar if any, if there isn't, set isGambarChanged to false
	isGambarChanged := true
	gambar, err := c.FormFile("gambar")
	if err != nil {
		isGambarChanged = false
	}

	// Prevent duplicate file name
	var title string
	if isGambarChanged {
		title = gambar.Filename
		var menus []model.Menu
		for {
			db.DB.Find(&menus, "gambar = ?", title)
			if len(menus) == 0 { break }
			t := strings.Split(title, ".")
			title = t[0] + "0." + t[1]
		}
	}

	// Parse body
	menu := new(UpdateMenuParams)
	if err := c.BodyParser(menu); err != nil {
		return err
	}

	// Return if jenis is neither makanan or minuman
	if menu.Jenis != "makanan" && menu.Jenis != "minuman" {
		return c.Status(fiber.StatusBadRequest).SendString("jenis is invalid")
	}

	// Return if any of the params is empty
	if menu.MenuId == 0 || menu.NamaMenu == "" || menu.Harga == 0 {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if !isGambarChanged {
		// Create menu model
		updatedMenu := model.Menu{MenuID: menu.MenuId, NamaMenu: menu.NamaMenu, Jenis: menu.Jenis, Deskripsi: menu.Deskripsi, Harga: menu.Harga}
		
		// Query to database and handle the error
		var mysqlErr *mysql.MySQLError
		if err := db.DB.Model(&model.Menu{MenuID: menu.MenuId}).Omit("gambar").Updates(updatedMenu).Error; err != nil {
			if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
				return c.Status(fiber.StatusBadRequest).SendString("nothing is changed")
			}
			return err
		}
		return c.SendStatus(fiber.StatusOK)
	}

	// Write gambar to file system
	if isGambarChanged {
		if err := c.SaveFile(gambar, fmt.Sprintf("./public/gambarmenu/%s", title)); err != nil {
			return err
		}
	}

	// Create entry model
	updatedMenu := model.Menu{MenuID: menu.MenuId, NamaMenu: menu.NamaMenu, Jenis: menu.Jenis, Deskripsi: menu.Deskripsi, Gambar: title, Harga: menu.Harga}

	// Update the menu and return error if any
	if err := db.DB.Save(&updatedMenu).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return c.Status(fiber.StatusBadRequest).SendString("nothing is changed")
		}
		return err
	}
	return c.SendStatus(fiber.StatusOK)
}

func DeleteMenu(c *fiber.Ctx) error {
	// Get token claims
	claims := helper.TokenClaims(c)
	role := claims["role"].(string)

	// Return if role is not admin
	if role != "admin" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Get userId, parse it, and then return the error if any
	menuId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Check if menu is exist. if not, return 404
	var menu model.Menu
	if err := db.DB.First(&menu, menuId).Error; err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	// Query to database and return the error if any
	if err := db.DB.Delete(&model.Menu{}, menuId).Error; err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
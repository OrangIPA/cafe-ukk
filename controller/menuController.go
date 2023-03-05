package controller

import (
	"fmt"
	"strings"

	"github.com/OrangIPA/ukekehfrozekakhyr/db"
	"github.com/OrangIPA/ukekehfrozekakhyr/helper"
	"github.com/OrangIPA/ukekehfrozekakhyr/model"
	"github.com/gofiber/fiber/v2"
)

type CreateMenuParams struct {
	NamaMenu  string `form:"namaMenu"`
	Jenis     string `form:"jenis"`
	Deskripsi string `form:"deskripsi"`
	Harga     int    `form:"harga"`
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
		return c.SendStatus(fiber.StatusBadRequest)
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

	// Crete the entry and Return the error if any
	err = db.DB.Create(&newMenu).Error
	if err != nil {
		return err
	}

	// Return 201 Created
	return c.SendStatus(fiber.StatusCreated)
}

package controller

import (
	"log"
	"strconv"
	"time"

	"github.com/OrangIPA/ukekehfrozekakhyr/db"
	"github.com/OrangIPA/ukekehfrozekakhyr/helper"
	"github.com/OrangIPA/ukekehfrozekakhyr/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type transaksiParams struct {
	MejaID          uint                    `form:"mejaId" json:"mejaId"`
	NamaPelanggan   string                  `form:"namaPelanggan" json:"namaPelanggan"`
	Status          string                  `form:"status" json:"status"`
	DetailTransaksi []detailTransaksiParams `form:"detailTranskasi" json:"detailTransaksi"`
}

type detailTransaksiParams struct {
	MenuID uint `json:"menuId" form:"menuId"`
	Jumlah int  `json:"jumlah" form:"jumlah"`
}

func CreateTransaksi(c *fiber.Ctx) error {
	// Only accept json as content type
	c.Accepts("application/json")

	// Get token claims
	claims := helper.TokenClaims(c)
	role := claims["role"].(string)
	userId := uint(claims["userId"].(float64))

	// Return if insufficient role
	if role != "kasir" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Parse body
	transaksi := new(transaksiParams)
	if err := c.BodyParser(transaksi); err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Initialize database transaction
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		// Create transaksi model
		t := model.Transaksi{
			TglTransaksi:  time.Now(),
			UserID:        userId,
			MejaID:        transaksi.MejaID,
			NamaPelanggan: transaksi.NamaPelanggan,
			Status:        transaksi.Status,
		}
		// Query transaksi to database and get the primary key
		if err := tx.Create(&t).Error; err != nil {
			return err
		}

		// Create detailTransaksi slice and fill it with all the detailTransaksi from the request
		var dT []model.DetailTransaksi
		for _, detil := range transaksi.DetailTransaksi {
			dT = append(dT, model.DetailTransaksi{
				TransaksiID: t.TransaksiID,
				MenuID:      detil.MenuID,
				Jumlah:      detil.Jumlah,
			})
		}

		// Query detailTransaksi to database and both return the error
		if err := tx.Create(&dT).Error; err != nil {
			return nil
		}

		// Transaction completed succesfully, return nil committing the transaction
		return nil
	})
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

func UpdateTransaksiStatus(c *fiber.Ctx) error {
	// Get token claims
	claims := helper.TokenClaims(c)
	role := claims["role"].(string)

	// Return if insufficient role
	if role != "kasir" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Get request info
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	userId := uint(id)
	status := c.FormValue("status")

	// Query to database and return error if any
	if err := db.DB.Model(&model.Transaksi{TransaksiID: userId}).Update("status", status).Error; err != nil {
		return err
	}

	// Return OK
	return c.SendStatus(fiber.StatusOK)
}

func GetAllTransaksi(c *fiber.Ctx) error {
	// If query parameter for filtering not empty, pass task to GetAllTransaksiFilter
	if c.Query("nama_karyawan") != "" || c.Query("from") != "" || c.Query("to") != "" {
		return getTransaksiFilter(c)
	}

	// Query all transaksi from database
	var transaksis []model.Transaksi
	if err := db.DB.Find(&transaksis).Error; err != nil {
		return err
	}

	// For each transaksi, query transaction detail from database then attach it to transaksi
	for i, transaksi := range transaksis {
		var dT []model.DetailTransaksi
		if err := db.DB.Find(&dT, "transaksi_id = ?", transaksi.TransaksiID).Error; err != nil {
			return err
		}
		transaksis[i].DetailTransaksi = dT
	}

	// Return all transaksi
	return c.JSON(transaksis)
}

func GetTransaksiById(c *fiber.Ctx) error {
	// Get id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Query transaksi based on the id
	var transaksi model.Transaksi
	if err := db.DB.First(&transaksi, model.Transaksi{TransaksiID: uint(id)}).Error; err != nil {
		return err
	}

	// Query detail transaksi based on transaksi id then attach it to transaksi
	var dT []model.DetailTransaksi
	if err := db.DB.Find(&dT, "transaksi_id = ?", transaksi.TransaksiID).Error; err != nil {
		return err
	}
	transaksi.DetailTransaksi = dT

	// Return the transaksi
	return c.JSON(transaksi)
}

func getTransaksiFilter(c *fiber.Ctx) error {
	// Get query parameter
	namaKaryawan := c.Query("nama_karyawan")
	from, until := c.Query("dari"), c.Query("sampai")

	// Parse tanggal
	fromParsed, err := time.Parse("2-1-2006", from)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	untilParsed, err := time.Parse("2-1-2006", until)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	from = fromParsed.Format("2006-1-2")
	until = untilParsed.Format("2006-1-2")

	// Query transaksi
	var t []model.Transaksi
	namaKaryawan = "%" + namaKaryawan + "%"
	err = db.DB.Joins("JOIN user ON user.user_id = transaksi.user_id").Where("user.nama_user LIKE ? AND transaksi.tgl_transaksi >= ? AND transaksi.tgl_transaksi <= ?", namaKaryawan, from, until).Find(&t).Error
	if err != nil {
		return err
	}

	// Return the result
	return c.JSON(t)
}

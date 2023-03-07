package controller

import (
	"log"
	"time"

	"github.com/OrangIPA/ukekehfrozekakhyr/db"
	"github.com/OrangIPA/ukekehfrozekakhyr/helper"
	"github.com/OrangIPA/ukekehfrozekakhyr/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type transaksiParams struct {
	MejaID	uint `form:"mejaId" json:"mejaId"`
	NamaPelanggan string `form:"namaPelanggan" json:"namaPelanggan"`
	Status	string `form:"status" json:"status"`
	DetailTransaksi []detailTransaksiParams `form:"detailTranskasi" json:"detailTransaksi"`
}

type detailTransaksiParams struct {
	MenuID uint `json:"menuId" form:"menuId"`
	Jumlah int `json:"jumlah" form:"jumlah"`
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
			TglTransaksi: time.Now(),
			UserID: userId,
			MejaID: transaksi.MejaID,
			NamaPelanggan: transaksi.NamaPelanggan,
			Status: transaksi.Status,
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
				MenuID: detil.MenuID,
				Jumlah: detil.Jumlah,
			})
		}

		// Query detailTransaksi to database and both return the error 
		if err := tx.Create(&dT).Error; err != nil {
			return nil
		}

		return nil 
	})
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}
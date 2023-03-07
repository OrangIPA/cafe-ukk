package model

type DetailTransaksi struct {
	DetailTransaksiID uint `json:"detailTransasksiId" gorm:"primaryKey; autoIncrement"`
	TransaksiID       uint `json:"transaksiId" gorm:"not null"`
	MenuID            uint `json:"menuId" gorm:"not null"`
	Jumlah            int  `json:"jumlah" gorm:"not null"`
}

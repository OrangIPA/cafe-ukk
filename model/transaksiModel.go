package model

import "time"

type Transaksi struct {
	TransaksiID     uint      `json:"transaksiId" gorm:"primaryKey; autoIncrement"`
	TglTransaksi    time.Time `json:"tglTransaksi" gorm:"not null"`
	UserID          uint      `json:"userId" gorm:"not null"`
	MejaID          uint      `json:"mejaId" gorm:"not null"`
	NamaPelanggan   string    `json:"namaPelanggan"`
	Status          string    `json:"status" gorm:"type:enum('belum_bayar', 'lunas');not null"`
	DetailTransaksi []DetailTransaksi
}

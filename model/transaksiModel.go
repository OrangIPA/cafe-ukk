package model

import "time"

type Transaksi struct {
	ID int `gorm:"column:id; primaryKey; autoIncrement"`
	TglTransaksi time.Time `gorm:"column:tgl_transaksi"`
}
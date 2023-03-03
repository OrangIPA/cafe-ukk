package model

import "time"

type Transaksi struct {
	TransaksiID   int       `gorm:"primaryKey; autoIncrement"`
	TglTransaksi  time.Time `gorm:"not null"`
	User          User      `gorm:"not null; foreignKey: UserID"`
	Meja          Meja      `gorm:"foreignKey: MejaID"`
	NamaPelanggan string
	Status        StatusTransaksi `sql:"type:ENUM('BELUM_BAYAR','LUNAS')" gorm:"not null"`
}

type StatusTransaksi string

const (
	BELUM_BAYAR StatusTransaksi = "BELUM_BAYAR"
	LUNAS       StatusTransaksi = "LUNAS"
)

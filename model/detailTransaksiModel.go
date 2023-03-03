package model

type DetailTransaksi struct {
	DetailTransaksiID int       `gorm:"primaryKey"`
	Transaksi         Transaksi `gorm:"foreignKey: TransaksiID; not null"`
	Menu              Menu      `gorm:"foreignKey: MenuID; not null"`
	Harga             int       `gorm:"not null"`
}

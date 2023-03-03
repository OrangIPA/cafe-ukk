package model

type Menu struct {
	MenuID int `gorm:"primaryKey; autoIncrement"`
	NamaMenu string `gorm:"not null"`
	Jenis JenisPesanan `sql:"type:ENUM('MAKANAN','MINUMAN')" gorm:"not null"`
	Deskripsi string
	Gambar string
	Harga int `gorm:"not null"`
}

type JenisPesanan string

const (
	MAKANAN JenisPesanan = "MAKANAN"
	MINUMAN JenisPesanan = "MINUMAN"
)
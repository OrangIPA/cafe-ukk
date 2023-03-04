package model

type Menu struct {
	MenuID          uint   `json:"menuId" gorm:"primaryKey; autoIncrement"`
	NamaMenu        string `json:"namaMenu" gorm:"not null"`
	Jenis           string `json:"jenis" gorm:"type:enum('makanan', 'minuman'); not null"`
	Deskripsi       string `json:"deskripsi"`
	Gambar          string `json:"gambar"`
	Harga           int    `json:"harga" gorm:"not null"`
	DetailTransaksi []DetailTransaksi
}
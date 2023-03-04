package model

type Meja struct {
	MejaID    uint   `json:"mejaId" gorm:"primaryKey; autoIncrement"`
	NomorMeja string `json:"nomorMeja" gorm:"not null"`
	Transaksi []Transaksi
}

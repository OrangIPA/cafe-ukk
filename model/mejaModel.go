package model

type Meja struct {
	MejaID    uint   `json:"mejaId" gorm:"primaryKey; autoIncrement"`
	NomorMeja string `json:"nomorMeja" gorm:"not null"`
	Status    string `json:"status" gorm:"type:enum('kosong', 'terisi'); default: kosong; not null"`
	Transaksi []Transaksi
}

package model

type Meja struct {
	MejaID int `gorm:"primaryKey; autoIncrement"`
	NomorMeja string
}
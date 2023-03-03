package model

type User struct {
	ID int `gorm:"primaryKey; autoIncrement"`
	NamaUser string `gorm:"not null"`
	Role userPriv `sql:"type:ENUM('ADMIN', 'KASIR', 'MANAJER')" gorm:"not null"`
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
}

type userPriv string

const (
	ADMIN userPriv = "ADMIN"
	KASIR userPriv = "KASIR"
	MANAJER userPriv = "MANAJER"
)
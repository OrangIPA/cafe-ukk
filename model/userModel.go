package model

type User struct {
	ID int `gorm:"column:id; primaryKey; autoIncrement"`
	NamaUser string `gorm:"column:nama_user; not null"`
	Role userPriv `sql:"type:ENUM('ADMIN', 'KASIR', 'MANAJER')" gorm:"column:role; not null"`
	Username string `gorm:"column:username; not null"`
	Password string `gorm:"column:password; not null"`
}

type userPriv string

const (
	ADMIN userPriv = "ADMIN"
	KASIR userPriv = "KASIR"
	MANAJER userPriv = "MANAJER"
)
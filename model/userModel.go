package model

type User struct {
	UserID    uint   `json:"userId" gorm:"primaryKey; autoIncrement"`
	NamaUser  string `json:"namaUser" gorm:"not null"`
	Role      string `json:"role" gorm:"type:enum('admin', 'manager', 'kasir'); default:'kasir'; not null"`
	Username  string `json:"username" gorm:"not null"`
	Password  string `json:"password" gorm:"not null"`
	Transaksi []Transaksi
}

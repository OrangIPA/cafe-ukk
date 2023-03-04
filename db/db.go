package db

import (
	"fmt"
	"os"

	"github.com/OrangIPA/ukekehfrozekakhyr/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("Failed to connect to database")
	}
}

func SyncDB() {
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Transaksi{})
	DB.AutoMigrate(&model.Meja{})
	DB.AutoMigrate(&model.Menu{})
	DB.AutoMigrate(&model.DetailTransaksi{})
}

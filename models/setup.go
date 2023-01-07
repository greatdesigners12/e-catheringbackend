package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/e-cathering?parseTime=true"))
	if err != nil {
		fmt.Println("Gagal koneksi database")
	}



	db.AutoMigrate(&User{})
	
	db.AutoMigrate(&Cathering{})
	db.AutoMigrate(&Category{})
	db.AutoMigrate(&Cart{})
	db.AutoMigrate(&UserInformation{})
	

	DB = db
}

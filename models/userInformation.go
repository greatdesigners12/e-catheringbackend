package models

import (


)

type UserInformation struct {
	Id          	int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	User_id     	int64 `gorm:"int(11)" json:"user_id"`
	Nama_lengkap    string `gorm:"text" json:"nama_lengkap"`
	Alamat          string `gorm:"text" json:"alamat"`
	Image_profile   string `gorm:"text" json:"image_profile"`
	Is_verified   int64 `gorm:"int(1)" json:"is_verified"`
}
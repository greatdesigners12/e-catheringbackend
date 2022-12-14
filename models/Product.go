package models

import "time"

type Product struct {
	Id           int64     `gorm:"primaryKey" json:"id"`
	Cathering_id string    `gorm:"type:int(11)" json:"cathering_id"`
	Nama         string    `gorm:"type:varchar(255)" json:"nama"`
	Deskripsi    string    `gorm:"type:text" json:"deskripsi"`
	Harga        int64     `gorm:"type:bigint(20)" json:"harga"`
	Image_Url    string    `gorm:"type:text" json:"image_url"`
	Cathering    Cathering `gorm:"foreignKey:Cathering_id"`
	Type         string    `gorm:"type:varchar(100)" json:"type"`
	Start_date   time.Time `gorm:"type:date" json:"start_date"`
	Due_date     time.Time `gorm:"type:date" json:"due_date"`
}

package models

type Category struct{
	Id int64 `gorm:"primaryKey" json:"id"`
	Nama_kategori string `gorm:"type:text" json:"nama_kategori"` 
	Image string `gorm:"type:text" json:"image"` 
}
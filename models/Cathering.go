package models

type Cathering struct{
	Id int64 `gorm:"primaryKey" json:"id"`
	User_id string `gorm:"type:int(11)" json:"user_id"`	
	Nama string `gorm:"type:varchar(255)" json:"nama"`
	Tanggal_register string `gorm:"type:date" json:"tanggal_register"`
	Deskripsi string `gorm:"type:text" json:"deskripsi"`
	Image_logo string `gorm:"type:text" json:"image_logo"`
	Image_menu string `gorm:"type:text" json:"image_menu"`
	Is_verified int64 `gorm:"type:tinyint(1)" json:"is_verified"`

}
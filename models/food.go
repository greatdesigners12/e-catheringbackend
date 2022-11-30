package models

type Food struct{
	Id int64 `gorm:"primaryKey" json:"id"`
	Nama string `gorm:"type:varchar(100)" json:"nama"`	
	Desc string `gorm:"type:varchar(300)" json:"desc"`
	Quantity int64 `gorm:"type:int(100)" json:"quantity"`
}
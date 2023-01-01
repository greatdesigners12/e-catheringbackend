package models

type Cart struct{
	Id int64 `gorm:"primaryKey" json:"id"`
	User_id int64 `gorm:"type:int(11)" json:"user_id"` 
	User User `gorm:"foreignKey:User_id"`
	Product_id int64 `gorm:"type:int(11)" json:"product_id"` 
	Product Product `gorm:"foreignKey:Product_id"`
	Cathering_id int64 `gorm:"type:int(11)" json:"cathering_id"`
	Cathering Cathering `gorm:"foreignKey:Cathering_id"`
}
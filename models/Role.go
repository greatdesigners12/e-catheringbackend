package models


type Role struct {
	Id          int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Role        string  `gorm:"varchar(10)" json:"role"`
}



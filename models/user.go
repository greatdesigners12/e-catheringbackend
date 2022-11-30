package models

type User struct {
	Id          int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Email string `gorm:"varchar(255)" json:"email"`
	Password    string `gorm:"text" json:"password"`
	Role    string `gorm:"varchar(100)" json:"role"`
}

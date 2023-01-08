package models


type Review struct {
	Id          int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	User_id        int  `gorm:"int(11)" json:"user_id"`
	Cathering_id        int  `gorm:"int(11)" json:"cathering_id"`
	Review_msg        string  `gorm:"text" json:"review_msg"`
	Rating        float64  `gorm:"float(10)" json:"rating"`
	User    User `gorm:"foreignKey:User_id"`
	Cathering Cathering `gorm:"foreignKey:Cathering_id"`
}
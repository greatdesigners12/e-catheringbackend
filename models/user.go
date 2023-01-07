package models
import "time"

type User struct {
	Id          int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Email string `gorm:"varchar(255)" json:"email"`
	Password    string `gorm:"text" json:"password"`
	Role_id    int `gorm:"int(11)" json:"role_id"`
	Role       Role `gorm:"foreignKey:role_id"`
	UserInformation UserInformation `gorm:"foreignKey:User_id"`
}

type UserInformation struct {
	Id          	int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	User_id     	int64 `gorm:"int(11)" json:"user_id"`
	Nama_lengkap    string `gorm:"text" json:"nama_lengkap"`
	Tanggal_lahir   time.Time `gorm:"datetime" json:"tanggal_lahir"`
	Alamat          string `gorm:"text" json:"alamat"`
	Image_profile   string `gorm:"text" json:"image_profile"`
	Is_verified   int64 `gorm:"int(1)" json:"is_verified"`
}

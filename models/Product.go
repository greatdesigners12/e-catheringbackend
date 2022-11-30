package models

type Product struct{
	Id int64 `gorm:"primaryKey" json:"id"`
	Cathering_id string `gorm:"type:int(11)" json:"cathering_id"`	
	Nama string `gorm:"type:varchar(255)" json:"nama"`
	Deskripsi string `gorm:"type:text" json:"deskripsi"`
	Harga int64 `gorm:"type:bigint(20)" json:"harga"`
	Stock_persediaan int64 `gorm:"type:bigint(20)" json:"stock_persediaan"`
}
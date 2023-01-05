package models

import "time"

type TransactionGroup struct {
	Id                       int64                    `gorm:"primaryKey" json:"id"`
	Id_transaction           string                    `gorm:"type:text" json:"id_transaction"`
	TransactionGroupRelation []TransactionGroupRelation `gorm:"foreignKey:Transaction_group_id"`
	TotalPrice               int64                    `gorm:"type:bigint(20)" json:"total_price"`
	DateTransaction          time.Time                `gorm:"type:datetime" json:"dateTransaction"`
	User_id                  int64                  `gorm:"type:int(11)" json:"user_id"`
	Cathering_id             int64                  `gorm:"type:int(11)" json:"cathering_id"`
	Cathering                Cathering               `gorm:"foreignKey:Cathering_id"`
	User                     User                     `gorm:"foreignKey:user_id"`
	Shipping_price           int64                    `gorm:"type:bigint(20)" json:"shipping_price"`
	Status                   string                 `gorm:"type:varchar(100)" json:"status"`  
	Snap_token               string                `gorm:"type:text" json:"snap_token"`
	Full_address             string                	`gorm:"type:text" json:"full_address"`
}

type TransactionGroupRelation struct {
	Id                       int64                    `gorm:"primaryKey" json:"id"`
	Transaction_group_id     int64                    `gorm:"type:int(11)" json:"transaction_group_id"`
	Transaction_product_id   int64                    `gorm:"type:int(11)" json:"transaction_product_id"`
	TransactionProduct       TransactionProduct       `gorm:"foreignKey:transaction_product_id"`
}

type TransactionProduct struct{
	Id       int64    `gorm:"primaryKey" json:"id"`
	Name     string   `gorm:"type:text" json:"name"`
	Price    int64    `gorm:"type:bigint(20)" json:"price"`
	Time     int   `gorm:"type:int(10)" json:"time"`
}
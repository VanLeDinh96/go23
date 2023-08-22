package entities

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	ProductID int `gorm:"type:varchar(50)"`
	Quantity int    `gorm:"type:varchar(50)"`
	Name     string `gorm:"type:varchar(50)"`
	Price    int    `gorm:"type:varchar(50)"`
}

func (Cart) TableName() string {
	return "carts"
}

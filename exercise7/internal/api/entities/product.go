package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name     string `gorm:"type:varchar(50)"`
	Price    float64    `gorm:"type:varchar(50)"`
	Quantity int    `gorm:"type:varchar(50)"`
	UserID   uint

	Creator   *User       `gorm:"foreignKey:CreatedById"`
}

func (Product) TableName() string {
	return "products"
}

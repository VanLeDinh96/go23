package entities

import (
	"gorm.io/gorm"
	"github.com/diegovanne/go23/exercise7/internal/constants/role_enums"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(50)"`
	Email    string `gorm:"type:varchar(50)"`
	Password string `gorm:"type:varchar(255)"`
	Role     role_enums.UserRole `gorm:"type:int"`
	Products []*Product `gorm:"foreignKey:CreatedById"`
}

func (User) TableName() string {
	return "users"
}

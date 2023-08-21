package database

import (
	"github.com/diegovanne/go23/exercise7/internal/api/entities"
  	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(connString string) {
	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN: connString,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(
		&entities.User{},
		&entities.Product{},
		&entities.Cart{},
	)

	if err != nil {
		panic("Failed to migrate database!")
	}

	DB = database
}

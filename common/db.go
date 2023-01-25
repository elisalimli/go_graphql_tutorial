package common

import (
	"github.com/elisalimli/meetmeup/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var tables []interface{} = []interface{}{
	models.User{},
	models.Todo{},
}

func InitDb() (*gorm.DB, error) {
	var err error
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/abc?sslmode=disable"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(tables...)

	return db, nil
}

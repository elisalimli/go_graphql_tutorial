package common

import (
	"log"
	"os"
	"time"

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
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	var err error
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/abc?sslmode=disable"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(tables...)

	return db, nil
}

package db

import (
	"log"

	"github.com/Lzrb0x/go-gorm-urlShortener-api/config"
	"github.com/Lzrb0x/go-gorm-urlShortener-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := config.Config.DbPath

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Database connection established")

	if config.Config.AutoMigrate {
		err = DB.AutoMigrate(
			&models.Url{},
		)

		if err != nil {
			log.Fatal("Failed to migrate database: ", err)
		}

		log.Println("Database migration completed")
	}

	log.Println("connect to db successfully!")
	return DB
}

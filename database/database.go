package database

import (
	"log"

	"simple-api-gorm/config"
	"simple-api-gorm/models"
)

func Migrate() {
	log.Printf("Migrate: Start")
	db := config.Db()
	db.AutoMigrate(
		&models.User{},
	)
	log.Printf("Migrate: Success")
}

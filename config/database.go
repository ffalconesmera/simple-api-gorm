package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Db() *gorm.DB {
	/*dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_TIMEZONE"),
	)*/

	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		"localhost",
		"5432",
		"test_gorm",
		"postgres",
		"root",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

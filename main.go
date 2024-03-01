package main

import (
	"simple-api-gorm/database"
	"simple-api-gorm/routes"
)

func main() {
	database.Migrate()
	// database.Seeder()
	routes.Run()
}

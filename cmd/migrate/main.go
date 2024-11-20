package main

import "github.com/haovoanh28/gai-webscraper/internal/infrastructure/database"

func main() {
	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	if err := database.Migrate(db); err != nil {
		panic(err)
	}

	if err := database.SeedDefaultData(db); err != nil {
		panic(err)
	}
}

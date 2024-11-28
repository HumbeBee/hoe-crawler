package main

import "github.com/HumbeBee/hoe-crawler/internal/infrastructure/database"

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

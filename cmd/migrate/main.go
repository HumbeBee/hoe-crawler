package main

import "github.com/haovoanh28/gai-webscraper/internal/infrastructure/database"

func main() {
	dbConfig := database.NewConfig()
	dbo, err := database.GetDB(dbConfig)
	if err != nil {
		panic(err)
	}

	if err := dbo.Migrate(); err != nil {
		panic(err)
	}
}

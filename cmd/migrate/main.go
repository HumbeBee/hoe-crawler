package main

import "github.com/haovoanh28/gai-webscraper/internal/infrastructure/database"

func main() {
	dbo, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	if err := dbo.Migrate(); err != nil {
		panic(err)
	}
}

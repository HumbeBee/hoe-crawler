package main

import (
	"github.com/haovoanh28/gai-webscraper/internal/db"
)

func main() {
	dbConfig := db.NewConfig()
	dbo, err := db.GetDB(dbConfig)
	if err != nil {
		panic(err)
	}

	if err := dbo.Migrate(); err != nil {
		panic(err)
	}
}

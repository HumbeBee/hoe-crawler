package testcmd

import "github.com/haovoanh28/gai-webscraper/internal/db"

func main() {
	dbConfig := db.NewConfig()
	db, err := db.GetDB(dbConfig)
	if err != nil {
		panic(err)
	}

}

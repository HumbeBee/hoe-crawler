package database

import (
	"fmt"

	"github.com/haovoanh28/gai-webscraper/internal/models"
)

func (dbo *DBO) InsertHoe(hoeInfo *models.HoeInfo) error {
	transaction := dbo.db.Begin()

	if err := transaction.Create(hoeInfo).Error; err != nil {
		transaction.Rollback()
		return err
	}

	if err := transaction.Commit().Error; err != nil {
		transaction.Rollback()
		return err
	}

	fmt.Println("Inserted")
	hoeInfo.Print()

	return nil
}

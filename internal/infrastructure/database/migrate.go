package database

import (
	"fmt"

	"github.com/haovoanh28/gai-webscraper/internal/models"
)

func (dbo *DBO) Migrate() error {
	err := dbo.db.AutoMigrate(&models.City{})
	if err != nil {
		return fmt.Errorf("failed to migrate City : %v", err)
	}

	// Create provinces table with city reference
	err = dbo.db.AutoMigrate(&models.Province{})
	if err != nil {
		return fmt.Errorf("failed to migrate Province : %v", err)
	}

	// Existing migrations
	err = dbo.db.AutoMigrate(&models.HoeInfo{})
	if err != nil {
		return fmt.Errorf("failed to migrate HoeInfo : %v", err)
	}

	err = dbo.db.AutoMigrate(&models.HoeReport{})
	if err != nil {
		return fmt.Errorf("failed to migrate HoeReport : %v", err)
	}

	return nil
}

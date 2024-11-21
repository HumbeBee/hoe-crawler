package database

import (
	"fmt"

	"github.com/haovoanh28/gai-webscraper/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	models := []interface{}{
		&models.Site{},
		&models.City{},
		&models.District{},
		&models.HoeInfo{},
		&models.HoeReport{},
		&models.WorkingHistory{},
	}

	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			return fmt.Errorf("failed to migrate %T: %v", model, err)
		}
	}

	return nil
}

package database

import (
	"fmt"

	"github.com/HumbeBee/hoe-crawler/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	_models := []interface{}{
		&models.Site{},
		&models.City{},
		&models.District{},
		&models.Street{},
		&models.HoeInfo{},
		&models.HoeReport{},
		&models.WorkingHistory{},
		&models.FailedURL{},
	}

	for _, model := range _models {
		if err := db.AutoMigrate(model); err != nil {
			return fmt.Errorf("failed to migrate %T: %v", model, err)
		}
	}

	return nil
}

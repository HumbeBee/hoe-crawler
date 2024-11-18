package database

import (
	"github.com/haovoanh28/gai-webscraper/internal/models"
	"gorm.io/gorm"
)

func SeedDefaultData(db *gorm.DB) error {
	if err := seedSites(db); err != nil {
		return err
	}

	return nil
}

func seedSites(db *gorm.DB) error {
	sites := []models.Site{
		{
			Name:      "gaito",
			IsExpired: false,
		},
		{
			Name:      "gaigu",
			IsExpired: false,
		},
	}

	for _, site := range sites {
		result := db.Where(models.Site{Name: site.Name}).
			FirstOrCreate(&site)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

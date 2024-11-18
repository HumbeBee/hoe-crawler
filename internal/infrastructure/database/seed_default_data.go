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
			BaseURL:   "https://gaito.love",
			IsExpired: false,
		},
		{
			Name:      "gaigu",
			BaseURL:   "https://gaigu31.tv",
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

func seedCities(db *gorm.DB) error {

	return nil
}

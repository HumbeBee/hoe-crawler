package database

import "github.com/haovoanh28/gai-webscraper/internal/models"

func (dbo *DBO) SeedDefaultData() error {
	if err := dbo.seedSites(); err != nil {
		return err
	}

	return nil
}

func (dbo *DBO) seedSites() error {
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
		result := dbo.db.Where(models.Site{Name: site.Name}).
			FirstOrCreate(&site)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

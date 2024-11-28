package scrapers

import (
	"github.com/HumbeBee/hoe-crawler/internal/definitions"
	"github.com/HumbeBee/hoe-crawler/internal/interfaces"
	"github.com/HumbeBee/hoe-crawler/internal/scrapers/gaito"
)

func CreateScraper(baseConfig definitions.ScraperConfig) interfaces.Scraper {
	switch baseConfig.SiteName {
	case "gaito":
		return gaito.NewScraper(baseConfig)
	case "gaigu":
		return gaito.NewScraper(baseConfig)
	default:
		return nil
	}
}

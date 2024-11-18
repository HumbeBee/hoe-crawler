package scrapers

import (
	"github.com/haovoanh28/gai-webscraper/internal/definitions"
	"github.com/haovoanh28/gai-webscraper/internal/interfaces"
	"github.com/haovoanh28/gai-webscraper/internal/scrapers/gaito"
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

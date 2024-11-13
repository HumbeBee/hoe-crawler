package scrapers

import (
	"github.com/haovoanh28/gai-webscraper/internal/definitions"
	"github.com/haovoanh28/gai-webscraper/internal/interfaces"
	"github.com/haovoanh28/gai-webscraper/internal/scrapers/gaito"
)

func CreateScraper(baseConfig definitions.ScraperConfig) interfaces.Scraper {
	switch baseConfig.Site {
	case definitions.Gaito:
		return gaito.NewScraper(baseConfig)
	case definitions.Gaigu:
		return gaito.NewScraper(baseConfig)
	default:
		return nil
	}
}

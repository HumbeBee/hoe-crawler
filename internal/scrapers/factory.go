package scrapers

import (
	"fmt"

	"github.com/haovoanh28/gai-webscraper/internal/definitions"
	"github.com/haovoanh28/gai-webscraper/internal/interfaces"
	"github.com/haovoanh28/gai-webscraper/internal/scrapers/gaito"
)

func CreateScraper(site string) (interfaces.Scraper, error) {

	baseConfigs := definitions.ScraperConfig{
		RequestsPerSecond: 1.0,
	}

	switch site {
	case "gaito":
		baseConfigs.BaseURL = "https://gaito.love"
		return gaito.NewScraper(baseConfigs), nil
	case "gaigu":
		baseConfigs.BaseURL = "https://gaigu31.tv"
		return gaito.NewScraper(baseConfigs), nil
	default:
		return nil, fmt.Errorf("unknown site: %s", site)
	}
}

package definitions

import (
	"github.com/haovoanh28/gai-webscraper/internal/interfaces"
	"github.com/haovoanh28/gai-webscraper/internal/repository"
	"github.com/haovoanh28/gai-webscraper/internal/utils/logutil"
)

type ScraperConfig struct {
	SiteID            uint
	SiteName          string
	BaseURL           string
	RequestsPerSecond float64
	Logger            *logutil.Logger
}

type HoeServiceConfig struct {
	HoeRepo            repository.HoeRepository
	WorkingHistoryRepo repository.WorkingHistoryRepository
	Logger             *logutil.Logger
	Scraper            interfaces.Scraper
	MapperService      MapperService
	ValidateService    ValidateService
}

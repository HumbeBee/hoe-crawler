package service

import (
	"fmt"
	"github.com/haovoanh28/gai-webscraper/internal/interfaces"
	"github.com/haovoanh28/gai-webscraper/internal/repository"
	"github.com/haovoanh28/gai-webscraper/internal/utils/logutil"
)

type HoeService struct {
	scraper interfaces.Scraper
	repo    repository.HoeRepository
	logger  *logutil.Logger
}

func NewHoeService(repo repository.HoeRepository, logger *logutil.Logger, scraper interfaces.Scraper) *HoeService {
	return &HoeService{repo: repo, logger: logger, scraper: scraper}
}

func (hs *HoeService) ProcessListPage() {
	detailURLs, err := hs.scraper.GetDetailURLs()
	if err != nil {
		hs.logger.Fatal("process list page", err.Error())
	}

	if len(detailURLs) == 0 {
		hs.logger.Warn("No items found (Maybe Cloudflare block ?)")
	} else {
		hs.logger.Info(fmt.Sprintf("Found %d items\n", len(detailURLs)))

		for _, url := range detailURLs {
			hs.logger.Info(url)
		}
	}
}

func (hs *HoeService) ProcessDetailPage() {
	url := "/gai-goi/chi-tiet/56042/hot-girl-diep-anhmat-xinh-nguc-dep-bu-cu-dieu-luyen"

	hoe, err := hs.scraper.GetRawHoeData(url)
	if err != nil {
		hs.logger.Fatal(err.Error())
	}
}

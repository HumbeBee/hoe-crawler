package service

import (
	"fmt"
	"github.com/haovoanh28/gai-webscraper/internal/definitions"

	"github.com/haovoanh28/gai-webscraper/internal/interfaces"
	"github.com/haovoanh28/gai-webscraper/internal/repository"
	"github.com/haovoanh28/gai-webscraper/internal/utils/errutil"
	"github.com/haovoanh28/gai-webscraper/internal/utils/logutil"
)

type HoeService interface {
	ProcessListPage() error
	ProcessDetailPage(url string) error
}

type hoeService struct {
	hoeRepo            repository.HoeRepository
	workingHistoryRepo repository.WorkingHistoryRepository
	logger             *logutil.Logger
	scraper            interfaces.Scraper
	mapperService      definitions.MapperService
	validateService    definitions.ValidateService
}

func (hs *hoeService) ProcessListPage() error {
	detailURLs, err := hs.scraper.GetDetailURLs()
	if err != nil {
		return errutil.WrapError("get detail urls", err, "no target")
	}

	if len(detailURLs) == 0 {
		hs.logger.Warn("No items found (Maybe Cloudflare block ?)")
	} else {
		hs.logger.Info(fmt.Sprintf("Found %d items\n", len(detailURLs)))

		for _, url := range detailURLs {
			hs.logger.Info(url)
		}
	}

	return nil
}

func (hs *hoeService) ProcessDetailPage(url string) error {
	url2 := "/gai-goi/chi-tiet/56042/hot-girl-diep-anhmat-xinh-nguc-dep-bu-cu-dieu-luyen"

	rawHoe, err := hs.scraper.GetRawHoeData(url2)
	if err != nil {
		return errutil.WrapError("get raw hoe data", err, url)
	}

	hoeInfo := hs.mapperService.TransformHoe(rawHoe)
	hoeInfo.Print()

	if err := hs.validateService.ValidateHoe(hoeInfo); err != nil {
		return err
	}

	hoeInfo.Print()

	return nil
}

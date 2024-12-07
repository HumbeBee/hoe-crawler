package service

import (
	"fmt"

	"github.com/HumbeBee/hoe-crawler/internal/models"

	"github.com/HumbeBee/hoe-crawler/internal/interfaces"
	"github.com/HumbeBee/hoe-crawler/internal/repository"
	"github.com/HumbeBee/hoe-crawler/internal/utils/errutil"
	"github.com/HumbeBee/hoe-crawler/internal/utils/logutil"
)

type hoeService struct {
	locationRepo       repository.LocationRepository
	hoeRepo            repository.HoeRepository
	workingHistoryRepo repository.WorkingHistoryRepository
	logger             *logutil.Logger
	scraper            interfaces.Scraper
	mapperService      interfaces.MapperService
	validateService    interfaces.ValidateService
}

func (hs *hoeService) ProcessListPage() error {
	detailURLs, err := hs.scraper.GetDetailURLs()
	if err != nil {
		return errutil.WrapError("get detail urls", err)
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
	//url2 := "/gai-goi/chi-tiet/56042/hot-girl-diep-anhmat-xinh-nguc-dep-bu-cu-dieu-luyen"

	rawHoe, err := hs.scraper.GetRawHoeData(url)
	if err != nil {
		return fmt.Errorf("get raw hoe data: %v", err)
	}

	// If we can get location from database, it means the location is already valid
	cityID, err := hs.locationRepo.GetCityIDFromName(rawHoe.CityName)
	if err != nil {
		return fmt.Errorf("get city id: %v", err)
	}
	districtID, err := hs.locationRepo.GetDistrictIDFromName(rawHoe.DistrictName)
	if err != nil {
		return fmt.Errorf("get district id: %v", err)
	}

	// raw data to domain models
	hoeInfo := hs.mapperService.TransformHoe(rawHoe)

	isNewLocation, err := hs.workingHistoryRepo.CheckIsNewLocation(hoeInfo.ID, cityID, districtID)
	if err != nil {
		return fmt.Errorf("check is new location: %v", err)
	}
	if isNewLocation {
		// Create new working history
		hoeInfo.WorkingHistories = append(hoeInfo.WorkingHistories, models.WorkingHistory{
			HoeID:      hoeInfo.ID,
			CityID:     cityID,
			DistrictID: districtID,
		})
	}

	hoeInfo.Print()

	if err := hs.validateService.ValidateHoe(hoeInfo); err != nil {
		return err
	}

	if err = hs.hoeRepo.Save(hoeInfo); err != nil {
		return err
	}

	//
	//hoeInfo.Print()

	return nil
}

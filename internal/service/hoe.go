package service

import (
	"fmt"
	"github.com/HumbeBee/hoe-crawler/internal/dto"
	"github.com/HumbeBee/hoe-crawler/internal/infrastructure/browser"

	"github.com/HumbeBee/hoe-crawler/internal/models"

	"github.com/HumbeBee/hoe-crawler/internal/interfaces"
	"github.com/HumbeBee/hoe-crawler/internal/repository"
	"github.com/HumbeBee/hoe-crawler/internal/utils/errutil"
	"github.com/HumbeBee/hoe-crawler/internal/utils/logutil"
)

type hoeService struct {
	logger             *logutil.Logger
	browserRateLimiter *browser.RateLimiter
	locationRepo       repository.LocationRepository
	hoeRepo            repository.HoeRepository
	workingHistoryRepo repository.WorkingHistoryRepository
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
	rawHoeData, err := hs.GetRawHoeData(url)
	if err != nil {
		return fmt.Errorf("get raw hoe data: %v", err)
	}

	return hs.ProcessRawHoeData(rawHoeData)
}

func (hs *hoeService) GetRawHoeData(url string) (*dto.RawHoeData, error) {
	return hs.scraper.GetRawHoeData(url)
}

func (hs *hoeService) ProcessRawHoeData(rawHoeData *dto.RawHoeData) error {
	// If we can get location from database, it means the location is already valid
	cityID, err := hs.locationRepo.GetCityIDFromName(rawHoeData.CityName)
	if err != nil {
		return fmt.Errorf("get city id: %v", err)
	}
	districtID, err := hs.locationRepo.GetDistrictIDFromName(rawHoeData.DistrictName)
	if err != nil {
		return fmt.Errorf("get district id: %v", err)
	}

	// raw data to domain models
	hoeInfo := hs.mapperService.TransformHoe(rawHoeData)

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

	if err := hs.validateService.ValidateHoe(hoeInfo); err != nil {
		return err
	}

	if err = hs.hoeRepo.Save(hoeInfo); err != nil {
		return err
	}

	hoeInfo.Print()

	return nil
}

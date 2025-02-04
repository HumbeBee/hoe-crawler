package service

import (
	"fmt"
	"github.com/HumbeBee/hoe-crawler/internal/definitions"
	"github.com/HumbeBee/hoe-crawler/internal/dto"
	"github.com/HumbeBee/hoe-crawler/internal/infrastructure/browser"
	"sync"

	"github.com/HumbeBee/hoe-crawler/internal/models"

	"github.com/HumbeBee/hoe-crawler/internal/interfaces"
	"github.com/HumbeBee/hoe-crawler/internal/repository"
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

func (hs *hoeService) ProcessListPage(baseURL string, relativeURL string) ([]definitions.FailedData, error) {
	detailURLs, err := hs.scraper.GetDetailURLs(baseURL, relativeURL)
	if err != nil {
		return nil, fmt.Errorf("get detail urls: %v", err)
	}

	if len(detailURLs) == 0 {
		return nil, fmt.Errorf("no items found: %s", relativeURL)
	}

	fmt.Printf("Total - %d\n", len(detailURLs))

	var wg sync.WaitGroup
	failedURLs := make(chan definitions.FailedData, len(detailURLs))

	for i, url := range detailURLs {
		// Wait before making new request to the host
		if i > 0 {
			hs.browserRateLimiter.Wait()
		}

		rawData, err := hs.GetRawHoeData(baseURL, url)
		if err != nil {
			failedURLs <- definitions.FailedData{
				URL: url,
				Err: err,
			}
			continue
		}

		wg.Add(1)
		go func(data *dto.RawHoeData) {
			defer wg.Done()
			if err := hs.ProcessRawHoeData(data); err != nil {
				failedURLs <- definitions.FailedData{
					URL: url,
					Err: err,
				}
			}
		}(rawData)

		hs.logger.Info(url)
	}

	go func() {
		wg.Wait()
		close(failedURLs)
	}()

	var failedResults []definitions.FailedData
	for result := range failedURLs {
		failedResults = append(failedResults, result)
	}

	return failedResults, nil
}

func (hs *hoeService) ProcessDetailPage(baseURL string, relativeURL string) error {
	rawHoeData, err := hs.GetRawHoeData(baseURL, relativeURL)
	if err != nil {
		return fmt.Errorf("get raw hoe data: %v", err)
	}

	return hs.ProcessRawHoeData(rawHoeData)
}

func (hs *hoeService) GetRawHoeData(baseURL string, relativeURL string) (*dto.RawHoeData, error) {
	return hs.scraper.GetRawHoeData(baseURL, relativeURL)
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

	//hoeInfo.Print()

	return nil
}

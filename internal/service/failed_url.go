package service

import (
	"fmt"
	"github.com/HumbeBee/hoe-crawler/internal/infrastructure/browser"
	"github.com/HumbeBee/hoe-crawler/internal/interfaces"
	"github.com/HumbeBee/hoe-crawler/internal/models"
	"github.com/HumbeBee/hoe-crawler/internal/repository"
	"github.com/HumbeBee/hoe-crawler/internal/utils/logutil"
)

type failedURLService struct {
	siteID             uint
	logger             *logutil.Logger
	browserRateLimiter *browser.RateLimiter
	failedURLRepo      repository.FailedURLRepository
	siteRepo           repository.SiteRepository
	hoeService         interfaces.HoeService
}

func (f *failedURLService) TrackFailedURL(failedType models.FailedType, url string, err error) {
	var errorText string
	if err.Error() != "" {
		errorText = err.Error()
	} else {
		errorText = "unknown error"
	}

	f.logger.Error(fmt.Sprintf("failed to process url - %s, siteID: %d, error: %v", url, f.siteID, err))

	failedUrl, err := f.failedURLRepo.FindFailedURL(url, f.siteID)
	if err != nil {
		f.logger.Error(fmt.Sprintf(
			"failed to check url existence - url: %s, siteID: %d, error: %v",
			url, f.siteID, err))
		return
	}
	if failedUrl != nil {
		failedUrl.RetryCount++
		failedUrl.LastError = errorText
		if err := f.failedURLRepo.Save(failedUrl); err != nil {
			//return fmt.Errorf("failed to update failed url: %w", err)
			f.logger.Error(fmt.Sprintf(
				"failed to update failed url - url: %s, siteID: %d, error: %v",
				url, f.siteID, err))
		}

		return
	}

	if err := f.failedURLRepo.Save(&models.FailedURL{
		Type:       failedType,
		URL:        url,
		SiteID:     f.siteID,
		RetryCount: 0,
		LastError:  errorText,
	}); err != nil {
		f.logger.Error(fmt.Sprintf(
			"failed to create failed url - url: %s, siteID: %d, error: %v",
			url, f.siteID, err))
	}
}

func (f *failedURLService) RetryFailedURLs() error {
	failedUrls, err := f.failedURLRepo.GetFailedURLs()
	if err != nil {
		return fmt.Errorf("failed to get failed urls: %w", err)
	}

	siteInfo, err := f.siteRepo.GetSiteByID(f.siteID)
	if err != nil {
		return fmt.Errorf("failed to get site info: %w", err)
	}

	for i, url := range failedUrls {
		if i > 0 {
			f.browserRateLimiter.Wait()
		}

		f.logger.Info(fmt.Sprintf("Processing failed url: %s", url.URL))

		switch url.Type {
		case models.FailedTypeList:
		case models.FailedTypeDetail:
			if err := f.hoeService.ProcessDetailPage(siteInfo.BaseURL, url.URL); err != nil {
				f.TrackFailedURL(models.FailedTypeDetail, url.URL, err)
				continue
			}
		case models.FailedTypeReport:
		case models.FailedTypeUnknown:
		default:
			f.TrackFailedURL(models.FailedTypeUnknown, url.URL, fmt.Errorf("unknown failed type: %s", url.Type))
			continue
		}

		if err := f.failedURLRepo.Delete(url); err != nil {
			return fmt.Errorf("failed to delete failed url: %w", err)
		}
	}

	return nil
}

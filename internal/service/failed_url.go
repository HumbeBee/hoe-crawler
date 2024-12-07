package service

import (
	"fmt"
	"github.com/HumbeBee/hoe-crawler/internal/interfaces"
	"github.com/HumbeBee/hoe-crawler/internal/models"
	"github.com/HumbeBee/hoe-crawler/internal/repository"
	"github.com/HumbeBee/hoe-crawler/internal/utils/logutil"
)

type failedURLService struct {
	siteID uint
	repo   repository.FailedURLRepository
	logger *logutil.Logger
}

func NewFailedUrlService(siteID uint, repo repository.FailedURLRepository, logger *logutil.Logger) interfaces.FailedURLService {
	return &failedURLService{siteID: siteID, repo: repo, logger: logger}
}

func (f failedURLService) HandleFailedURL(url string, err error) {
	var errorText string
	if err.Error() != "" {
		errorText = err.Error()
	} else {
		errorText = "unknown error"
	}

	f.logger.Error(fmt.Sprintf(
		"failed to process url - %s, siteID: %d, error: %v",
		url, f.siteID, err))

	failedUrl, err := f.repo.FindFailedURL(url, f.siteID)
	if err != nil {
		f.logger.Error(fmt.Sprintf(
			"failed to check url existence - url: %s, siteID: %d, error: %v",
			url, f.siteID, err))
		return
	}
	if failedUrl != nil {
		failedUrl.RetryCount++
		failedUrl.LastError = errorText
		if err := f.repo.Save(failedUrl); err != nil {
			//return fmt.Errorf("failed to update failed url: %w", err)
			f.logger.Error(fmt.Sprintf(
				"failed to update failed url - url: %s, siteID: %d, error: %v",
				url, f.siteID, err))
		}

		return
	}

	if err := f.repo.Save(&models.FailedURL{
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

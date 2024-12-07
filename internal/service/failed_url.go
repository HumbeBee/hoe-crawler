package service

import (
	"fmt"
	"github.com/HumbeBee/hoe-crawler/internal/interfaces"
	"github.com/HumbeBee/hoe-crawler/internal/models"
	"github.com/HumbeBee/hoe-crawler/internal/repository"
)

type failedURLService struct {
	siteID uint
	repo   repository.FailedURLRepository
}

func NewFailedUrlService(siteID uint, repo repository.FailedURLRepository) interfaces.FailedURLService {
	return &failedURLService{siteID: siteID, repo: repo}
}

func (f failedURLService) HandleFailedURL(url string, siteID uint, err error) error {
	// TODO
	// 1. Check if url with given siteID exists in database
	// 2. If it does, increment retry count
	// 3. If it doesn't, create new entry
	failedUrl, err := f.repo.FindFailedURL(url, siteID)
	if err != nil {
		return fmt.Errorf("failed to check url existence (url=%s, siteID=%d): %w", url, siteID, err)
	}
	if failedUrl != nil {
		failedUrl.RetryCount++
		failedUrl.LastError = err.Error()
		if err := f.repo.Save(failedUrl); err != nil {
			return fmt.Errorf("failed to update failed url: %w", err)
		}
		return nil

	}

	if err := f.repo.Save(&models.FailedURL{
		URL:        url,
		SiteID:     siteID,
		RetryCount: 0,
		LastError:  err.Error(),
	}); err != nil {
		return fmt.Errorf("failed to create failed url: %w", err)
	}

	return nil
}

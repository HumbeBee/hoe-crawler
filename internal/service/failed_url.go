package service

import (
	"github.com/HumbeBee/hoe-crawler/internal/interfaces"
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

	//failedURL := &models.FailedURL{
	//	URL:        url,
	//	SiteID:     siteID,
	//	RetryCount: 0,
	//	LastError:  err.Error(),
	//}
	//
	//return f.repo.Save(failedURL)

	return nil
}

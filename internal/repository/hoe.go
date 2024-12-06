package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/HumbeBee/hoe-crawler/internal/models"
	"github.com/HumbeBee/hoe-crawler/internal/utils/logutil"
	"gorm.io/gorm"
)

type HoeRepository interface {
	Save(hoe *models.HoeInfo) error
	GetHoeByPhone(phone string) (*models.HoeInfo, error)
}

type hoeRepo struct {
	db     *gorm.DB
	logger *logutil.Logger
}

func NewHoeRepository(db *gorm.DB, logger *logutil.Logger) HoeRepository {
	return &hoeRepo{db: db, logger: logger}
}

func (r *hoeRepo) Save(hoe *models.HoeInfo) error {
	tx := r.db.Begin()

	// Try to find existing hoe by phone
	var existingHoe models.HoeInfo
	err := tx.Where("phone = ?", hoe.Phone).First(&existingHoe).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// First time this hoe is scraped
		if err := tx.Create(hoe).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to create new hoe: %v", err)
		}

		r.logger.Info(fmt.Sprintf("Created new hoe: %s with profile from %d",
			hoe.Name, hoe.Profiles[0].SiteID))

	} else if err != nil {
		tx.Rollback()
		return fmt.Errorf("database error: %v", err)

	} else {
		// Hoe exists, handle profile update
		if err := r.updateExistingHoe(tx, &existingHoe, hoe); err != nil {
			tx.Rollback()
			return err
		}

		r.logger.Info(fmt.Sprintf("Updated existing hoe: %s with profile from %s",
			hoe.Name, hoe.Profiles[0].Site.Name))
	}

	return tx.Commit().Error
}

func (r *hoeRepo) updateExistingHoe(tx *gorm.DB, existing *models.HoeInfo, new *models.HoeInfo) error {
	// Update common info
	if err := tx.Model(existing).Updates(models.HoeInfo{
		Name:          new.Name,
		BirthYear:     new.BirthYear,
		Height:        new.Height,
		Weight:        new.Weight,
		Country:       new.Country,
		LastScrapedAt: time.Now(), // Always non-zero since it's current time
	}).Error; err != nil {
		return fmt.Errorf("failed to update hoe info: %v", err)
	}

	// Handle profile update
	newProfile := new.Profiles[0]
	var existingProfile models.HoeProfile

	err := tx.Where("hoe_id = ? AND site_id = ?",
		existing.ID, newProfile.SiteID).First(&existingProfile).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// First profile from this site
		newProfile.HoeID = existing.ID
		if err := tx.Create(&newProfile).Error; err != nil {
			return fmt.Errorf("failed to create new profile: %v", err)
		}
	} else if err != nil {
		return fmt.Errorf("database error: %v", err)
	} else {
		// Update existing profile
		if err := tx.Model(&existingProfile).Updates(models.HoeProfile{
			OriginID: newProfile.OriginID,
			Url:      newProfile.Url,
			ImageUrl: newProfile.ImageUrl,
			Price:    newProfile.Price,
			Provider: newProfile.Provider,
			Area:     newProfile.Area,
			Status:   newProfile.Status,
			Service:  newProfile.Service,
			Duration: newProfile.Duration,
			WorkTime: newProfile.WorkTime,
		}).Error; err != nil {
			return fmt.Errorf("failed to update profile: %v", err)
		}
	}

	return nil
}

func (r *hoeRepo) GetHoeByPhone(phone string) (*models.HoeInfo, error) {
	var hoe models.HoeInfo
	if err := r.db.Where("phone = ?", phone).First(&hoe).Error; err != nil {
		return nil, err
	}
	return &hoe, nil
}

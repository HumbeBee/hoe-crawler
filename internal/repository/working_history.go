package repository

import (
	"github.com/HumbeBee/hoe-crawler/internal/models"
	"github.com/HumbeBee/hoe-crawler/internal/utils/logutil"
	"gorm.io/gorm"
)

type WorkingHistoryRepository interface {
	CheckIsNewLocation(hoeID, cityID uint, districtID uint) (bool, error)
}

type workingHistoryRepo struct {
	db     *gorm.DB
	logger *logutil.Logger
}

func NewWorkingHistoryRepository(db *gorm.DB, logger *logutil.Logger) WorkingHistoryRepository {
	return &workingHistoryRepo{db: db, logger: logger}
}

func (wh *workingHistoryRepo) CheckIsNewLocation(hoeID uint, cityID uint, districtID uint) (bool, error) {
	var count int64

	if err := wh.db.Model(&models.WorkingHistory{}).Where("hoe_id = ? AND city_id = ? AND district_id = ?", hoeID, cityID, districtID).Count(&count).Error; err != nil {
		return false, err
	}

	return count == 0, nil
}

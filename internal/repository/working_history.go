package repository

import (
	"github.com/HumbeBee/hoe-crawler/internal/models"
	"github.com/HumbeBee/hoe-crawler/internal/utils/logutil"
	"gorm.io/gorm"
)

type WorkingHistoryRepository interface {
	CheckIsNewLocation(hoeID, cityID uint, districtID uint) (bool, error)
}

type workingHistory struct {
	db     *gorm.DB
	logger *logutil.Logger
}

func NewWorkingHistoryRepository(db *gorm.DB, logger *logutil.Logger) WorkingHistoryRepository {
	return &workingHistory{db: db, logger: logger}
}

func (w *workingHistory) CheckIsNewLocation(hoeID uint, cityID uint, districtID uint) (bool, error) {
	var count int64
	//if err := w.db.Model(&models.WorkingHistory{}).Where("city_id = ? AND district_id = ?", cityID, districtID).Count(&count).Error; err != nil {
	//	return false, err
	//}

	if err := w.db.Model(&models.WorkingHistory{}).Where("hoe_id = ?", hoeID).Where("city_id = ?", cityID).Where("district_id = ?", districtID).Count(&count).Error; err != nil {
		return false, err
	}

	return count == 0, nil
}

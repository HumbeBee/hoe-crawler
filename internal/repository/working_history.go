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

func (wh *workingHistory) CheckIsNewLocation(hoeID uint, cityID uint, districtID uint) (bool, error) {
	var count int64
	//if err := wh.db.Model(&models.WorkingHistory{}).Where("city_id = ? AND district_id = ?", cityID, districtID).Count(&count).Error; err != nil {
	//	return false, err
	//}

	if err := wh.db.Model(&models.WorkingHistory{}).Where("hoe_id = ? AND city_id = ? AND district_id = ?", hoeID, cityID, districtID).Count(&count).Error; err != nil {
		return false, err
	}

	return count == 0, nil
}

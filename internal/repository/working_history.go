package repository

import (
	"github.com/haovoanh28/gai-webscraper/internal/utils/logutil"
	"gorm.io/gorm"
)

type WorkingHistoryRepository interface {
}

type workingHistory struct {
	db     *gorm.DB
	logger *logutil.Logger
}

func NewWorkingHistoryRepository(db *gorm.DB, logger *logutil.Logger) WorkingHistoryRepository {
	return &workingHistory{db: db, logger: logger}
}

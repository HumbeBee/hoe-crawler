package models

import "time"

type FailedURL struct {
	ID         uint      `gorm:"primaryKey"`
	URL        string    `gorm:"column:url"`
	SiteID     uint      `gorm:"column:site_id"`
	RetryCount uint      `gorm:"column:retry_count"`
	LastError  string    `gorm:"column:last_error;type:text"`
	LastTry    time.Time `gorm:"column:last_try"`
	CreatedAt  string    `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`

	Site Site `gorm:"foreignKey:SiteID"`
}

func (FailedURL) TableName() string {
	return "failed_urls"
}

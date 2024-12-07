package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type FailedType string

const (
	FailedTypeList    FailedType = "list"
	FailedTypeDetail  FailedType = "detail"
	FailedTypeReport  FailedType = "report"
	FailedTypeUnknown FailedType = "unknown"
)

type FailedURL struct {
	ID         uint       `gorm:"primaryKey"`
	URL        string     `gorm:"column:url"`
	SiteID     uint       `gorm:"column:site_id"`
	RetryCount uint       `gorm:"column:retry_count"`
	Type       FailedType `gorm:"column:type"`
	LastError  string     `gorm:"column:last_error;type:text"`
	LastTry    time.Time  `gorm:"column:last_try"`
	CreatedAt  time.Time  `gorm:"column:created_at"`
	UpdatedAt  time.Time  `gorm:"column:updated_at;autoUpdateTime"`

	Site Site `gorm:"foreignKey:SiteID"`
}

func (FailedURL) TableName() string {
	return "failed_urls"
}

func (f FailedType) Value() (driver.Value, error) {
	return string(f), nil
}

func (f *FailedType) Scan(value interface{}) error {
	if value == nil {
		*f = FailedTypeUnknown
		return nil
	}

	var str string
	switch v := value.(type) {
	case string:
		str = v
	case []byte:
		str = string(v)
	default:
		return fmt.Errorf("invalid failed type: %T", value)
	}

	*f = FailedType(str)
	return nil
}

package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type HoeStatus string

type ParsedAddress struct {
	Street   string
	District string
}

const (
	HoeStatusActive   HoeStatus = "active"
	HoeStatusInactive HoeStatus = "inactive"
	HoeStatusUnknown  HoeStatus = "unknown"
)

type HoeProfile struct {
	ID        uint      `gorm:"primaryKey"`
	HoeID     uint      `gorm:"column:hoe_id"`
	SiteID    uint      `gorm:"column:site_id;index"`
	OriginID  string    `gorm:"column:origin_id"`
	Url       string    `gorm:"column:url"`
	ImageUrl  string    `gorm:"column:image_url"`
	Price     string    `gorm:"column:price"`
	Area      string    `gorm:"column:area"`
	Provider  string    `gorm:"column:provider"`
	Status    HoeStatus `gorm:"column:status;type:enum('active','inactive','unknown')"`
	Service   string    `gorm:"column:service"`
	Duration  string    `gorm:"column:duration"`
	WorkTime  string    `gorm:"column:work_time"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	Hoe     HoeInfo     `gorm:"foreignKey:HoeID;constraint:OnDelete:CASCADE"`
	Site    Site        `gorm:"foreignKey:SiteID;constraint:OnDelete:CASCADE"`
	Reports []HoeReport `gorm:"foreignKey:HoeProfileID"`
}

func (p *HoeProfile) TableName() string {
	return "hoe_profiles"
}

// Using value receiver for Value() as we only need to read the status
// Using pointer receiver for Scan() as we need to modify the status
// This mixed receiver pattern is recommended by GORM
// Docs: https://gorm.io/docs/data_types.html

func (s HoeStatus) Value() (driver.Value, error) {
	return string(s), nil
}

func (s *HoeStatus) Scan(value interface{}) error {
	if value == nil {
		*s = HoeStatusUnknown
		return nil
	}

	// MySQL ENUM values can be returned as either string or []byte
	// Need to handle both cases to avoid panic when scanning
	var str string
	switch v := value.(type) {
	case string:
		str = v
	case []byte:
		str = string(v) // Convert []byte to string for ENUM values from DB
	default:
		return fmt.Errorf("invalid status value type: %T", value)
	}

	*s = HoeStatus(str)
	return nil
}

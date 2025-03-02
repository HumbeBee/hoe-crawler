package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type SiteRenderType string

const (
	SiteRenderTypeClient = "client"
	SiteRenderTypeServer = "server"
)

type Site struct {
	ID             uint           `gorm:"primaryKey"`
	Name           string         `gorm:"column:name"`
	BaseURL        string         `gorm:"column:base_url"`
	IsExpired      bool           `gorm:"column:is_expired"`
	SiteRenderType SiteRenderType `gorm:"column:site_render_type;type:enum('client','server')"`
	CreatedAt      time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;autoUpdateTime"`
}

func (Site) TableName() string {
	return "sites"
}

func (s SiteRenderType) Value() (driver.Value, error) {
	return string(s), nil
}

func (s *SiteRenderType) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var str string
	switch v := value.(type) {
	case string:
		str = v
	case []byte:
		str = string(v)
	default:
		return fmt.Errorf("unsupported type %T", v)
	}

	*s = SiteRenderType(str)
	return nil
}

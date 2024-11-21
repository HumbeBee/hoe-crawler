package models

import "time"

type WorkingHistory struct {
	ID         uint       `gorm:"primaryKey"`
	HoeID      uint       `gorm:"column:hoe_id;index"`
	StartDate  time.Time  `gorm:"column:start_date"`
	EndDate    *time.Time `gorm:"column:end_date"`
	CityID     uint       `gorm:"column:city_id;index"`
	DistrictID uint       `gorm:"column:district_id;index"`
	CreatedAt  time.Time  `gorm:"column:created_at"`
	UpdatedAt  time.Time  `gorm:"column:updated_at"`

	Hoe      HoeInfo  `gorm:"foreignKey:HoeID;"`
	City     City     `gorm:"foreignKey:CityID;references:ID;constraint:OnDelete:CASCADE"`
	District District `gorm:"foreignKey:DistrictID;references:ID;constraint:OnDelete:CASCADE"`
}

func (WorkingHistory) TableName() string {
	return "working_histories"
}

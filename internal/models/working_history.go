package models

import "time"

type WorkingHistory struct {
	ID         uint       `gorm:"primaryKey"`
	HoeID      uint       `gorm:"column:hoe_id"`
	StartDate  time.Time  `gorm:"column:start_date"`
	EndDate    *time.Time `gorm:"column:end_date"` // Make it pointer to allow NULL
	CityID     uint       `gorm:"column:city_id"`
	City       City       `gorm:"foreignKey:CityID;references:ID;constraint:OnDelete:CASCADE"`
	ProvinceID uint       `gorm:"column:province_id"`
	Province   Province   `gorm:"foreignKey:ProvinceID;references:ID;constraint:OnDelete:CASCADE"`
	CreatedAt  time.Time  `gorm:"column:created_at"`
	UpdatedAt  time.Time  `gorm:"column:updated_at"`
}

func (WorkingHistory) TableName() string {
	return "working_histories"
}

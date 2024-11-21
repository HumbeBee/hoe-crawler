package models

type Street struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"column:name"`
	Code       string `gorm:"column:code"`
	DistrictID uint   `gorm:"column:district_id;index"`
}

func (Street) TableName() string {
	return "street"
}

package models

type Street struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"column:name"`
	Code       string `gorm:"column:code"`
	CityID     string `gorm:"column:city"`
	DistrictID uint   `gorm:"column:district_id;index"`

	City     City     `gorm:"foreignKey:CityID;references:ID;constraint:OnDelete:CASCADE"`
	District District `gorm:"foreignKey:DistrictID;references:ID;constraint:OnDelete:CASCADE"`
}

func (Street) TableName() string {
	return "street"
}

package models

type Hotel struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"column:name;index"`
	EngName    string `gorm:"column:eng_name;index"`
	CityID     uint   `gorm:"column:city_id;index"`
	DistrictID uint   `gorm:"column:district_id;index"`

	City     City     `gorm:"foreignKey:CityID;references:ID;constraint:OnDelete:CASCADE"`
	District District `gorm:"foreignKey:DistrictID;references:ID;constraint:OnDelete:CASCADE"`
}

func (Hotel) TableName() string {
	return "hotels"
}

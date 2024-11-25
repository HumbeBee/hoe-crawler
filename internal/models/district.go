package models

type District struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"column:name;index"`
	ShortName string `gorm:"column:short_name;index"`
	Code      string `gorm:"column:code;index"`
	CityID    uint   `gorm:"column:city_id;index"`

	City    City     `gorm:"foreignKey:CityID;"`
	Streets []Street `gorm:"foreignKey:DistrictID;references:ID;constraint:OnDelete:CASCADE"`
}

func (District) TableName() string {
	return "districts"
}

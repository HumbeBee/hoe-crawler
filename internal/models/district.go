package models

type District struct {
	ID          uint   `gorm:"primaryKey"`
	CityID      uint   `gorm:"column:city_id;index"`
	DisplayName string `gorm:"column:display_name"`
	Name        string `gorm:"column:name;index"`
	ShortName   string `gorm:"column:short_name;index"`
	EngName     string `gorm:"column:eng_name;index"`
	Code        string `gorm:"column:code;index"`

	City    City     `gorm:"foreignKey:CityID;constraint:OnDelete:CASCADE"`
	Streets []Street `gorm:"foreignKey:DistrictID;references:ID;"`
}

func (District) TableName() string {
	return "districts"
}

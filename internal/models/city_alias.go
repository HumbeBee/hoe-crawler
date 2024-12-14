package models

type CityAlias struct {
	ID     uint   `gorm:"primaryKey"`
	CityID uint   `gorm:"column:city_id;index"`
	Alias  string `gorm:"column:alias;index;unique"`

	City City `gorm:"foreignKey:CityID;constraint:OnDelete:CASCADE"`
}

func (CityAlias) TableName() string {
	return "city_aliases"
}

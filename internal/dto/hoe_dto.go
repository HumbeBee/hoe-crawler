package dto

type RawHoeData struct {
	SiteID   uint
	OriginID string
	Url      string

	// All scraped fields without transformation
	Name      string
	ImageUrl  string
	Price     string
	Phone     string
	Address   string
	Provider  string
	Status    string
	BirthYear string
	Height    string
	Weight    string
	Country   string
	Service   string
	Duration  string
	WorkTime  string
}

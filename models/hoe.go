package models

type Hoe struct {
	ID         int            `json:"id"`
	Url        string         `json:"url"`
	MainInfo   *HoeMainInfo   `json:"main_info"`
	DetailInfo *HoeDetailInfo `json:"detail_info"`
	ReportInfo *HoeReportInfo `json:"report_info"`
}

type HoeMainInfo struct {
	Name        string `json:"name"`
	Price       string `json:"price"`
	ReportCount int    `json:"report_count"`
	ImageUrl    string `json:"image_url"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	Author      string `json:"author"`
	Status      string `json:"status"`
}

type HoeDetailInfo struct {
	From      string `json:"from"`
	BirthYear string `json:"birth_year"`
	Height    string `json:"height"`
	Weight    string `json:"weight"`
	Service   string `json:"service"`
	Duration  string `json:"duration"`
	WorkTime  string `json:"work_time"`
}

type HoeReportInfo struct {
	ID          string `json:"id"`
	Rating      string `json:"rating"`
	Author      string `json:"author"`
	Time        string `json:"time"`
	Description string `json:"description"`
}

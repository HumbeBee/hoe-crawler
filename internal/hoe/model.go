package hoe

import "github.com/haovoanh28/gai-webscraper/internal/reporter"

type Hoe struct {
	ID            string           `json:"id"`
	Url           string           `json:"url"`
	MainInfo      *HoeMainInfo     `json:"main_info"`
	DetailInfo    *HoeDetailInfo   `json:"detail_info"`
	ReportURLs    []string         `json:"report_urls"`
	DetailReports []*HoeReportInfo `json:"detail_reports"`
}

type HoeMainInfo struct {
	Name     string `json:"name"`
	Price    string `json:"price"`
	ImageUrl string `json:"image_url"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Author   string `json:"author"`
	Status   string `json:"status"`
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
	ID          string             `json:"id"`
	DetailUrl   string             `json:"detail_url"`
	Rating      string             `json:"rating"`
	Author      *reporter.Reporter `json:"author"`
	Time        string             `json:"time"`
	Description string             `json:"description"`
}

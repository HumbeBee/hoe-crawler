package models

type ReportInfo struct {
	ID          string    `json:"id"`
	DetailUrl   string    `json:"detail_url"`
	Rating      string    `json:"rating"`
	Author      *Reporter `json:"author"`
	Time        string    `json:"time"`
	Description string    `json:"description"`
}

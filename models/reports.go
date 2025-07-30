package models

// ReportsRequest represents report filter parameters
type ReportsRequest struct {
	StartDate string
	EndDate   string
	State     string
	Source    string
	Type      string
	ReportID  string
	Page      string
	Count     string
}

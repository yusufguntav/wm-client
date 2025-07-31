package models

import "time"

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

type ReportResponse struct {
	SMS           interface{} `json:"sms"`
	SMSTotalCount int         `json:"smsTotalCount"`
	Status        int         `json:"status"`
	WP            []WPReport  `json:"wp"`
	WPTotalCount  int         `json:"wpTotalCount"`
}

type WPReport struct {
	ID            int        `json:"ID"`
	CreatedAt     time.Time  `json:"CreatedAt"`
	UpdatedAt     time.Time  `json:"UpdatedAt"`
	DeletedAt     *time.Time `json:"DeletedAt"`
	ReportID      string     `json:"report_id"`
	UserID        int        `json:"user_id"`
	RegID         string     `json:"reg_id"`
	Phone         string     `json:"phone"`
	TotalCount    int        `json:"total_count"`
	Success       int        `json:"success"`
	Fail          int        `json:"fail"`
	CampaignName  string     `json:"campaign_name"`
	CreditBefore  int        `json:"credit_before"`
	State         int        `json:"state"`
	ReportType    int        `json:"report_type"`
	SendTime      time.Time  `json:"send_time"`
	OldID         *int       `json:"old_id"`
	TZString      string     `json:"tz_string"`
	AutoMessageID int        `json:"auto_message_id"`
	UserName      string     `json:"user_name"`
	Content       string     `json:"content"`
}

type ReportDetailResponse struct {
	SMS           interface{}      `json:"sms"`
	SMSTotalCount int              `json:"smsTotalCount"`
	Status        int              `json:"status"`
	WP            []WPDetailReport `json:"wp"`
	WPTotalCount  int              `json:"wpTotalCount"`
}

type WPDetailReport struct {
	ID               int        `json:"ID"`
	CreatedAt        time.Time  `json:"CreatedAt"`
	UpdatedAt        time.Time  `json:"UpdatedAt"`
	DeletedAt        *time.Time `json:"DeletedAt"`
	Phone            string     `json:"phone"`
	CountryCode      string     `json:"country_code"`
	UserID           int        `json:"user_id"`
	ReportID         string     `json:"report_id"`
	SendDateFromApp  time.Time  `json:"send_date_from_app"`
	SendDateFromCore time.Time  `json:"send_date_from_core"`
	CurrentWPCredit  int        `json:"current_wp_credit"`
	AfterWPCredit    int        `json:"after_wp_credit"`
	Content          string     `json:"content"`
	State            int        `json:"state"`
	Type             int        `json:"type"`
	SendDate         time.Time  `json:"send_date"`
	UserName         string     `json:"user_name"`
}

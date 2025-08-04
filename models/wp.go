package models

import (
	"time"
)

type WpSendResponse struct {
	Message  string `json:"message"`
	ReportID string `json:"report_id"`
}

type SendWpRequest struct {
	ID string `json:"id"`
}

type PreviewWpRequest struct {
	Numbers            string `json:"numbers" form:"numbers"`                             // "+905347882505"
	Message            string `json:"message" form:"message"`                             // "dawdawdwa"
	CampaignName       string `json:"campaign_name" form:"campaign_name"`                 // "testtt"
	RegID              string `json:"reg_id" form:"reg_id"`                               // "2817162021"
	Now                bool   `json:"now" form:"now"`                                     // true
	SendSpeed          int    `json:"send_speed" form:"send_speed"`                       // 4
	SendDate           string `json:"send_date" form:"send_date"`                         // "2025-08-01T13:53:46+03:00"
	BulkWpExcelNumbers string `json:"bulk_wp_excel_numbers" form:"bulk_wp_excel_numbers"` // "[object Object]"
	BulkWpFile         string `json:"-" form:"bulk_wp_file"`                              // Dosya (JSON gönderiminde yok)
	AddCancelLink      bool   `json:"add_cancel_link" form:"add_cancel_link"`             // false
	TzString           string `json:"tz_string" form:"tz_string"`                         // "(UTC+03:00) Istanbul"
}
type BulkWpPreviewResponse struct {
	ID                      string    `json:"id"`
	TotalSuccessMembers     int       `json:"total_success_members"`
	TotalWpCredit           int       `json:"total_wp_credit"`
	Message                 string    `json:"message"`
	LastStr                 string    `json:"last_str"`
	AddCancelLink           bool      `json:"AddCancelLink"`
	Top10ValidNumbers       []string  `json:"top10_valid_numbers"`
	Top10BlacklistedNumbers []string  `json:"top10_blacklisted_numbers"`
	Top10InvalidNumbers     []string  `json:"top10_invalid_numbers"`
	AllValidNumbers         []string  `json:"all_valid_numbers"`
	AllBlacklistNumbers     []string  `json:"all_blacklist_numbers"`
	AllInvalidNumbers       []string  `json:"all_invalid_numbers"`
	CountBlacklistedNumbers int       `json:"count_blacklisted_numbers"`
	CountInvalidNumbers     int       `json:"count_invalid_numbers"`
	CountValidNumbers       int       `json:"count_valid_numbers"`
	CountExcelNumbers       int       `json:"count_excel_numbers"`
	CampaignName            string    `json:"campaign_name"`
	Now                     bool      `json:"now"`
	SendDate                time.Time `json:"send_date"`
	WpCreditCurrent         int       `json:"wp_credit_current"`
	WpCreditAfterSend       int       `json:"wp_credit_after_send"`
	File                    string    `json:"file"`
	FileName                string    `json:"file_name"`
	RegID                   string    `json:"reg_id"`
	TzString                string    `json:"tz_string"`
	SendSpeed               int       `json:"send_speed"`
	DelayMin                int       `json:"delay_min"`
	DelayMax                int       `json:"delay_max"`
	UserID                  int       `json:"user_id"`
}

package models

import (
	"time"
)

// /bulk/preview/sms
// /bulk/sms

type CharacterType int

const (
	CharacterType_Normal CharacterType = iota
	CharacterType_Turkish
	CharacterType_Unicode
)

type MessageType int

const (
	MessageType_Information MessageType = iota
	MessageType_Commercial
)

type SmsSendResponse string

// type SmsSendResponse struct {
// 	Message  string `json:"message"`
// 	ReportID string `json:"report_id"`
// }

type SendSmsRequest struct {
	ID string `json:"id"`
}

type PreviewSmsRequest struct {
	Numbers       []string      `form:"numbers"`     // "+905555555555"
	Message       string        `form:"message"`     // "message content"
	SenderName    int           `form:"sender_name"` // ID of sender name
	SendDate      time.Time     `form:"send_date"`   // "2025-08-01T13:53:46+03:00"
	CharacterType CharacterType `form:"character_type"`
	MessageType   MessageType   `form:"message_type"`
	AddCancelLink bool          `form:"add_cancel_link"`
}

type BulkSmsPreviewResponse struct {
	ID                  string `json:"id"`
	TotalSuccessMembers int    `json:"total_success_members"`
	TotalSmsCredit      int    `json:"total_sms_credit"`
	TotalCount          int    `json:"total_count"`
	Message             string `json:"message"`
	AddCancelLink       bool   `json:"add_cancel_link"`

	Top10ValidNumbers       []string `json:"top10_valid_numbers"`
	Top10BlacklistedNumbers []string `json:"top10_blacklisted_numbers"`
	Top10InvalidNumbers     []string `json:"top10_invalid_numbers"`

	AllValidNumbers     []string `json:"all_valid_numbers"`
	AllBlacklistNumbers []string `json:"all_blacklist_numbers"`
	AllInvalidNumbers   []string `json:"all_invalid_numbers"`

	CountBlacklistedNumbers int `json:"count_blacklisted_numbers"`
	CountInvalidNumbers     int `json:"count_invalid_numbers"`
	CountValidNumbers       int `json:"count_valid_numbers"`
	CountExcelNumbers       int `json:"count_excel_numbers"`

	SenderName    string        `json:"sender_name"`
	SendDate      time.Time     `json:"send_date"`
	CharacterType CharacterType `json:"character_type"`
	MessageType   MessageType   `json:"message_type"`
}

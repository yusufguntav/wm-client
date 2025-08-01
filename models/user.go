package models

import (
	"time"
)

type SubuserData struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Surname    string  `json:"surname"`
	Username   string  `json:"username"`
	Phone      string  `json:"phone"`
	Email      string  `json:"email"`
	Status     int     `json:"status"`
	SMSCredit  int     `json:"sms_credit"`
	WPCredit   int     `json:"wp_credit"`
	Sendername *string `json:"sendername"`
}

type ResponseForSubuser struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Username   string `json:"username"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Status     int    `json:"status"`
	SMSCredit  uint   `json:"sms_credit"`
	WPCredit   uint   `json:"wp_credit"`
	Sendername []uint `json:"sendername"`
}

type CreateSubuser struct {
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	SMSCreditLimit int    `json:"sms_credit_limit"`
	WPCreditLimit  int    `json:"wp_credit_limit"`
	Senders        []uint `json:"senders"`
}

type UpdateSubuser struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Senders  []uint `json:"senders"`
}

type UserResponse struct {
	ID                uint       `json:"id"`
	Name              string     `json:"name"`
	Surname           string     `json:"surname"`
	Username          string     `json:"username"`
	Phone             string     `json:"phone"`
	Email             string     `json:"email"`
	CompanyName       string     `json:"company_name"`
	Status            int        `json:"status"`
	CancelLinkEndDate *time.Time `json:"cancel_link_end_date"`
	RequireLoginOTP   bool       `json:"require_login_otp"`
	Project           int        `json:"project"`
}

type Birthday struct {
	Time  time.Time `json:"Time"`
	Valid bool      `json:"Valid"`
}

type UserDetailResponse struct {
	ID                     uint         `json:"id"`
	UserID                 uint         `json:"user_id"`
	User                   UserResponse `json:"user"`
	Birthday               Birthday     `json:"birthday"`
	City                   string       `json:"city"`
	District               string       `json:"district"`
	Address                string       `json:"address"`
	SmsCredit              uint64       `json:"sms_credit"`
	WpCredit               uint64       `json:"wp_credit"`
	CancelLinkEndsAt       *time.Time   `json:"cancel_link_ends_at"`
	IncomingMessagesEndsAt *time.Time   `json:"incoming_messages_ends_at"`
	FeedbackDate           *time.Time   `json:"feedback_date"`
	SmsWpSettings          int          `json:"sms_wp_settings"`
}

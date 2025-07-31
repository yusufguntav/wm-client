package models

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

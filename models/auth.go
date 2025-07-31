package models

type LoginVerifyCodePayload struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type LoginPayload struct {
	Identifier string `json:"identifier"`
	VerifyCode string `json:"verify_code"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type VerificationData struct {
	ID               int    `json:"id"`
	PhoneNumber      string `json:"phone_number"`
	VerificationCode string `json:"verification_code"`
	UserID           int    `json:"user_id"`
}

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

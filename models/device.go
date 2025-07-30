package models

type ConnectCodeRequest struct {
	Phone string `json:"phone"`
}

type CheckDeviceRequest struct {
	RegID string `json:"reg_id"`
	Phone string `json:"phone"`
}

package models

type ConnectCodeRequest struct {
	Phone string `json:"phone"`
}

type CheckDeviceRequest struct {
	RegID string `json:"reg_id"`
	Phone string `json:"phone"`
}

type GetDevicesRequest struct {
	RegID string `json:"reg_id"`
	Phone string `json:"phone"`
}

type DeviceResponse struct {
	JID            string `json:"j_id"`
	RegistrationID string `json:"registration_id"`
	Platform       string `json:"platform"`
	PushName       string `json:"push_name"`
	BusinessName   string `json:"business_name"`
	DeviceNumber   string `json:"device_number"`
	State          string `json:"state"`
	UserID         int    `json:"user_id"`
	UserName       string `json:"user_name"`
	LogoutDate     string `json:"logout_date"`
}

type ConnectCodeResponse struct {
	Code  string `json:"code"`
	RegID string `json:"regId"`
}

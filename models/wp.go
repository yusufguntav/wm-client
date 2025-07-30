package models

type PreviewWpRequest struct {
	Numbers            string `json:"numbers"`
	Message            string `json:"message"`
	CampaignName       string `json:"campaign_name"`
	RegID              string `json:"reg_id"`
	Now                string `json:"now"`
	SendSpeed          string `json:"send_speed"`
	SendDate           string `json:"send_date,omitempty"`
	AddCancelLink      string `json:"add_cancel_link"`
	BulkWpExcelNumbers string `json:"bulk_wp_excel_numbers,omitempty"`
	BulkWpFile         string `json:"bulk_wp_file,omitempty"`
}

type PreviewWpResponse struct {
	ID string `json:"id"`
}

type SendWpRequest struct {
	ID string `json:"id"`
}

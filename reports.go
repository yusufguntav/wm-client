package wmclient

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/yusufguntav/wm-client/models"
)

func (c *Client) GetReports(req models.ReportsRequest) (models.ReportResponse, error) {
	params := url.Values{}
	params.Add("start_date", req.StartDate)
	params.Add("end_date", req.EndDate)
	params.Add("state", req.State)
	params.Add("source", req.Source)
	params.Add("type", req.Type)
	params.Add("report_id", req.ReportID)
	params.Add("page", req.Page)
	params.Add("count", req.Count)

	endpoint := fmt.Sprintf("/reports/multi?%s", params.Encode())

	respBody, err := c.doRequest("GET", endpoint, nil)
	if err != nil {
		return models.ReportResponse{}, err
	}

	var result models.ReportResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return models.ReportResponse{}, fmt.Errorf("unmarshal error: %w", err)
	}

	return result, nil
}

func (c *Client) GetReportDetail(reportID, source, state, page string) (models.ReportDetailResponse, error) {
	params := url.Values{}
	params.Add("report_id", reportID)
	params.Add("source", source)
	params.Add("state", state)
	params.Add("page", page)

	endpoint := fmt.Sprintf("/reports?%s", params.Encode())
	respBody, err := c.doRequest("GET", endpoint, nil)
	if err != nil {
		return models.ReportDetailResponse{}, err
	}

	var result models.ReportDetailResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return models.ReportDetailResponse{}, fmt.Errorf("unmarshal error: %w", err)
	}

	return result, nil
}

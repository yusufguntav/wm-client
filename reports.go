package wmclient

import (
	"fmt"
	"net/url"

	"github.com/yusufguntav/wm-client/models"
)

// GetReports gets multiple reports with filters
func (c *Client) GetReports(req models.ReportsRequest) ([]byte, error) {
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
	return c.doRequest("GET", endpoint, nil)
}

// GetReportDetail retrieves a detailed single report
func (c *Client) GetReportDetail(reportID, source, state, page string) ([]byte, error) {
	params := url.Values{}
	params.Add("report_id", reportID)
	params.Add("source", source)
	params.Add("state", state)
	params.Add("page", page)

	endpoint := fmt.Sprintf("/reports?%s", params.Encode())
	return c.doRequest("GET", endpoint, nil)
}

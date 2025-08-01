package wmclient

import (
	"encoding/json"
	"fmt"

	"github.com/yusufguntav/wm-client/models"
)

func (c *Client) SendWpPreview(req models.PreviewWpRequest) (models.BulkWpPreviewResponse, error) {
	respBody, err := c.doRequest("POST", "/bulk/preview/wp", req)
	if err != nil {
		return models.BulkWpPreviewResponse{}, fmt.Errorf("send wp preview error: %w, response: %s", err, respBody)
	}

	var previewResp models.APIResponse[models.BulkWpPreviewResponse]
	if err := json.Unmarshal(respBody, &previewResp); err != nil {
		return models.BulkWpPreviewResponse{}, fmt.Errorf("unmarshal error: %w", err)
	}

	return previewResp.Data, nil
}

func (c *Client) SendWp(req models.SendWpRequest) (models.WpSendResponse, error) {
	respBody, err := c.doRequest("POST", "/bulk/wp", req)
	if err != nil {
		return models.WpSendResponse{}, err
	}

	var resp models.APIResponse[models.WpSendResponse]
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return models.WpSendResponse{}, fmt.Errorf("unmarshal error: %w", err)
	}

	return resp.Data, nil
}

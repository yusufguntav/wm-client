package wmclient

import (
	"encoding/json"
	"fmt"

	"github.com/yusufguntav/wmclient/models"
)

// SendWp sends a WhatsApp message after preview
func (c *Client) SendWp(req models.SendWpRequest) error {

	_, err := c.doRequest("POST", "/bulk/wp", req)
	return err
}

// PreviewWp generates WhatsApp message preview and returns an ID
func (c *Client) PreviewWp(req models.PreviewWpRequest) (string, error) {
	respBody, err := c.doRequest("POST", "/bulk/preview/wp", req)
	if err != nil {
		return "", err
	}

	var resp PreviewWpResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return "", fmt.Errorf("unmarshal error: %w", err)
	}

	return resp.ID, nil
}

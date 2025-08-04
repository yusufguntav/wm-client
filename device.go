package wmclient

import (
	"encoding/json"
	"fmt"

	"github.com/yusufguntav/wm-client/models"
)

func (c *Client) ConnectCode(req models.ConnectCodeRequest) (models.ConnectCodeResponse, error) {

	resp, err := c.doRequest("POST", "/wp/login/code", req)

	if err != nil {
		return models.ConnectCodeResponse{}, fmt.Errorf("connect code error: %w, response: %s", err, resp)
	}

	var connectResp models.APIResponse[models.ConnectCodeResponse]
	if err := json.Unmarshal(resp, &connectResp); err != nil {
		return models.ConnectCodeResponse{}, fmt.Errorf("unmarshal error: %w", err)
	}

	return connectResp.Data, nil

}

func (c *Client) CheckDevice(req models.CheckDeviceRequest) error {
	_, err := c.doRequest("POST", "/wp/device/check", req)
	return err
}

func (c *Client) DeleteDevice(regID string) error {
	endpoint := fmt.Sprintf("/wp/delete/%s", regID)
	_, err := c.doRequest("DELETE", endpoint, nil)
	return err
}

func (c *Client) LogoutDevice(regID string) error {
	endpoint := fmt.Sprintf("/wp/logout/%s", regID)
	_, err := c.doRequest("POST", endpoint, nil)
	return err
}

func (c *Client) GetDevices(req models.GetDevicesRequest) ([]models.DeviceResponse, error) {
	resp, err := c.doRequest("GET", "/wp/device", req)
	if err != nil {
		return nil, fmt.Errorf("get devices error: %w", err)
	}

	var deviceResp models.APIResponse[[]models.DeviceResponse]
	if err := json.Unmarshal(resp, &deviceResp); err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}
	if len(deviceResp.Data) == 0 {
		return nil, fmt.Errorf("no devices found")
	}

	return deviceResp.Data, nil
}

package wmclient

import (
	"fmt"

	"github.com/yusufguntav/wm-client/models"
)

// ConnectCode initiates a connection with device
func (c *Client) ConnectCode(req models.ConnectCodeRequest) error {
	_, err := c.doRequest("POST", "/wp/login/code", req)
	return err
}

// CheckDevice checks if the device is connected
func (c *Client) CheckDevice(req models.CheckDeviceRequest) error {
	_, err := c.doRequest("POST", "/wp/device/check", req)
	return err
}

// DeleteDevice removes a device connection
func (c *Client) DeleteDevice(regID string) error {
	endpoint := fmt.Sprintf("/wp/delete/%s", regID)
	_, err := c.doRequest("DELETE", endpoint, nil)
	return err
}

// GetDevices fetches devices based on phone and reg_id
func (c *Client) GetDevices(phone, regID string) ([]byte, error) {
	payload := map[string]string{
		"phone":  phone,
		"reg_id": regID,
	}
	return c.doRequest("GET", "/wp/device", payload)
}

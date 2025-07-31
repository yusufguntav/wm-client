package wmclient

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/yusufguntav/wm-client/models"
)

func (c *Client) CreateSubuser(req models.CreateSubuser) error {
	_, err := c.doRequest("POST", "/subuser", req)
	return err
}

func (c *Client) GetSubusers() (models.APIResponse[[]models.SubuserData], error) {
	respBody, err := c.doRequest("GET", "/subuser", nil)
	if err != nil {
		return models.APIResponse[[]models.SubuserData]{}, err
	}

	var result models.APIResponse[[]models.SubuserData]
	if err := json.Unmarshal(respBody, &result); err != nil {
		return models.APIResponse[[]models.SubuserData]{}, fmt.Errorf("unmarshal error: %w", err)
	}

	return result, nil
}

func (c *Client) GetSubuserByID(id string) (models.ResponseForSubuser, error) {
	endpoint := fmt.Sprintf("/subuser/%s", id)
	respBody, err := c.doRequest("GET", endpoint, nil)
	if err != nil {
		return models.ResponseForSubuser{}, err
	}

	var result models.ResponseForSubuser
	if err := json.Unmarshal(respBody, &result); err != nil {
		return models.ResponseForSubuser{}, fmt.Errorf("unmarshal error: %w", err)
	}

	return result, nil
}

func (c *Client) UpdateSubuser(req models.UpdateSubuser) error {
	endpoint := fmt.Sprintf("/subuser/%s", strconv.FormatUint(uint64(req.ID), 10))
	_, err := c.doRequest("PUT", endpoint, req)
	return err
}

func (c *Client) DeleteSubuser(id string) error {
	endpoint := fmt.Sprintf("/subuser/%s", id)
	_, err := c.doRequest("DELETE", endpoint, nil)
	return err
}

// UserDetail fetches user detail info
func (c *Client) UserDetail() ([]byte, error) {
	return c.doRequest("GET", "/user/detail", nil)
}

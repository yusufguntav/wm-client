package wmclient

import (
	"fmt"
	"net/url"
)

// CreateSubuser creates a new subuser
func (c *Client) CreateSubuser(params url.Values) error {
	endpoint := fmt.Sprintf("/subuser?%s", params.Encode())
	_, err := c.doRequest("POST", endpoint, nil)
	return err
}

// GetSubusers fetches all subusers
func (c *Client) GetSubusers() ([]byte, error) {
	return c.doRequest("GET", "/subuser", nil)
}

// GetSubuserByID fetches a subuser by ID
func (c *Client) GetSubuserByID(id string) ([]byte, error) {
	endpoint := fmt.Sprintf("/subuser/%s", id)
	return c.doRequest("GET", endpoint, nil)
}

// UpdateSubuser updates subuser details by ID
func (c *Client) UpdateSubuser(id string, params url.Values) error {
	endpoint := fmt.Sprintf("/subuser/%s?%s", id, params.Encode())
	_, err := c.doRequest("PUT", endpoint, nil)
	return err
}

// DeleteSubuser deletes subuser by ID
func (c *Client) DeleteSubuser(id string) error {
	endpoint := fmt.Sprintf("/subuser/%s", id)
	_, err := c.doRequest("DELETE", endpoint, nil)
	return err
}

// UserDetail fetches user detail info
func (c *Client) UserDetail() ([]byte, error) {
	return c.doRequest("GET", "/user/detail", nil)
}

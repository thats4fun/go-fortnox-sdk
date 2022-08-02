package client

import "context"

const (
	meURI = "me"
)

// GetMeInformation does _GET https://api.fortnox.se/3/me
func (c *Client) GetMeInformation(ctx context.Context) (*MeInformation, error) {
	resp := &GetMeInformationResp{}

	err := c._GET(ctx, meURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.MeInformation, nil
}

type MeInformation struct {
	Id       string `json:"Id,omitempty"`
	Name     string `json:"Name,omitempty"`
	Email    string `json:"Email,omitempty"`
	SysAdmin bool   `json:"SysAdmin,omitempty"`
	Locale   string `json:"Locale,omitempty"`
}

type GetMeInformationResp struct {
	MeInformation MeInformation `json:"MeInformation"`
}

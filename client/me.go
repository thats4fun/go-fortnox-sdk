package client

import "context"

const (
	meURI = "me"
)

// GetMeInformation does _GET https://api.fortnox.se/3/me
func (c *Client) GetMeInformation(ctx context.Context) (*GetMeInformationResp, error) {
	resp := &GetMeInformationResp{}

	err := c._GET(ctx, meURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type MeInformation struct {
	Id       string `json:"Id"`
	Name     string `json:"Name"`
	Email    string `json:"Email"`
	SysAdmin bool   `json:"SysAdmin"`
	Locale   string `json:"Locale"`
}

type GetMeInformationResp struct {
	MeInformation MeInformation `json:"MeInformation"`
}

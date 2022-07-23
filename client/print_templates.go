package client

import "context"

const (
	printTemplatesURI = "printtemplates"
)

// GetAllPrintTemplates does _GET https://api.fortnox.se/3/printtemplates
func (c *Client) GetAllPrintTemplates(ctx context.Context) (*GetAllPrintTemplatesResp, error) {
	resp := &GetAllPrintTemplatesResp{}

	err := c._GET(ctx, printTemplatesURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetAllPrintTemplatesResp struct {
	PrintTemplates []struct {
		Template string `json:"Template"`
		Name     string `json:"Name"`
	} `json:"PrintTemplates"`
}

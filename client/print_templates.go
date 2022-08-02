package client

import "context"

const (
	printTemplatesURI = "printtemplates"
)

// GetAllPrintTemplates does _GET https://api.fortnox.se/3/printtemplates
func (c *Client) GetAllPrintTemplates(ctx context.Context) ([]PrintTemplate, error) {
	resp := &GetAllPrintTemplatesResp{}

	err := c._GET(ctx, printTemplatesURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.PrintTemplates, nil
}

type GetAllPrintTemplatesResp struct {
	PrintTemplates []PrintTemplate `json:"PrintTemplates"`
}

type PrintTemplate struct {
	Template string `json:"Template,omitempty"`
	Name     string `json:"Name,omitempty"`
}

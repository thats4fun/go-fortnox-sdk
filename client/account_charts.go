package client

import "context"

const (
	accountChartsURI = "accountcharts"
)

// GetAccountCharts does _GET https://api.fortnox.se/3/accountcharts
func (c *Client) GetAccountCharts(ctx context.Context) ([]AccountChart, error) {
	resp := &AccountChartResp{}

	err := c._GET(ctx, accountChartsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.AccountCharts, nil
}

type AccountChart struct {
	Name string `json:"Name,omitempty"`
}

type AccountChartResp struct {
	AccountCharts []AccountChart `json:"AccountCharts"`
}

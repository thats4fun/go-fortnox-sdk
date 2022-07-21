package client

import "context"

const (
	accountChartsURI = "accountcharts"
)

// GetAccountCharts does _GET https://api.fortnox.se/3/accountcharts
func (c *Client) GetAccountCharts(ctx context.Context) (*AccountChartResp, error) {
	resp := &AccountChartResp{}

	err := c._GET(ctx, accountChartsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type AccountChartResp struct {
	AccountCharts []struct {
		Name string `json:"Name"`
	} `json:"AccountCharts"`
}

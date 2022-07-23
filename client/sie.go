package client

import (
	"context"
	"fmt"
)

const (
	sieURI = "sie"
)

// GetSIEFile does _GET https://api.fortnox.se/3/sie/{Type}
//
// typ - type
//
// filter - FinancialYearFilter
func (c *Client) GetSIEFile(ctx context.Context, typ string, filter *FinancialYearFilter) (*[]byte, error) {
	resp := &[]byte{}

	uri := fmt.Sprintf("%s/%s", sieURI, typ)

	params := filter.urlValues()

	err := c._GET(ctx, uri, params, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

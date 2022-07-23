package client

import (
	"context"
	"fmt"
)

const (
	predefinedVoucherSeriesURI = "predefinedvoucherseries"
)

// GetAllPredefinedVoucherSeries does _GET https://api.fortnox.se/3/predefinedvoucherseries/
func (c *Client) GetAllPredefinedVoucherSeries(ctx context.Context) (*GetAllPredefinedVoucherSeriesResp, error) {
	resp := &GetAllPredefinedVoucherSeriesResp{}

	err := c._GET(ctx, predefinedVoucherSeriesURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetPredefinedVoucherSeries does _GET https://api.fortnox.se/3/predefinedvoucherseries/{Name}
//
// name - identifies the predefined voucher series
func (c *Client) GetPredefinedVoucherSeries(ctx context.Context, name string) (*GetPredefinedVoucherSeriesResp, error) {
	resp := &GetPredefinedVoucherSeriesResp{}

	uri := fmt.Sprintf("%s/%s", predefinedVoucherSeriesURI, name)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdatePredefinedVoucherSeries does _PUT https://api.fortnox.se/3/predefinedvoucherseries/{Name}
//
// name - identifies the predefined voucher series
//
// req - predefined voucher series to update
func (c *Client) UpdatePredefinedVoucherSeries(
	ctx context.Context,
	name string,
	req *UpdatePredefinedVoucherSeriesReq) (*UpdatePredefinedVoucherSeriesResp, error) {

	resp := &UpdatePredefinedVoucherSeriesResp{}

	uri := fmt.Sprintf("%s/%s", predefinedVoucherSeriesURI, name)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetAllPredefinedVoucherSeriesResp struct {
	PreDefinedVoucherSeriesCollection []struct {
		Url           string `json:"@url"`
		Name          string `json:"Name"`
		VoucherSeries string `json:"VoucherSeries"`
	} `json:"PreDefinedVoucherSeriesCollection"`
}

type GetPredefinedVoucherSeriesResp struct {
	PreDefinedVoucherSeries struct {
		Url           string `json:"@url"`
		Name          string `json:"Name"`
		VoucherSeries string `json:"VoucherSeries"`
	} `json:"PreDefinedVoucherSeries"`
}

type UpdatePredefinedVoucherSeriesReq struct {
	PreDefinedVoucherSeries struct {
		Url           string `json:"@url"`
		Name          string `json:"Name"`
		VoucherSeries string `json:"VoucherSeries"`
	} `json:"PreDefinedVoucherSeries"`
}

type UpdatePredefinedVoucherSeriesResp struct {
	PreDefinedVoucherSeries struct {
		Url           string `json:"@url"`
		Name          string `json:"Name"`
		VoucherSeries string `json:"VoucherSeries"`
	} `json:"PreDefinedVoucherSeries"`
}

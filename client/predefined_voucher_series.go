package client

import (
	"context"
	"fmt"
)

const (
	predefinedVoucherSeriesURI = "predefinedvoucherseries"
)

// GetAllPredefinedVoucherSeries does _GET https://api.fortnox.se/3/predefinedvoucherseries/
func (c *Client) GetAllPredefinedVoucherSeries(ctx context.Context) ([]PreDefinedVoucherSeries, error) {
	resp := &GetAllPredefinedVoucherSeriesResp{}

	err := c._GET(ctx, predefinedVoucherSeriesURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.PreDefinedVoucherSeriesCollection, nil
}

// GetPredefinedVoucherSeries does _GET https://api.fortnox.se/3/predefinedvoucherseries/{Name}
//
// name - identifies the predefined voucher series
func (c *Client) GetPredefinedVoucherSeries(ctx context.Context, name string) (*PreDefinedVoucherSeries, error) {
	resp := &GetPredefinedVoucherSeriesResp{}

	uri := fmt.Sprintf("%s/%s", predefinedVoucherSeriesURI, name)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.PreDefinedVoucherSeries, nil
}

// UpdatePredefinedVoucherSeries does _PUT https://api.fortnox.se/3/predefinedvoucherseries/{Name}
//
// name - identifies the predefined voucher series
//
// pdvs - predefined voucher series to update
func (c *Client) UpdatePredefinedVoucherSeries(
	ctx context.Context,
	name string,
	pdvs *PreDefinedVoucherSeries) (*PreDefinedVoucherSeries, error) {

	req := &UpdatePredefinedVoucherSeriesReq{PreDefinedVoucherSeries: *pdvs}
	resp := &UpdatePredefinedVoucherSeriesResp{}

	uri := fmt.Sprintf("%s/%s", predefinedVoucherSeriesURI, name)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.PreDefinedVoucherSeries, nil
}

type PreDefinedVoucherSeries struct {
	Url           string `json:"@url,omitempty"`
	Name          string `json:"Name,omitempty"`
	VoucherSeries string `json:"VoucherSeries,omitempty"`
}

type GetAllPredefinedVoucherSeriesResp struct {
	PreDefinedVoucherSeriesCollection []PreDefinedVoucherSeries `json:"PreDefinedVoucherSeriesCollection"`
}

type GetPredefinedVoucherSeriesResp struct {
	PreDefinedVoucherSeries PreDefinedVoucherSeries `json:"PreDefinedVoucherSeries"`
}

type UpdatePredefinedVoucherSeriesReq struct {
	PreDefinedVoucherSeries PreDefinedVoucherSeries `json:"PreDefinedVoucherSeries"`
}

type UpdatePredefinedVoucherSeriesResp struct {
	PreDefinedVoucherSeries PreDefinedVoucherSeries `json:"PreDefinedVoucherSeries"`
}

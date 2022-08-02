package client

import (
	"context"
	"fmt"
)

const (
	voucherSeriesURI = "voucherseries"
)

// GetAllVoucherSeries does _GET https://api.fortnox.se/3/voucherseries
func (c *Client) GetAllVoucherSeries(ctx context.Context) ([]VoucherSeries, error) {
	resp := &GetAllVoucherSeriesResp{}

	err := c._GET(ctx, voucherSeriesURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.VoucherSeriesCollection, nil
}

// CreateVoucherSeries does _POST https://api.fortnox.se/3/voucherseries
//
// req - voucher series to create
func (c *Client) CreateVoucherSeries(ctx context.Context, vs *VoucherSeries) (*VoucherSeries, error) {
	req := &CreateVoucherSeriesReq{VoucherSeries: *vs}
	resp := &CreateVoucherSeriesResp{}

	err := c._POST(ctx, vouchersURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.VoucherSeries, nil
}

// GetVoucherSeriesByCode doe _GET https://api.fortnox.se/3/voucherseries/{Code}
//
// code - identifies the voucher series
func (c *Client) GetVoucherSeriesByCode(ctx context.Context, code string) (*VoucherSeries, error) {
	resp := &GetVoucherSeriesByCodeResp{}

	uri := fmt.Sprintf("%s/%s", vouchersURI, code)
	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.VoucherSeries, nil
}

// UpdateVoucherSeries does _PUT https://api.fortnox.se/3/voucherseries/{Code}
//
// code - identifies the voucher series
//
// req - to update
func (c *Client) UpdateVoucherSeries(
	ctx context.Context,
	code string,
	vs *VoucherSeries) (*VoucherSeries, error) {

	req := &UpdateVoucherSeriesReq{VoucherSeries: *vs}
	resp := &UpdateVoucherSeriesResp{}

	uri := fmt.Sprintf("%s/%s", vouchersURI, code)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.VoucherSeries, nil
}

type VoucherSeries struct {
	Url               string   `json:"@url,omitempty"`
	Code              string   `json:"Code,omitempty"`
	Description       string   `json:"Description,omitempty"`
	Manual            bool     `json:"Manual,omitempty"`
	NextVoucherNumber int      `json:"NextVoucherNumber,omitempty"`
	Year              int      `json:"Year,omitempty"`
	Approver          Approver `json:"Approver,omitempty"`
}

type Approver struct {
	Id   int    `json:"Id,omitempty"`
	Name string `json:"Name,omitempty"`
}

type GetAllVoucherSeriesResp struct {
	VoucherSeriesCollection []VoucherSeries `json:"VoucherSeriesCollection"`
}

type CreateVoucherSeriesReq struct {
	VoucherSeries VoucherSeries `json:"VoucherSeries"`
}

type CreateVoucherSeriesResp struct {
	VoucherSeries VoucherSeries `json:"VoucherSeries"`
}

type GetVoucherSeriesByCodeResp struct {
	VoucherSeries VoucherSeries `json:"VoucherSeries"`
}

type UpdateVoucherSeriesReq struct {
	VoucherSeries VoucherSeries `json:"VoucherSeries"`
}

type UpdateVoucherSeriesResp struct {
	VoucherSeries VoucherSeries `json:"VoucherSeries"`
}

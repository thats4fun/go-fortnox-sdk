package client

import (
	"context"
	"fmt"
)

const (
	voucherSeriesURI = "voucherseries"
)

// GetAllVoucherSeries does _GET https://api.fortnox.se/3/voucherseries
func (c *Client) GetAllVoucherSeries(ctx context.Context) (*GetAllVoucherSeriesResp, error) {
	resp := &GetAllVoucherSeriesResp{}

	err := c._GET(ctx, voucherSeriesURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateVoucherSeries doe _POST https://api.fortnox.se/3/voucherseries
//
// req - voucher series to create
func (c *Client) CreateVoucherSeries(ctx context.Context, req *CreateVoucherSeriesReq) (*CreateVoucherSeriesResp, error) {
	resp := &CreateVoucherSeriesResp{}

	err := c._POST(ctx, vouchersURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetVoucherSeriesByCode doe _GET https://api.fortnox.se/3/voucherseries/{Code}
//
// code - identifies the voucher series
func (c *Client) GetVoucherSeriesByCode(ctx context.Context, code string) (*GetVoucherSeriesByCodeResp, error) {
	resp := &GetVoucherSeriesByCodeResp{}

	uri := fmt.Sprintf("%s/%s", vouchersURI, code)
	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateVoucherSeries does _PUT https://api.fortnox.se/3/voucherseries/{Code}
//
// code - identifies the voucher series
//
// req - to update
func (c *Client) UpdateVoucherSeries(
	ctx context.Context,
	code string,
	req *UpdateVoucherSeriesReq) (*UpdateVoucherSeriesResp, error) {

	resp := &UpdateVoucherSeriesResp{}

	uri := fmt.Sprintf("%s/%s", vouchersURI, code)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetAllVoucherSeriesResp struct {
	VoucherSeriesCollection []struct {
		Url         string `json:"@url"`
		Code        string `json:"Code"`
		Description string `json:"Description"`
		Manual      bool   `json:"Manual"`
		Year        int    `json:"Year"`
		Approver    struct {
			Id   int    `json:"Id"`
			Name string `json:"Name"`
		} `json:"Approver"`
	} `json:"VoucherSeriesCollection"`
}

type CreateVoucherSeriesReq struct {
	VoucherSeries struct {
		Url               string `json:"@url"`
		Code              string `json:"Code"`
		Description       string `json:"Description"`
		Manual            bool   `json:"Manual"`
		NextVoucherNumber int    `json:"NextVoucherNumber"`
		Year              int    `json:"Year"`
		Approver          struct {
			Id   int    `json:"Id"`
			Name string `json:"Name"`
		} `json:"Approver"`
	} `json:"VoucherSeries"`
}

type CreateVoucherSeriesResp struct {
	VoucherSeries struct {
		Url               string `json:"@url"`
		Code              string `json:"Code"`
		Description       string `json:"Description"`
		Manual            bool   `json:"Manual"`
		NextVoucherNumber int    `json:"NextVoucherNumber"`
		Year              int    `json:"Year"`
		Approver          struct {
			Id   int    `json:"Id"`
			Name string `json:"Name"`
		} `json:"Approver"`
	} `json:"VoucherSeries"`
}

type GetVoucherSeriesByCodeResp struct {
	VoucherSeries struct {
		Url               string `json:"@url"`
		Code              string `json:"Code"`
		Description       string `json:"Description"`
		Manual            bool   `json:"Manual"`
		NextVoucherNumber int    `json:"NextVoucherNumber"`
		Year              int    `json:"Year"`
		Approver          struct {
			Id   int    `json:"Id"`
			Name string `json:"Name"`
		} `json:"Approver"`
	} `json:"VoucherSeries"`
}

type UpdateVoucherSeriesReq struct {
	VoucherSeries struct {
		Url               string `json:"@url"`
		Code              string `json:"Code"`
		Description       string `json:"Description"`
		Manual            bool   `json:"Manual"`
		NextVoucherNumber int    `json:"NextVoucherNumber"`
		Year              int    `json:"Year"`
		Approver          struct {
			Id   int    `json:"Id"`
			Name string `json:"Name"`
		} `json:"Approver"`
	} `json:"VoucherSeries"`
}

type UpdateVoucherSeriesResp struct {
	VoucherSeries struct {
		Url               string `json:"@url"`
		Code              string `json:"Code"`
		Description       string `json:"Description"`
		Manual            bool   `json:"Manual"`
		NextVoucherNumber int    `json:"NextVoucherNumber"`
		Year              int    `json:"Year"`
		Approver          struct {
			Id   int    `json:"Id"`
			Name string `json:"Name"`
		} `json:"Approver"`
	} `json:"VoucherSeries"`
}

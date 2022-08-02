package client

import (
	"context"
	"fmt"
)

const (
	vouchersURI = "vouchers"
)

// GetVoucher does _GET https://api.fortnox.se/3/vouchers/{VoucherSeries}/{VoucherNumber}
//
// voucherSeries - identifies the voucher series
//
// voucherNumber - identifies the voucher number
//
// financialYear - filter on financial year
func (c *Client) GetVoucher(
	ctx context.Context,
	voucherSeries, voucherNumber string,
	filter *FinancialYearFilter) (*GetVoucherResp, error) {

	resp := &GetVoucherResp{}

	uri := fmt.Sprintf("%s/%s/%s", vouchersURI, voucherSeries, voucherNumber)

	params := filter.urlValues()

	err := c._GET(ctx, uri, params, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetAllVouchers does _GET https://api.fortnox.se/3/vouchers/
//
// filter - filter on financial year
func (c *Client) GetAllVouchers(ctx context.Context, filter *FinancialYearFilter) (*GetAllVouchersResp, error) {
	resp := &GetAllVouchersResp{}

	params := filter.urlValues()

	err := c._GET(ctx, vouchersURI, params, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateVoucher does _POST https://api.fortnox.se/3/vouchers/
//
// filter - financial year id, used to determine which financial year the voucher is created in
//
// req - voucher to create
func (c *Client) CreateVoucher(
	ctx context.Context,
	filter FinancialYearFilter,
	req *CreateVoucherReq) (*CreateVoucherResp, error) {

	resp := &CreateVoucherResp{}

	params := filter.urlValues()

	err := c._POST(ctx, vouchersURI, params, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetVouchersBySeries does _GET https://api.fortnox.se/3/vouchers/sublist/{VoucherSeries}
//
// voucherSeries - identifies the voucher series
//
// filter - filter on financial year
func (c *Client) GetVouchersBySeries(
	ctx context.Context,
	voucherSeries string,
	filter FinancialYearFilter) (*GetVouchersBySeriesResp, error) {

	resp := &GetVouchersBySeriesResp{}

	uri := fmt.Sprintf("%s/sublist/%s", vouchersURI, voucherSeries)

	params := filter.urlValues()

	err := c._GET(ctx, uri, params, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetVoucherResp struct {
	Voucher Voucher `json:"Voucher"`
}

type GetAllVouchersResp struct {
	Vouchers []Voucher `json:"Vouchers"`
}

type CreateVoucherReq struct {
	Voucher Voucher `json:"Voucher"`
}

type CreateVoucherResp struct {
	Voucher Voucher `json:"Voucher"`
}

type GetVouchersBySeriesResp struct {
	Vouchers []Voucher `json:"Vouchers"`
}

type VoucherRow struct {
	Account                int    `json:"Account,omitempty"`
	CostCenter             string `json:"CostCenter,omitempty"`
	Credit                 int    `json:"Credit,omitempty"`
	Description            string `json:"Description,omitempty"`
	Debit                  int    `json:"Debit,omitempty"`
	Project                string `json:"Project,omitempty"`
	Removed                bool   `json:"Removed,omitempty"`
	TransactionInformation string `json:"TransactionInformation,omitempty"`
	Quantity               int    `json:"Quantity,omitempty"`
}

type Voucher struct {
	Url             string       `json:"@url,omitempty"`
	Comments        string       `json:"Comments,omitempty"`
	CostCenter      string       `json:"CostCenter,omitempty"`
	Description     string       `json:"Description,omitempty"`
	Project         string       `json:"Project,omitempty"`
	ReferenceNumber string       `json:"ReferenceNumber,omitempty"`
	ReferenceType   string       `json:"ReferenceType,omitempty"`
	TransactionDate string       `json:"TransactionDate,omitempty"`
	VoucherNumber   int          `json:"VoucherNumber,omitempty"`
	VoucherRows     []VoucherRow `json:"VoucherRows,omitempty"`
	VoucherSeries   string       `json:"VoucherSeries,omitempty"`
	Year            int          `json:"Year,omitempty"`
	ApprovalState   int          `json:"ApprovalState,omitempty"`
}

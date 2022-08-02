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
	Voucher struct {
		Url             string `json:"@url"`
		Comments        string `json:"Comments"`
		CostCenter      string `json:"CostCenter"`
		Description     string `json:"Description"`
		Project         string `json:"Project"`
		ReferenceNumber string `json:"ReferenceNumber"`
		ReferenceType   string `json:"ReferenceType"`
		TransactionDate string `json:"TransactionDate"`
		VoucherNumber   int    `json:"VoucherNumber"`
		VoucherRows     []struct {
			Account                int    `json:"Account"`
			CostCenter             string `json:"CostCenter"`
			Credit                 int    `json:"Credit"`
			Description            string `json:"Description"`
			Debit                  int    `json:"Debit"`
			Project                string `json:"Project"`
			Removed                bool   `json:"Removed"`
			TransactionInformation string `json:"TransactionInformation"`
			Quantity               int    `json:"Quantity"`
		} `json:"VoucherRows"`
		VoucherSeries string `json:"VoucherSeries"`
		Year          int    `json:"Year"`
		ApprovalState int    `json:"ApprovalState"`
	} `json:"Voucher"`
}

type GetAllVouchersResp struct {
	Vouchers []struct {
		Url             string `json:"@url"`
		Comments        string `json:"Comments"`
		Description     string `json:"Description"`
		ReferenceNumber string `json:"ReferenceNumber"`
		ReferenceType   string `json:"ReferenceType"`
		TransactionDate string `json:"TransactionDate"`
		VoucherNumber   int    `json:"VoucherNumber"`
		VoucherSeries   string `json:"VoucherSeries"`
		Year            int    `json:"Year"`
		ApprovalState   int    `json:"ApprovalState"`
	} `json:"Vouchers"`
}

type CreateVoucherReq struct {
	Voucher struct {
		Url             string `json:"@url"`
		Comments        string `json:"Comments"`
		CostCenter      string `json:"CostCenter"`
		Description     string `json:"Description"`
		Project         string `json:"Project"`
		ReferenceNumber string `json:"ReferenceNumber"`
		ReferenceType   string `json:"ReferenceType"`
		TransactionDate string `json:"TransactionDate"`
		VoucherNumber   int    `json:"VoucherNumber"`
		VoucherRows     []struct {
			Account                int    `json:"Account"`
			CostCenter             string `json:"CostCenter"`
			Credit                 int    `json:"Credit"`
			Description            string `json:"Description"`
			Debit                  int    `json:"Debit"`
			Project                string `json:"Project"`
			Removed                bool   `json:"Removed"`
			TransactionInformation string `json:"TransactionInformation"`
			Quantity               int    `json:"Quantity"`
		} `json:"VoucherRows"`
		VoucherSeries string `json:"VoucherSeries"`
		Year          int    `json:"Year"`
		ApprovalState int    `json:"ApprovalState"`
	} `json:"Voucher"`
}

type CreateVoucherResp struct {
	Voucher struct {
		Url             string `json:"@url"`
		Comments        string `json:"Comments"`
		CostCenter      string `json:"CostCenter"`
		Description     string `json:"Description"`
		Project         string `json:"Project"`
		ReferenceNumber string `json:"ReferenceNumber"`
		ReferenceType   string `json:"ReferenceType"`
		TransactionDate string `json:"TransactionDate"`
		VoucherNumber   int    `json:"VoucherNumber"`
		VoucherRows     []struct {
			Account                int    `json:"Account"`
			CostCenter             string `json:"CostCenter"`
			Credit                 int    `json:"Credit"`
			Description            string `json:"Description"`
			Debit                  int    `json:"Debit"`
			Project                string `json:"Project"`
			Removed                bool   `json:"Removed"`
			TransactionInformation string `json:"TransactionInformation"`
			Quantity               int    `json:"Quantity"`
		} `json:"VoucherRows"`
		VoucherSeries string `json:"VoucherSeries"`
		Year          int    `json:"Year"`
		ApprovalState int    `json:"ApprovalState"`
	} `json:"Voucher"`
}

type GetVouchersBySeriesResp struct {
	Vouchers []struct {
		Url             string `json:"@url"`
		Comments        string `json:"Comments"`
		Description     string `json:"Description"`
		ReferenceNumber string `json:"ReferenceNumber"`
		ReferenceType   string `json:"ReferenceType"`
		TransactionDate string `json:"TransactionDate"`
		VoucherNumber   int    `json:"VoucherNumber"`
		VoucherSeries   string `json:"VoucherSeries"`
		Year            int    `json:"Year"`
		ApprovalState   int    `json:"ApprovalState"`
	} `json:"Vouchers"`
}

type VoucherRow struct {
	Account                int    `json:"Account"`
	CostCenter             string `json:"CostCenter"`
	Credit                 int    `json:"Credit"`
	Description            string `json:"Description"`
	Debit                  int    `json:"Debit"`
	Project                string `json:"Project"`
	Removed                bool   `json:"Removed"`
	TransactionInformation string `json:"TransactionInformation"`
	Quantity               int    `json:"Quantity"`
}

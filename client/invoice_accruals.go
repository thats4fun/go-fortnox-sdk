package client

import (
	"context"
	"fmt"
)

const (
	invoiceAccrualsURI = "invoiceaccruals"
)

// GetAllInvoiceAccruals does _GET https://api.fortnox.se/3/invoiceaccruals/
func (c *Client) GetAllInvoiceAccruals(ctx context.Context) (*GetAllInvoiceAccrualsResp, error) {
	resp := &GetAllInvoiceAccrualsResp{}

	err := c._GET(ctx, invoiceAccrualsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateInvoiceAccrual does _POST https://api.fortnox.se/3/invoiceaccruals/
//
// req - invoice accrual to create
func (c *Client) CreateInvoiceAccrual(
	ctx context.Context, req *CreateInvoiceAccrualReq) (*CreateInvoiceAccrualResp, error) {

	resp := &CreateInvoiceAccrualResp{}

	err := c._POST(ctx, invoiceAccrualsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetInvoiceAccrual does _GET https://api.fortnox.se/3/invoiceaccruals/{InvoiceNumber}
//
// invoiceNumber - identifies the invoice accrual
func (c *Client) GetInvoiceAccrual(ctx context.Context, invoiceNumber int) (*GetInvoiceAccrualResp, error) {
	resp := &GetInvoiceAccrualResp{}

	uri := fmt.Sprintf("%s/%d", invoiceAccrualsURI, invoiceNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateInvoiceAccrual does _PUT https://api.fortnox.se/3/invoiceaccruals/{InvoiceNumber}
//
// invoiceNumber - identifies the invoice accrual
//
// req - invoice accrual to update
func (c *Client) UpdateInvoiceAccrual(
	ctx context.Context,
	invoiceNumber int,
	req *UpdateInvoiceAccrualReq) (*UpdateInvoiceAccrualResp, error) {

	resp := &UpdateInvoiceAccrualResp{}

	uri := fmt.Sprintf("%s/%d", invoiceAccrualsURI, invoiceNumber)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// RemoveInvoiceAccrual does _DELETE https://api.fortnox.se/3/invoiceaccruals/{InvoiceNumber}
//
// invoiceNumber - identifies the invoice accrual
func (c *Client) RemoveInvoiceAccrual(ctx context.Context, invoiceNumber int) error {
	uri := fmt.Sprintf("%s/%d", invoiceAccrualsURI, invoiceNumber)
	return c._DELETE(ctx, uri)
}

type GetAllInvoiceAccrualsResp struct {
	InvoiceAccruals []struct {
		Url           string `json:"@url"`
		Description   string `json:"Description"`
		InvoiceNumber int    `json:"InvoiceNumber"`
		Period        string `json:"Period"`
	} `json:"InvoiceAccruals"`
}

type CreateInvoiceAccrualReq struct {
	InvoiceAccrual struct {
		Url                string `json:"@url"`
		AccrualAccount     int    `json:"AccrualAccount"`
		Description        string `json:"Description"`
		EndDate            string `json:"EndDate"`
		InvoiceAccrualRows []struct {
			Account                int    `json:"Account"`
			CostCenter             string `json:"CostCenter"`
			Credit                 int    `json:"Credit"`
			Debit                  int    `json:"Debit"`
			Project                string `json:"Project"`
			TransactionInformation string `json:"TransactionInformation"`
		} `json:"InvoiceAccrualRows"`
		InvoiceNumber  int    `json:"InvoiceNumber"`
		Period         string `json:"Period"`
		RevenueAccount int    `json:"RevenueAccount"`
		StartDate      string `json:"StartDate"`
		Times          int    `json:"Times"`
		Total          int    `json:"Total"`
		VATIncluded    bool   `json:"VATIncluded"`
	} `json:"InvoiceAccrual"`
}

type CreateInvoiceAccrualResp struct {
	InvoiceAccrual struct {
		Url                string `json:"@url"`
		AccrualAccount     int    `json:"AccrualAccount"`
		Description        string `json:"Description"`
		EndDate            string `json:"EndDate"`
		InvoiceAccrualRows []struct {
			Account                int    `json:"Account"`
			CostCenter             string `json:"CostCenter"`
			Credit                 int    `json:"Credit"`
			Debit                  int    `json:"Debit"`
			Project                string `json:"Project"`
			TransactionInformation string `json:"TransactionInformation"`
		} `json:"InvoiceAccrualRows"`
		InvoiceNumber  int    `json:"InvoiceNumber"`
		Period         string `json:"Period"`
		RevenueAccount int    `json:"RevenueAccount"`
		StartDate      string `json:"StartDate"`
		Times          int    `json:"Times"`
		Total          int    `json:"Total"`
		VATIncluded    bool   `json:"VATIncluded"`
	} `json:"InvoiceAccrual"`
}

type GetInvoiceAccrualResp struct {
	InvoiceAccrual struct {
		Url                string `json:"@url"`
		AccrualAccount     int    `json:"AccrualAccount"`
		Description        string `json:"Description"`
		EndDate            string `json:"EndDate"`
		InvoiceAccrualRows []struct {
			Account                int    `json:"Account"`
			CostCenter             string `json:"CostCenter"`
			Credit                 int    `json:"Credit"`
			Debit                  int    `json:"Debit"`
			Project                string `json:"Project"`
			TransactionInformation string `json:"TransactionInformation"`
		} `json:"InvoiceAccrualRows"`
		InvoiceNumber  int    `json:"InvoiceNumber"`
		Period         string `json:"Period"`
		RevenueAccount int    `json:"RevenueAccount"`
		StartDate      string `json:"StartDate"`
		Times          int    `json:"Times"`
		Total          int    `json:"Total"`
		VATIncluded    bool   `json:"VATIncluded"`
	} `json:"InvoiceAccrual"`
}

type UpdateInvoiceAccrualReq struct {
	InvoiceAccrual struct {
		Url                string `json:"@url"`
		AccrualAccount     int    `json:"AccrualAccount"`
		Description        string `json:"Description"`
		EndDate            string `json:"EndDate"`
		InvoiceAccrualRows []struct {
			Account                int    `json:"Account"`
			CostCenter             string `json:"CostCenter"`
			Credit                 int    `json:"Credit"`
			Debit                  int    `json:"Debit"`
			Project                string `json:"Project"`
			TransactionInformation string `json:"TransactionInformation"`
		} `json:"InvoiceAccrualRows"`
		InvoiceNumber  int    `json:"InvoiceNumber"`
		Period         string `json:"Period"`
		RevenueAccount int    `json:"RevenueAccount"`
		StartDate      string `json:"StartDate"`
		Times          int    `json:"Times"`
		Total          int    `json:"Total"`
		VATIncluded    bool   `json:"VATIncluded"`
	} `json:"InvoiceAccrual"`
}

type UpdateInvoiceAccrualResp struct {
	InvoiceAccrual struct {
		Url                string `json:"@url"`
		AccrualAccount     int    `json:"AccrualAccount"`
		Description        string `json:"Description"`
		EndDate            string `json:"EndDate"`
		InvoiceAccrualRows []struct {
			Account                int    `json:"Account"`
			CostCenter             string `json:"CostCenter"`
			Credit                 int    `json:"Credit"`
			Debit                  int    `json:"Debit"`
			Project                string `json:"Project"`
			TransactionInformation string `json:"TransactionInformation"`
		} `json:"InvoiceAccrualRows"`
		InvoiceNumber  int    `json:"InvoiceNumber"`
		Period         string `json:"Period"`
		RevenueAccount int    `json:"RevenueAccount"`
		StartDate      string `json:"StartDate"`
		Times          int    `json:"Times"`
		Total          int    `json:"Total"`
		VATIncluded    bool   `json:"VATIncluded"`
	} `json:"InvoiceAccrual"`
}

package client

import (
	"context"
	"fmt"
)

const (
	supplierInvoiceAccrualsURI = "supplierinvoiceaccruals"
)

// GetAllSupplierInvoiceAccruals does _GET https://api.fortnox.se/3/supplierinvoiceaccruals/
func (c *Client) GetAllSupplierInvoiceAccruals(ctx context.Context) ([]SupplierInvoiceAccrual, error) {
	resp := &GetAllSupplierInvoiceAccrualsResp{}

	err := c._GET(ctx, supplierInvoiceAccrualsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.SupplierInvoiceAccruals, nil
}

// CreateSupplierInvoiceAccruals does _POST https://api.fortnox.se/3/supplierinvoiceaccruals/
//
// req - supplier invoice accruals to create
func (c *Client) CreateSupplierInvoiceAccruals(
	ctx context.Context,
	sia *SupplierInvoiceAccrual) (*SupplierInvoiceAccrual, error) {

	req := &CreateSupplierInvoiceAccrualsReq{SupplierInvoiceAccrual: *sia}
	resp := &CreateSupplierInvoiceAccrualsResp{}

	err := c._POST(ctx, supplierInvoiceAccrualsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.SupplierInvoiceAccrual, nil
}

// GetSupplierInvoiceAccruals does _GET https://api.fortnox.se/3/supplierinvoiceaccruals/{SupplierInvoiceNumber}
//
// supplierInvoiceNumber - identifies the supplier invoice accrual
func (c *Client) GetSupplierInvoiceAccruals(
	ctx context.Context,
	supplierInvoiceNumber int) (*SupplierInvoiceAccrual, error) {

	resp := &GetSupplierInvoiceAccrualsResp{}

	uri := fmt.Sprintf("%s/%d", supplierInvoiceAccrualsURI, supplierInvoiceNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.SupplierInvoiceAccrual, nil
}

// UpdateSupplierInvoiceAccruals does _PUT https://api.fortnox.se/3/supplierinvoiceaccruals/{SupplierInvoiceNumber}
//
// supplierInvoiceNumber - identifies the supplier invoice accrual
//
// req - supplier invoice accruals to update
func (c *Client) UpdateSupplierInvoiceAccruals(
	ctx context.Context,
	supplierInvoiceNumber int,
	sia *SupplierInvoiceAccrual) (*SupplierInvoiceAccrual, error) {

	req := UpdateSupplierInvoiceAccrualReq{SupplierInvoiceAccrual: *sia}
	resp := &UpdateSupplierInvoiceAccrualResp{}

	uri := fmt.Sprintf("%s/%d", supplierInvoiceAccrualsURI, supplierInvoiceNumber)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.SupplierInvoiceAccrual, nil
}

// DeleteSupplierInvoiceAccruals does _DELETE https://api.fortnox.se/3/supplierinvoiceaccruals/{SupplierInvoiceNumber}
//
// supplierInvoiceNumber - identifies the supplier invoice accrual
func (c *Client) DeleteSupplierInvoiceAccruals(ctx context.Context, supplierInvoiceNumber int) error {
	uri := fmt.Sprintf("%s/%d", supplierInvoiceAccrualsURI, supplierInvoiceNumber)
	return c._DELETE(ctx, uri)
}

type GetAllSupplierInvoiceAccrualsResp struct {
	SupplierInvoiceAccruals []SupplierInvoiceAccrual `json:"SupplierInvoiceAccruals"`
}

type CreateSupplierInvoiceAccrualsReq struct {
	SupplierInvoiceAccrual SupplierInvoiceAccrual `json:"SupplierInvoiceAccrual"`
}

type CreateSupplierInvoiceAccrualsResp struct {
	SupplierInvoiceAccrual SupplierInvoiceAccrual `json:"SupplierInvoiceAccrual"`
}

type GetSupplierInvoiceAccrualsResp struct {
	SupplierInvoiceAccrual SupplierInvoiceAccrual `json:"SupplierInvoiceAccrual"`
}

type UpdateSupplierInvoiceAccrualReq struct {
	SupplierInvoiceAccrual SupplierInvoiceAccrual `json:"SupplierInvoiceAccrual"`
}

type UpdateSupplierInvoiceAccrualResp struct {
	SupplierInvoiceAccrual SupplierInvoiceAccrual `json:"SupplierInvoiceAccrual"`
}

type SupplierInvoiceAccrualRow struct {
	Account                int    `json:"Account,omitempty"`
	CostCenter             string `json:"CostCenter,omitempty"`
	Credit                 int    `json:"Credit,omitempty"`
	Debit                  int    `json:"Debit,omitempty"`
	Project                string `json:"Project,omitempty"`
	TransactionInformation string `json:"TransactionInformation,omitempty"`
}

type SupplierInvoiceAccrual struct {
	Url                        string                      `json:"@url,omitempty"`
	AccrualAccount             int                         `json:"AccrualAccount,omitempty"`
	CostAccount                int                         `json:"CostAccount,omitempty"`
	Description                string                      `json:"Description,omitempty"`
	EndDate                    string                      `json:"EndDate,omitempty"`
	SupplierInvoiceNumber      int                         `json:"SupplierInvoiceNumber,omitempty"`
	Period                     string                      `json:"Period,omitempty"`
	StartDate                  string                      `json:"StartDate,omitempty"`
	Times                      int                         `json:"Times,omitempty"`
	Total                      int                         `json:"Total,omitempty"`
	VATIncluded                bool                        `json:"VATIncluded,omitempty"`
	SupplierInvoiceAccrualRows []SupplierInvoiceAccrualRow `json:"SupplierInvoiceAccrualRows,omitempty"`
}

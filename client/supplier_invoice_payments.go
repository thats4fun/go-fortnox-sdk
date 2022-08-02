package client

import (
	"context"
	"fmt"
)

const (
	supplierInvoicePaymentsURI = "supplierinvoicepayments"
)

// GetAllSupplierInvoicePayments does _GET https://api.fortnox.se/3/supplierinvoicepayments/
func (c *Client) GetAllSupplierInvoicePayments(ctx context.Context) (*GetAllSupplierInvoicePaymentsResp, error) {
	resp := &GetAllSupplierInvoicePaymentsResp{}

	err := c._GET(ctx, supplierInvoicePaymentsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateSupplierInvoicePayment does _POST https://api.fortnox.se/3/supplierinvoicepayments/
//
// req - supplier invoice payment to create
func (c *Client) CreateSupplierInvoicePayment(
	ctx context.Context,
	req *CreateSupplierInvoicePaymentReq) (*CreateSupplierInvoicePaymentResp, error) {

	resp := &CreateSupplierInvoicePaymentResp{}

	err := c._POST(ctx, supplierInvoicePaymentsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetSupplierInvoicePayment does _GET https://api.fortnox.se/3/supplierinvoicepayments/{Number}
//
// number - identifies the supplier invoice payment
func (c *Client) GetSupplierInvoicePayment(ctx context.Context, number int) (*GetSupplierInvoicePaymentResp, error) {
	resp := &GetSupplierInvoicePaymentResp{}

	uri := fmt.Sprintf("%s/%d", supplierInvoicePaymentsURI, number)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateSupplierInvoicePayment does _PUT https://api.fortnox.se/3/supplierinvoicepayments/{Number}
//
// number - identifies the supplier invoice payment
//
// req - supplier invoice payment to update
func (c *Client) UpdateSupplierInvoicePayment(
	ctx context.Context,
	number int,
	req *UpdateSupplierInvoicePaymentReq) (*UpdateSupplierInvoicePaymentResp, error) {

	resp := &UpdateSupplierInvoicePaymentResp{}

	uri := fmt.Sprintf("%s/%d", supplierInvoicePaymentsURI, number)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// RemoveSupplierInvoicePayment does _DELETE https://api.fortnox.se/3/supplierinvoicepayments/{Number}
//
// number - identifies the supplier invoice payment
func (c *Client) RemoveSupplierInvoicePayment(ctx context.Context, number int) error {
	uri := fmt.Sprintf("%s/%d", supplierInvoicePaymentsURI, number)
	return c._DELETE(ctx, uri)
}

// BookKeepSupplierInvoicePayment does _PUT https://api.fortnox.se/3/supplierinvoicepayments/{Number}/bookkeep
//
// number - identifies the supplier invoice payment
func (c *Client) BookKeepSupplierInvoicePayment(
	ctx context.Context,
	number int) (*BookKeepSupplierInvoicePaymentResp, error) {

	resp := &BookKeepSupplierInvoicePaymentResp{}

	uri := fmt.Sprintf("%s/%d/bookkeep", supplierInvoicePaymentsURI, number)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetAllSupplierInvoicePaymentsResp struct {
	SupplierInvoicePayments []SupplierInvoicePayment `json:"SupplierInvoicePayments"`
}

type CreateSupplierInvoicePaymentReq struct {
	SupplierInvoicePayment SupplierInvoicePayment `json:"SupplierInvoicePayment"`
}

type CreateSupplierInvoicePaymentResp struct {
	SupplierInvoicePayment SupplierInvoicePayment `json:"SupplierInvoicePayment"`
}

type GetSupplierInvoicePaymentResp struct {
	SupplierInvoicePayment SupplierInvoicePayment `json:"SupplierInvoicePayment"`
}

type UpdateSupplierInvoicePaymentReq struct {
	SupplierInvoicePayment SupplierInvoicePayment `json:"SupplierInvoicePayment"`
}

type UpdateSupplierInvoicePaymentResp struct {
	SupplierInvoicePayment SupplierInvoicePayment `json:"SupplierInvoicePayment"`
}

type BookKeepSupplierInvoicePaymentResp struct {
	SupplierInvoicePayment SupplierInvoicePayment `json:"SupplierInvoicePayment"`
}

type SupplierInvoicePayment struct {
	Url                   string     `json:"@url,omitempty"`
	Amount                int        `json:"Amount,omitempty"`
	AmountCurrency        int        `json:"AmountCurrency,omitempty"`
	Booked                bool       `json:"Booked,omitempty"`
	Currency              string     `json:"Currency,omitempty"`
	CurrencyRate          int        `json:"CurrencyRate,omitempty"`
	CurrencyUnit          int        `json:"CurrencyUnit,omitempty"`
	Information           string     `json:"Information,omitempty"`
	InvoiceNumber         string     `json:"InvoiceNumber,omitempty"`
	InvoiceDueDate        string     `json:"InvoiceDueDate,omitempty"`
	InvoiceOCR            string     `json:"InvoiceOCR,omitempty"`
	InvoiceSupplierName   string     `json:"InvoiceSupplierName,omitempty"`
	InvoiceSupplierNumber string     `json:"InvoiceSupplierNumber,omitempty"`
	InvoiceTotal          string     `json:"InvoiceTotal,omitempty"`
	ModeOfPayment         string     `json:"ModeOfPayment,omitempty"`
	Number                int        `json:"Number,omitempty"`
	PaymentDate           string     `json:"PaymentDate,omitempty"`
	Source                string     `json:"Source,omitempty"`
	VoucherNumber         int        `json:"VoucherNumber,omitempty"`
	VoucherSeries         string     `json:"VoucherSeries,omitempty"`
	VoucherYear           int        `json:"VoucherYear,omitempty"`
	WriteOffs             []WriteOff `json:"WriteOffs,omitempty"`
}

type WriteOff struct {
	Amount                 int    `json:"Amount,omitempty"`
	AccountNumber          int    `json:"AccountNumber,omitempty"`
	CostCenter             string `json:"CostCenter,omitempty"`
	Currency               string `json:"Currency,omitempty"`
	Description            string `json:"Description,omitempty"`
	TransactionInformation string `json:"TransactionInformation,omitempty"`
	Project                string `json:"Project,omitempty"`
}

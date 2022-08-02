package client

import (
	"context"
	"fmt"
)

const (
	invoicePaymentsURI = "invoicepayments"
)

// GetAllInvoicePayments does _GET https://api.fortnox.se/3/invoicepayments/
func (c *Client) GetAllInvoicePayments(ctx context.Context) ([]InvoicePayment, error) {
	resp := &GetAllInvoicePaymentsResp{}

	err := c._GET(ctx, invoicePaymentsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.InvoicePayments, nil
}

// CreateInvoicePayment does _POST https://api.fortnox.se/3/invoicepayments/
//
// ip - invoice payment to create
func (c *Client) CreateInvoicePayment(ctx context.Context, ip *InvoicePayment) (*InvoicePayment, error) {
	req := &CreateInvoicePaymentReq{InvoicePayment: *ip}
	resp := &CreateInvoicePaymentResp{}

	err := c._POST(ctx, invoicePaymentsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.InvoicePayment, nil
}

// GetInvoicePayment does _GET https://api.fortnox.se/3/invoicepayments/{Number}
//
// number - identifies the invoice payment
func (c *Client) GetInvoicePayment(ctx context.Context, number string) (*InvoicePayment, error) {
	resp := &GetInvoicePaymentResp{}

	uri := fmt.Sprintf("%s/%s", invoicePaymentsURI, number)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.InvoicePayment, nil
}

// UpdateInvoicePayment does _PUT https://api.fortnox.se/3/invoicepayments/{Number}
//
// number - identifies the invoice payment
//
// ip - invoice payment to update
func (c *Client) UpdateInvoicePayment(ctx context.Context, number string, ip *InvoicePayment) (*InvoicePayment, error) {
	req := &UpdateInvoicePaymentReq{InvoicePayment: *ip}
	resp := &UpdateInvoicePaymentResp{}

	uri := fmt.Sprintf("%s/%s", invoicePaymentsURI, number)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.InvoicePayment, nil
}

// RemoveInvoicePayment does _DELETE https://api.fortnox.se/3/invoicepayments/{Number}
//
// number - identifies the invoice payment
func (c *Client) RemoveInvoicePayment(ctx context.Context, number string) error {
	uri := fmt.Sprintf("%s/%s", invoicePaymentsURI, number)
	return c._DELETE(ctx, uri)
}

// BookKeepInvoicePayment does _PUT https://api.fortnox.se/3/invoicepayments/{Number}/bookkeep
//
// number - identifies the invoice payment
func (c *Client) BookKeepInvoicePayment(
	ctx context.Context,
	number string,
	ip *InvoicePayment) (*InvoicePayment, error) {

	req := BookKeepInvoicePaymentReq{InvoicePayment: *ip}
	resp := &BookKeepInvoicePaymentResp{}

	uri := fmt.Sprintf("%s/%s/bookkeep", invoicePaymentsURI, number)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.InvoicePayment, nil
}

type GetAllInvoicePaymentsResp struct {
	InvoicePayments []InvoicePayment `json:"InvoicePayments"`
}

type CreateInvoicePaymentReq struct {
	InvoicePayment InvoicePayment `json:"InvoicePayment"`
}

type CreateInvoicePaymentResp struct {
	InvoicePayment InvoicePayment `json:"InvoicePayment"`
}

type BookKeepInvoicePaymentReq struct {
	InvoicePayment InvoicePayment `json:"InvoicePayment"`
}

type BookKeepInvoicePaymentResp struct {
	InvoicePayment InvoicePayment `json:"InvoicePayment"`
}

type GetInvoicePaymentResp struct {
	InvoicePayment InvoicePayment `json:"InvoicePayment"`
}

type UpdateInvoicePaymentReq struct {
	InvoicePayment InvoicePayment `json:"InvoicePayment"`
}

type UpdateInvoicePaymentResp struct {
	InvoicePayment InvoicePayment `json:"InvoicePayment"`
}

type InvoicePayment struct {
	Url                       string     `json:"@url,omitempty"`
	Amount                    int        `json:"Amount,omitempty"`
	AmountCurrency            int        `json:"AmountCurrency,omitempty"`
	Booked                    bool       `json:"Booked,omitempty"`
	Currency                  string     `json:"Currency,omitempty"`
	CurrencyRate              int        `json:"CurrencyRate,omitempty"`
	CurrencyUnit              int        `json:"CurrencyUnit,omitempty"`
	ExternalInvoiceReference1 string     `json:"ExternalInvoiceReference1,omitempty"`
	ExternalInvoiceReference2 string     `json:"ExternalInvoiceReference2,omitempty"`
	InvoiceCustomerName       string     `json:"InvoiceCustomerName,omitempty"`
	InvoiceCustomerNumber     string     `json:"InvoiceCustomerNumber,omitempty"`
	InvoiceNumber             int        `json:"InvoiceNumber,omitempty"`
	InvoiceDueDate            string     `json:"InvoiceDueDate,omitempty"`
	InvoiceOCR                string     `json:"InvoiceOCR,omitempty"`
	InvoiceTotal              string     `json:"InvoiceTotal,omitempty"`
	ModeOfPayment             string     `json:"ModeOfPayment,omitempty"`
	ModeOfPaymentAccount      int        `json:"ModeOfPaymentAccount,omitempty"`
	Number                    string     `json:"Number,omitempty"`
	PaymentDate               string     `json:"PaymentDate,omitempty"`
	VoucherNumber             int        `json:"VoucherNumber,omitempty"`
	VoucherSeries             string     `json:"VoucherSeries,omitempty"`
	VoucherYear               int        `json:"VoucherYear,omitempty"`
	Source                    string     `json:"Source,omitempty"`
	WriteOffs                 []WriteOff `json:"WriteOffs,omitempty"`
}

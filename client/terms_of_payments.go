package client

import (
	"context"
	"fmt"
)

const (
	termsOfPaymentsURI = "termsofpayments"
)

// GetAllTermsOfPayments does _GET https://api.fortnox.se/3/termsofpayments
func (c *Client) GetAllTermsOfPayments(ctx context.Context) ([]TermsOfPayment, error) {
	resp := &GetAllTermsOfPaymentsResp{}

	err := c._GET(ctx, termsOfPaymentsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.TermsOfPayments, nil
}

// CreateTermOfPayment does _POST https://api.fortnox.se/3/termsofpayments
//
// req - term of payment to create
func (c Client) CreateTermOfPayment(ctx context.Context, top *TermsOfPayment) (*TermsOfPayment, error) {
	req := &CreateTermOfPaymentReq{TermsOfPayment: *top}
	resp := &CreateTermOfPaymentResp{}

	err := c._POST(ctx, termsOfPaymentsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.TermsOfPayment, nil
}

// GetTermOfPayment does _GET https://api.fortnox.se/3/termsofpayments/{Code}
//
// code - identifies the terms of payment
func (c *Client) GetTermOfPayment(ctx context.Context, code string) (*TermsOfPayment, error) {
	resp := &GetTermOfPaymentResp{}

	uri := fmt.Sprintf("%s/%s", termsOfPaymentsURI, code)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.TermsOfPayment, nil
}

// UpdateTermOfPayment does _PUT https://api.fortnox.se/3/termsofpayments/{Code}
//
// code - identifies the terms of payment
//
// req - term of payment to update
func (c *Client) UpdateTermOfPayment(
	ctx context.Context,
	code string,
	top *TermsOfPayment) (*TermsOfPayment, error) {

	req := &UpdateTermOfPaymentReq{TermsOfPayment: *top}
	resp := &UpdateTermOfPaymentResp{}

	uri := fmt.Sprintf("%s/%s", termsOfPaymentsURI, code)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.TermsOfPayment, nil
}

// RemoveTermOfPayment does _DELETE https://api.fortnox.se/3/termsofpayments/{Code}
// code - identifies the terms of payment
func (c *Client) RemoveTermOfPayment(ctx context.Context, code string) error {
	uri := fmt.Sprintf("%s/%s", termsOfPaymentsURI, code)
	return c._DELETE(ctx, uri)
}

type TermsOfPayment struct {
	Url         string `json:"@url,omitempty"`
	Code        string `json:"Code,omitempty"`
	Description string `json:"Description,omitempty"`
}

type GetAllTermsOfPaymentsResp struct {
	TermsOfPayments []TermsOfPayment `json:"TermsOfPayments"`
}

type CreateTermOfPaymentReq struct {
	TermsOfPayment TermsOfPayment `json:"TermsOfPayment"`
}

type CreateTermOfPaymentResp struct {
	TermsOfPayment TermsOfPayment `json:"TermsOfPayment"`
}

type GetTermOfPaymentResp struct {
	TermsOfPayment TermsOfPayment `json:"TermsOfPayment"`
}

type UpdateTermOfPaymentReq struct {
	TermsOfPayment TermsOfPayment `json:"TermsOfPayment"`
}

type UpdateTermOfPaymentResp struct {
	TermsOfPayment TermsOfPayment `json:"TermsOfPayment"`
}

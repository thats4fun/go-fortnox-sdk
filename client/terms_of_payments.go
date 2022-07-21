package client

import (
	"context"
	"fmt"
)

const (
	termsOfPaymentsURI = "termsofpayments"
)

// GetAllTermsOfPayments does _GET https://api.fortnox.se/3/termsofpayments
func (c *Client) GetAllTermsOfPayments(ctx context.Context) (*GetAllTermsOfPaymentsResp, error) {
	resp := &GetAllTermsOfPaymentsResp{}

	err := c._GET(ctx, termsOfPaymentsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

//CreateTermOfPayment does _POST https://api.fortnox.se/3/termsofpayments
//
// req - term of payment to create
func (c Client) CreateTermOfPayment(ctx context.Context, req *CreateTermOfPaymentReq) (*CreateTermOfPaymentResp, error) {
	resp := &CreateTermOfPaymentResp{}

	err := c._POST(ctx, termsOfPaymentsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetTermOfPayment does _GET https://api.fortnox.se/3/termsofpayments/{Code}
//
// code - identifies the terms of payment
func (c *Client) GetTermOfPayment(ctx context.Context, code string) (*GetTermOfPaymentResp, error) {
	resp := &GetTermOfPaymentResp{}

	uri := fmt.Sprintf("%s/%s", termsOfPaymentsURI, code)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateTermOfPayment does _PUT https://api.fortnox.se/3/termsofpayments/{Code}
//
// code - identifies the terms of payment
//
// req - term of payment to update
func (c *Client) UpdateTermOfPayment(
	ctx context.Context,
	code string,
	req *UpdateTermOfPaymentReq) (*UpdateTermOfPaymentResp, error) {

	resp := &UpdateTermOfPaymentResp{}

	uri := fmt.Sprintf("%s/%s", termsOfPaymentsURI, code)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// RemoveTermOfPayment does _DELETE https://api.fortnox.se/3/termsofpayments/{Code}
// code - identifies the terms of payment
func (c *Client) RemoveTermOfPayment(ctx context.Context, code string) error {
	uri := fmt.Sprintf("%s/%s", termsOfPaymentsURI, code)
	return c._DELETE(ctx, uri)
}

type GetAllTermsOfPaymentsResp struct {
	TermsOfPayments []struct {
		Url         string `json:"@url"`
		Code        string `json:"Code"`
		Description string `json:"Description"`
	} `json:"TermsOfPayments"`
}

type CreateTermOfPaymentReq struct {
	TermsOfPayment struct {
		Url         string `json:"@url"`
		Code        string `json:"Code"`
		Description string `json:"Description"`
	} `json:"TermsOfPayment"`
}

type CreateTermOfPaymentResp struct {
	TermsOfPayment struct {
		Url         string `json:"@url"`
		Code        string `json:"Code"`
		Description string `json:"Description"`
	} `json:"TermsOfPayment"`
}

type GetTermOfPaymentResp struct {
	TermsOfPayment struct {
		Url         string `json:"@url"`
		Code        string `json:"Code"`
		Description string `json:"Description"`
	} `json:"TermsOfPayment"`
}

type UpdateTermOfPaymentReq struct {
	TermsOfPayment struct {
		Url         string `json:"@url"`
		Code        string `json:"Code"`
		Description string `json:"Description"`
	} `json:"TermsOfPayment"`
}

type UpdateTermOfPaymentResp struct {
	TermsOfPayment struct {
		Url         string `json:"@url"`
		Code        string `json:"Code"`
		Description string `json:"Description"`
	} `json:"TermsOfPayment"`
}

package client

import (
	"context"
	"fmt"
)

const (
	supplierInvoiceExternalUrlConnectionsURI = "supplierinvoiceexternalurlconnections"
)

// GetSupplierInvoiceExternalUrlConnection does _GET
func (c *Client) GetSupplierInvoiceExternalUrlConnection(
	ctx context.Context,
	id int) (*GetSupplierInvoiceExternalUrlConnectionResp, error) {

	resp := &GetSupplierInvoiceExternalUrlConnectionResp{}

	uri := fmt.Sprintf("%s/%d", supplierInvoiceExternalUrlConnectionsURI, id)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateSupplierInvoiceExternalUrlConnection does _PUT
func (c *Client) UpdateSupplierInvoiceExternalUrlConnection(
	ctx context.Context,
	id int,
	req *UpdateSupplierInvoiceExternalUrlConnectionReq) (*UpdateSupplierInvoiceExternalUrlConnectionResp, error) {

	resp := &UpdateSupplierInvoiceExternalUrlConnectionResp{}

	uri := fmt.Sprintf("%s/%d", supplierInvoiceExternalUrlConnectionsURI, id)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteSupplierInvoiceExternalUrlConnection does _DELETE
func (c *Client) DeleteSupplierInvoiceExternalUrlConnection(ctx context.Context, id int) error {
	uri := fmt.Sprintf("%s/%d", supplierInvoiceExternalUrlConnectionsURI, id)
	return c._DELETE(ctx, uri)
}

// CreateSupplierInvoiceExternalUrlConnection does _POST
func (c *Client) CreateSupplierInvoiceExternalUrlConnection(
	ctx context.Context,
	req *CreateSupplierInvoiceExternalUrlConnectionReq) (*CreateSupplierInvoiceExternalUrlConnectionResp, error) {

	resp := &CreateSupplierInvoiceExternalUrlConnectionResp{}

	err := c._POST(ctx, supplierInvoiceExternalUrlConnectionsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetSupplierInvoiceExternalUrlConnectionResp struct {
	SupplierInvoiceExternalURLConnection struct {
		Url                   string `json:"Url"`
		Id                    int    `json:"Id"`
		SupplierInvoiceNumber int    `json:"SupplierInvoiceNumber"`
		ExternalUrlConnection string `json:"ExternalUrlConnection"`
	} `json:"SupplierInvoiceExternalURLConnection"`
}

type UpdateSupplierInvoiceExternalUrlConnectionReq struct {
	SupplierInvoiceNumber int    `json:"SupplierInvoiceNumber"`
	ExternalURLConnection string `json:"ExternalURLConnection"`
}

type UpdateSupplierInvoiceExternalUrlConnectionResp struct {
	SupplierInvoiceExternalURLConnection struct {
		Url                   string `json:"Url"`
		Id                    int    `json:"Id"`
		SupplierInvoiceNumber int    `json:"SupplierInvoiceNumber"`
		ExternalUrlConnection string `json:"ExternalUrlConnection"`
	} `json:"SupplierInvoiceExternalURLConnection"`
}

type CreateSupplierInvoiceExternalUrlConnectionReq struct {
	SupplierInvoiceNumber int    `json:"SupplierInvoiceNumber"`
	ExternalURLConnection string `json:"ExternalURLConnection"`
}

type CreateSupplierInvoiceExternalUrlConnectionResp struct {
	SupplierInvoiceExternalURLConnection struct {
		Url                   string `json:"Url"`
		Id                    int    `json:"Id"`
		SupplierInvoiceNumber int    `json:"SupplierInvoiceNumber"`
		ExternalUrlConnection string `json:"ExternalUrlConnection"`
	} `json:"SupplierInvoiceExternalURLConnection"`
}

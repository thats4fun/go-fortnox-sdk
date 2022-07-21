package client

import (
	"context"
	"fmt"
)

const (
	supplierInvoiceFileConnectionsURI = "supplierinvoicefileconnections"
)

// GetAllSupplierInvoiceFileConnections does _GET
func (c *Client) GetAllSupplierInvoiceFileConnections(
	ctx context.Context) (*GetAllSupplierInvoiceFileConnectionsResp, error) {

	resp := &GetAllSupplierInvoiceFileConnectionsResp{}

	err := c._GET(ctx, supplierInvoiceFileConnectionsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateSupplierInvoiceFileConnection does _POST
func (c *Client) CreateSupplierInvoiceFileConnection(
	ctx context.Context,
	req *CreateSupplierInvoiceFileConnectionReq) (*CreateSupplierInvoiceFileConnectionResp, error) {

	resp := &CreateSupplierInvoiceFileConnectionResp{}

	err := c._POST(ctx, supplierInvoiceFileConnectionsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetSupplierInvoiceFileConnection does _GET
func (c *Client) GetSupplierInvoiceFileConnection(
	ctx context.Context,
	fileID string) (*GetSupplierInvoiceFileConnectionResp, error) {

	resp := &GetSupplierInvoiceFileConnectionResp{}

	uri := fmt.Sprintf("%s/%s", supplierInvoiceFileConnectionsURI, fileID)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// RemoveSupplierInvoiceFileConnections does _DELETE
func (c *Client) RemoveSupplierInvoiceFileConnections(ctx context.Context, fieldID string) error {
	uri := fmt.Sprintf("%s/%s", supplierInvoiceFileConnectionsURI, fieldID)
	return c._DELETE(ctx, uri)
}

type GetAllSupplierInvoiceFileConnectionsResp struct {
	SupplierInvoiceFileConnections []struct {
		Url                   string `json:"@url"`
		FileId                string `json:"FileId"`
		Name                  string `json:"Name"`
		SupplierInvoiceNumber string `json:"SupplierInvoiceNumber"`
		SupplierName          string `json:"SupplierName"`
	} `json:"SupplierInvoiceFileConnections"`
}

type CreateSupplierInvoiceFileConnectionReq struct {
	SupplierInvoiceFileConnection struct {
		Url                   string `json:"@url"`
		FileId                string `json:"FileId"`
		Name                  string `json:"Name"`
		SupplierInvoiceNumber string `json:"SupplierInvoiceNumber"`
		SupplierName          string `json:"SupplierName"`
	} `json:"SupplierInvoiceFileConnection"`
}

type CreateSupplierInvoiceFileConnectionResp struct {
	SupplierInvoiceFileConnection struct {
		Url                   string `json:"@url"`
		FileId                string `json:"FileId"`
		Name                  string `json:"Name"`
		SupplierInvoiceNumber string `json:"SupplierInvoiceNumber"`
		SupplierName          string `json:"SupplierName"`
	} `json:"SupplierInvoiceFileConnection"`
}

type GetSupplierInvoiceFileConnectionResp struct {
	SupplierInvoiceFileConnection struct {
		Url                   string `json:"@url"`
		FileId                string `json:"FileId"`
		Name                  string `json:"Name"`
		SupplierInvoiceNumber string `json:"SupplierInvoiceNumber"`
		SupplierName          string `json:"SupplierName"`
	} `json:"SupplierInvoiceFileConnection"`
}

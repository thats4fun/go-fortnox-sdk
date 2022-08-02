package client

import (
	"context"
	"fmt"
)

const (
	supplierInvoiceFileConnectionsURI = "supplierinvoicefileconnections"
)

// GetAllSupplierInvoiceFileConnections does _GET https://api.fortnox.se/3/supplierinvoicefileconnections/
func (c *Client) GetAllSupplierInvoiceFileConnections(ctx context.Context) ([]SupplierInvoiceFileConnection, error) {
	resp := &GetAllSupplierInvoiceFileConnectionsResp{}

	err := c._GET(ctx, supplierInvoiceFileConnectionsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.SupplierInvoiceFileConnections, nil
}

// CreateSupplierInvoiceFileConnection does _POST https://api.fortnox.se/3/supplierinvoicefileconnections/
//
// sifc - supplier invoice file connection to create
func (c *Client) CreateSupplierInvoiceFileConnection(
	ctx context.Context,
	sifc *SupplierInvoiceFileConnection) (*SupplierInvoiceFileConnection, error) {

	req := &CreateSupplierInvoiceFileConnectionReq{SupplierInvoiceFileConnection: *sifc}
	resp := &CreateSupplierInvoiceFileConnectionResp{}

	err := c._POST(ctx, supplierInvoiceFileConnectionsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.SupplierInvoiceFileConnection, nil
}

// GetSupplierInvoiceFileConnection does _GET https://api.fortnox.se/3/supplierinvoicefileconnections/{FileId}
//
// fileID - identifies the file connection
func (c *Client) GetSupplierInvoiceFileConnection(
	ctx context.Context,
	fileID string) (*SupplierInvoiceFileConnection, error) {

	resp := &GetSupplierInvoiceFileConnectionResp{}

	uri := fmt.Sprintf("%s/%s", supplierInvoiceFileConnectionsURI, fileID)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.SupplierInvoiceFileConnection, nil
}

// RemoveSupplierInvoiceFileConnections does _DELETE https://api.fortnox.se/3/supplierinvoicefileconnections/{FileId}
//
// fileID - identifies the file connection
func (c *Client) RemoveSupplierInvoiceFileConnections(ctx context.Context, fieldID string) error {
	uri := fmt.Sprintf("%s/%s", supplierInvoiceFileConnectionsURI, fieldID)
	return c._DELETE(ctx, uri)
}

type SupplierInvoiceFileConnection struct {
	Url                   string `json:"@url,omitempty"`
	FileId                string `json:"FileId,omitempty"`
	Name                  string `json:"Name,omitempty"`
	SupplierInvoiceNumber string `json:"SupplierInvoiceNumber,omitempty"`
	SupplierName          string `json:"SupplierName,omitempty"`
}

type GetAllSupplierInvoiceFileConnectionsResp struct {
	SupplierInvoiceFileConnections []SupplierInvoiceFileConnection `json:"SupplierInvoiceFileConnections"`
}

type CreateSupplierInvoiceFileConnectionReq struct {
	SupplierInvoiceFileConnection SupplierInvoiceFileConnection `json:"SupplierInvoiceFileConnection"`
}

type CreateSupplierInvoiceFileConnectionResp struct {
	SupplierInvoiceFileConnection SupplierInvoiceFileConnection `json:"SupplierInvoiceFileConnection"`
}

type GetSupplierInvoiceFileConnectionResp struct {
	SupplierInvoiceFileConnection SupplierInvoiceFileConnection `json:"SupplierInvoiceFileConnection"`
}

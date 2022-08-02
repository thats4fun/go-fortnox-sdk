package client

import (
	"context"
	"fmt"
)

const (
	voucherFileConnectionsURI = "voucherfileconnections"
)

// GetAllVoucherFileConnections does _GET https://api.fortnox.se/3/voucherfileconnections/
func (c *Client) GetAllVoucherFileConnections(ctx context.Context) ([]VoucherFileConnection, error) {
	resp := &GetAllVoucherFileConnectionsResp{}

	err := c._GET(ctx, voucherFileConnectionsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.VoucherFileConnections, nil
}

// CreateVoucherFileConnection does _POST https://api.fortnox.se/3/voucherfileconnections/
//
// req - voucher file connection to create
func (c *Client) CreateVoucherFileConnection(
	ctx context.Context,
	vfc *VoucherFileConnection) (*VoucherFileConnection, error) {

	req := &CreateVoucherFileConnectionReq{VoucherFileConnection: *vfc}
	resp := &CreateVoucherFileConnectionResp{}

	err := c._POST(ctx, voucherFileConnectionsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.VoucherFileConnection, nil
}

// GetVoucherFileConnectionByFileID does _GET https://api.fortnox.se/3/voucherfileconnections/{FileId}
//
// fileID - identifies the voucher file connection
func (c *Client) GetVoucherFileConnectionByFileID(
	ctx context.Context,
	fileID string) (*VoucherFileConnection, error) {

	resp := &GetVoucherFileConnectionByFileIDResp{}

	uri := fmt.Sprintf("%s/%s", voucherFileConnectionsURI, fileID)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.VoucherFileConnection, nil
}

// RemoveVoucherFileConnections does _DELETE https://api.fortnox.se/3/voucherfileconnections/{FileId}
//
// fileID - identifies the voucher file connection
func (c *Client) RemoveVoucherFileConnections(ctx context.Context, fieldID string) error {
	uri := fmt.Sprintf("%s/%s", voucherFileConnectionsURI, fieldID)
	return c._DELETE(ctx, uri)
}

type VoucherFileConnection struct {
	Url                string `json:"@url"`
	FileId             string `json:"FileId"`
	VoucherDescription string `json:"VoucherDescription"`
	VoucherNumber      string `json:"VoucherNumber"`
	VoucherSeries      string `json:"VoucherSeries"`
	VoucherYear        int    `json:"VoucherYear"`
}

type GetAllVoucherFileConnectionsResp struct {
	VoucherFileConnections []VoucherFileConnection `json:"VoucherFileConnections"`
}

type CreateVoucherFileConnectionReq struct {
	VoucherFileConnection VoucherFileConnection `json:"VoucherFileConnection"`
}

type CreateVoucherFileConnectionResp struct {
	VoucherFileConnection VoucherFileConnection `json:"VoucherFileConnection"`
}

type GetVoucherFileConnectionByFileIDResp struct {
	VoucherFileConnection VoucherFileConnection `json:"VoucherFileConnection"`
}

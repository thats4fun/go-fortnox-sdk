package client

import (
	"context"
	"fmt"
)

const (
	assetFileConnectionsURI = "assetfileconnections"
)

// GetAllAssetFileConnections does _GET https://api.fortnox.se/3/assetfileconnections
func (c *Client) GetAllAssetFileConnections(ctx context.Context) (*GetAllAssetFileConnectionsResp, error) {
	resp := &GetAllAssetFileConnectionsResp{}

	err := c._GET(ctx, assetFileConnectionsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateAssetFileConnection does _POST https://api.fortnox.se/3/assetfileconnections
//
// req - asset file connection to create
func (c *Client) CreateAssetFileConnection(
	ctx context.Context,
	req *CreateAssetFileConnectionReq) (*AssetFileConnection, error) {

	resp := &AssetFileConnection{}

	err := c._POST(ctx, assetFileConnectionsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteAssetFileConnection does _DELETE https://api.fortnox.se/3/assetfileconnections/{fileId}
//
// fileID - fileId
func (c *Client) DeleteAssetFileConnection(ctx context.Context, fileID string) error {
	uri := fmt.Sprintf("%s/%s", articleFileConnectionsURI, fileID)
	return c._DELETE(ctx, uri)
}

type AssetFileConnection struct {
	Url     string `json:"@url"`
	FileId  string `json:"FileId"`
	Name    string `json:"Name"`
	AssetId string `json:"AssetId"`
}

type AssetFileConnectionMetaInformation struct {
	TotalResources int `json:"@TotalResources"`
	TotalPages     int `json:"@TotalPages"`
	CurrentPage    int `json:"@CurrentPage"`
}

type GetAllAssetFileConnectionsResp struct {
	AssetFileConnections []AssetFileConnection              `json:"AssetFileConnections"`
	MetaInformation      AssetFileConnectionMetaInformation `json:"MetaInformation"`
}

type CreateAssetFileConnectionReq struct {
	FileId  string `json:"FileId"`
	AssetId string `json:"AssetId"`
}

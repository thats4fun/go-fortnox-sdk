package client

import (
	"context"
	"fmt"
)

const (
	assetFileConnectionsURI = "assetfileconnections"
)

// GetAllAssetFileConnections does _GET https://api.fortnox.se/3/assetfileconnections
func (c *Client) GetAllAssetFileConnections(ctx context.Context) (*GetAllAssetFileConnections, error) {
	resp := &GetAllAssetFileConnections{}

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
	req *CreateAssetFileConnectionReq) (*CreateAssetFileConnectionResp, error) {

	resp := &CreateAssetFileConnectionResp{}

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

type GetAllAssetFileConnections struct {
	AssetFileConnections []struct {
		Url     string `json:"@url"`
		FileId  string `json:"FileId"`
		Name    string `json:"Name"`
		AssetId string `json:"AssetId"`
	} `json:"AssetFileConnections"`
	MetaInformation struct {
		TotalResources int `json:"@TotalResources"`
		TotalPages     int `json:"@TotalPages"`
		CurrentPage    int `json:"@CurrentPage"`
	} `json:"MetaInformation"`
}

type CreateAssetFileConnectionReq struct {
	FileID  string
	AssetID string
}

type CreateAssetFileConnectionResp struct {
	Url     string `json:"@url"`
	FileId  string `json:"FileId"`
	Name    string `json:"Name"`
	AssetId string `json:"AssetId"`
}

package client

import (
	"context"
	"fmt"
)

const (
	articleFileConnectionsURI = "articlefileconnections"
)

// GetAllArticleFileConnections does _GET https://api.fortnox.se/3/articlefileconnections/
func (c *Client) GetAllArticleFileConnections(ctx context.Context) ([]ArticleFileConnection, error) {
	resp := &GetAllArticleFileConnectionsResp{}

	err := c._GET(ctx, articleFileConnectionsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.ArticleFileConnections, nil
}

// CreateArticleFileConnection does _POST https://api.fortnox.se/3/articlefileconnections/
//
// req - article file connection to create
func (c *Client) CreateArticleFileConnection(ctx context.Context, afc *ArticleFileConnection) (*ArticleFileConnection, error) {
	req := CreateArticleFileConnectionReq{ArticleFileConnection: *afc}
	resp := &CreateArticleFileConnectionResp{}

	err := c._POST(ctx, articleFileConnectionsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.ArticleFileConnection, nil
}

// GetArticleFileConnectionByID does _GET https://api.fortnox.se/3/articlefileconnections/{FileId}
//
// fileID - identifies the article file connection
func (c *Client) GetArticleFileConnectionByID(ctx context.Context, fileID string) (*ArticleFileConnection, error) {
	resp := &GetArticleFileConnectionByIDResp{}

	uri := fmt.Sprintf("%s/%s", articleFileConnectionsURI, fileID)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.ArticleFileConnection, nil
}

// DeleteArticleFileConnection does _DELETE https://api.fortnox.se/3/articlefileconnections/{FileId}
//
// fileID - identifies the article file connection
func (c *Client) DeleteArticleFileConnection(ctx context.Context, fileID string) error {
	uri := fmt.Sprintf("%s/%s", articleFileConnectionsURI, fileID)
	return c._DELETE(ctx, uri)
}

type ArticleFileConnection struct {
	Url           string `json:"@url,omitempty"`
	FileId        string `json:"FileId,omitempty"`
	ArticleNumber string `json:"ArticleNumber,omitempty"`
}

type GetAllArticleFileConnectionsResp struct {
	ArticleFileConnections []ArticleFileConnection `json:"ArticleFileConnections"`
}

type CreateArticleFileConnectionReq struct {
	ArticleFileConnection ArticleFileConnection `json:"ArticleFileConnection"`
}

type CreateArticleFileConnectionResp struct {
	ArticleFileConnection ArticleFileConnection `json:"ArticleFileConnection"`
}

type GetArticleFileConnectionByIDResp struct {
	ArticleFileConnection ArticleFileConnection `json:"ArticleFileConnection"`
}

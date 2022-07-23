package client

import (
	"context"
	"fmt"
)

const (
	articleFileConnectionsURI = "articlefileconnections"
)

// GetAllArticleFileConnections does _GET https://api.fortnox.se/3/articlefileconnections/
func (c *Client) GetAllArticleFileConnections(ctx context.Context) (*GetAllArticleFileConnectionsResp, error) {
	resp := &GetAllArticleFileConnectionsResp{}

	err := c._GET(ctx, articleFileConnectionsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateArticleFileConnection does _POST https://api.fortnox.se/3/articlefileconnections/
//
// req - article file connection to create
func (c *Client) CreateArticleFileConnection(ctx context.Context, req *CreateArticleFileConnectionReq) (*CreateArticleFileConnectionResp, error) {
	resp := &CreateArticleFileConnectionResp{}

	err := c._POST(ctx, articleFileConnectionsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetArticleFileConnectionByID does _GET https://api.fortnox.se/3/articlefileconnections/{FileId}
//
// fileID - identifies the article file connection
func (c *Client) GetArticleFileConnectionByID(ctx context.Context, fileID string) (*GetArticleFileConnectionByIDResp, error) {
	resp := &GetArticleFileConnectionByIDResp{}

	uri := fmt.Sprintf("%s/%s", articleFileConnectionsURI, fileID)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteArticleFileConnection does _DELETE https://api.fortnox.se/3/articlefileconnections/{FileId}
//
// fileID - identifies the article file connection
func (c *Client) DeleteArticleFileConnection(ctx context.Context, fileID string) error {
	uri := fmt.Sprintf("%s/%s", articleFileConnectionsURI, fileID)
	return c._DELETE(ctx, uri)
}

type GetAllArticleFileConnectionsResp struct {
	ArticleFileConnections []struct {
		Url           string `json:"@url"`
		FileId        string `json:"FileId"`
		ArticleNumber string `json:"ArticleNumber"`
	} `json:"ArticleFileConnections"`
}

type CreateArticleFileConnectionReq struct {
	ArticleFileConnection struct {
		Url           string `json:"@url"`
		FileId        string `json:"FileId"`
		ArticleNumber string `json:"ArticleNumber"`
	} `json:"ArticleFileConnection"`
}

type CreateArticleFileConnectionResp struct {
	ArticleFileConnection struct {
		Url           string `json:"@url"`
		FileId        string `json:"FileId"`
		ArticleNumber string `json:"ArticleNumber"`
	} `json:"ArticleFileConnection"`
}

type GetArticleFileConnectionByIDResp struct {
	ArticleFileConnection struct {
		Url           string `json:"@url"`
		FileId        string `json:"FileId"`
		ArticleNumber string `json:"ArticleNumber"`
	} `json:"ArticleFileConnection"`
}

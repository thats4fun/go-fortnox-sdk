package client

import (
	"context"
	"fmt"
)

const (
	articleFileConnectionsURI = "articlefileconnections"
)

// GetAllArticleFileConnections does _GET https://api.fortnox.se/3/articlefileconnections/
func (c *Client) GetAllArticleFileConnections(ctx context.Context) ([]*ArticleFileConnection, error) {
	resp := &GetAllArticleFileConnectionsResp{}

	err := c._GET(ctx, articleFileConnectionsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.ArticleFileConnections, nil
}

// CreateArticleFileConnection does _POST https://api.fortnox.se/3/articlefileconnections/
// req - object to create
func (c *Client) CreateArticleFileConnection(ctx context.Context, req *CreateArticleFileConnectionReq) (*CreateArticleFileConnectionResp, error) {
	resp := &CreateArticleFileConnectionResp{}

	err := c._POST(ctx, articleFileConnectionsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetArticleFileConnectionByID does _GET https://api.fortnox.se/3/articlefileconnections/{FileId}
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

// Delete does _DELETE https://api.fortnox.se/3/articlefileconnections/{FileId}
// fileID - identifies the article file connection
func (c *Client) Delete(ctx context.Context, fileID string) error {
	uri := fmt.Sprintf("%s/%s", articleFileConnectionsURI, fileID)
	return c._DELETE(ctx, uri)
}

type GetAllArticleFileConnectionsResp struct {
	ArticleFileConnections []*ArticleFileConnection `json:"ArticleFileConnections"`
}

type ArticleFileConnection struct {
	Url           string `json:"@url"`
	FileId        string `json:"FileId"`
	ArticleNumber string `json:"ArticleNumber"`
}

type CreateArticleFileConnectionReq struct {
	ArticleFileConnection ArticleFileConnection `json:"ArticleFileConnection"`
}

type CreateArticleFileConnectionResp CreateArticleFileConnectionReq

type GetArticleFileConnectionByIDResp CreateArticleFileConnectionResp

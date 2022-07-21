package client

import (
	"context"
	"fmt"
	"time"
)

const (
	pricesURI = "prices"
)

// GetPriceForArticle does _GET
//
// priceList -
//
// articleNumber -
//
// fromQuantity -
func (c *Client) GetPriceForArticle(
	ctx context.Context,
	priceList, articleNumber string,
	fromQuantity int) (*GetPriceForArticleResp, error) {

	resp := &GetPriceForArticleResp{}

	uri := fmt.Sprintf("%s/%s/%s/%d", pricesURI, priceList, articleNumber, fromQuantity)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdatePrice doe _PUT
//
// priceList -
//
// articleNumber -
//
// fromQuantity -
//
// req -
func (c *Client) UpdatePrice(
	ctx context.Context,
	priceList, articleNumber string,
	fromQuantity int,
	req *UpdatePriceReq) (*UpdatePriceResp, error) {

	resp := &UpdatePriceResp{}

	uri := fmt.Sprintf("%s/%s/%s/%d", pricesURI, priceList, articleNumber, fromQuantity)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeletePrice does _DELETE
//
// priceList -
//
// articleNumber -
//
// fromQuantity -
func (c *Client) DeletePrice(ctx context.Context, priceList, articleNumber string, fromQuantity int) error {
	uri := fmt.Sprintf("%s/%s/%s/%d", pricesURI, priceList, articleNumber, fromQuantity)
	return c._DELETE(ctx, uri)
}

// GetAllArticlesWithPricesInPriceList does _GET
//
// priceList -
//
// articleNumber -
func (c *Client) GetAllArticlesWithPricesInPriceList(
	ctx context.Context,
	priceList, articleNumber string) (*GetAllArticlesWithPricesInPriceListResp, error) {

	resp := &GetAllArticlesWithPricesInPriceListResp{}

	uri := fmt.Sprintf("%s/%s/%s", pricesURI, priceList, articleNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetFirstPriceForArticle does _GET
//
// priceList -
//
// articleNumber -
func (c *Client) GetFirstPriceForArticle(
	ctx context.Context,
	priceList, articleNumber string) (*GetFirstPriceForArticleResp, error) {

	resp := &GetFirstPriceForArticleResp{}

	uri := fmt.Sprintf("%s/%s/%s", pricesURI, priceList, articleNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateFirstPriceForArticle does _PUT
//
// priceList -
//
// articleNumber -
//
// req -
func (c *Client) UpdateFirstPriceForArticle(
	ctx context.Context,
	priceList, articleNumber string,
	req *UpdateFirstPriceForArticleReq) (*UpdateFirstPriceForArticleResp, error) {

	resp := &UpdateFirstPriceForArticleResp{}

	uri := fmt.Sprintf("%s/%s/%s", pricesURI, priceList, articleNumber)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreatePrice does _POST
//
// req -
func (c *Client) CreatePrice(ctx context.Context, req *CreatePriceReq) (*CreatePriceResp, error) {
	resp := &CreatePriceResp{}

	err := c._POST(ctx, pricesURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetPriceForArticleResp struct {
	Price struct {
		Url           string    `json:"@url"`
		ArticleNumber string    `json:"ArticleNumber"`
		Date          time.Time `json:"Date"`
		FromQuantity  int       `json:"FromQuantity"`
		Percent       int       `json:"Percent"`
		Price         int       `json:"Price"`
		PriceList     string    `json:"PriceList"`
	} `json:"Price"`
}

type UpdatePriceReq struct {
	Price struct {
		Url           string    `json:"@url"`
		ArticleNumber string    `json:"ArticleNumber"`
		Date          time.Time `json:"Date"`
		FromQuantity  int       `json:"FromQuantity"`
		Percent       int       `json:"Percent"`
		Price         int       `json:"Price"`
		PriceList     string    `json:"PriceList"`
	} `json:"Price"`
}

type UpdatePriceResp struct {
	Price struct {
		Url           string    `json:"@url"`
		ArticleNumber string    `json:"ArticleNumber"`
		Date          time.Time `json:"Date"`
		FromQuantity  int       `json:"FromQuantity"`
		Percent       int       `json:"Percent"`
		Price         int       `json:"Price"`
		PriceList     string    `json:"PriceList"`
	} `json:"Price"`
}

type GetAllArticlesWithPricesInPriceListResp struct {
	Prices []struct {
		Url           string `json:"@url"`
		ArticleNumber string `json:"ArticleNumber"`
		FromQuantity  int    `json:"FromQuantity"`
		PriceList     string `json:"PriceList"`
		Price         int    `json:"Price"`
	} `json:"Prices"`
}

type GetFirstPriceForArticleResp struct {
	Price struct {
		Url           string    `json:"@url"`
		ArticleNumber string    `json:"ArticleNumber"`
		Date          time.Time `json:"Date"`
		FromQuantity  int       `json:"FromQuantity"`
		Percent       int       `json:"Percent"`
		Price         int       `json:"Price"`
		PriceList     string    `json:"PriceList"`
	} `json:"Price"`
}

type UpdateFirstPriceForArticleReq struct {
	Price struct {
		Url           string    `json:"@url"`
		ArticleNumber string    `json:"ArticleNumber"`
		Date          time.Time `json:"Date"`
		FromQuantity  int       `json:"FromQuantity"`
		Percent       int       `json:"Percent"`
		Price         int       `json:"Price"`
		PriceList     string    `json:"PriceList"`
	} `json:"Price"`
}

type UpdateFirstPriceForArticleResp struct {
	Price struct {
		Url           string    `json:"@url"`
		ArticleNumber string    `json:"ArticleNumber"`
		Date          time.Time `json:"Date"`
		FromQuantity  int       `json:"FromQuantity"`
		Percent       int       `json:"Percent"`
		Price         int       `json:"Price"`
		PriceList     string    `json:"PriceList"`
	} `json:"Price"`
}

type CreatePriceReq struct {
	Price struct {
		Url           string    `json:"@url"`
		ArticleNumber string    `json:"ArticleNumber"`
		Date          time.Time `json:"Date"`
		FromQuantity  int       `json:"FromQuantity"`
		Percent       int       `json:"Percent"`
		Price         int       `json:"Price"`
		PriceList     string    `json:"PriceList"`
	} `json:"Price"`
}

type CreatePriceResp struct {
	Price struct {
		Url           string    `json:"@url"`
		ArticleNumber string    `json:"ArticleNumber"`
		Date          time.Time `json:"Date"`
		FromQuantity  int       `json:"FromQuantity"`
		Percent       int       `json:"Percent"`
		Price         int       `json:"Price"`
		PriceList     string    `json:"PriceList"`
	} `json:"Price"`
}

package client

import (
	"context"
	"fmt"
	"time"
)

const (
	pricesURI = "prices"
)

// GetPriceForArticle does _GET https://api.fortnox.se/3/prices/{PriceList}/{ArticleNumber}/{FromQuantity}
//
// priceList - identifies the price list
//
// articleNumber - identifies the article
//
// fromQuantity - identifies from quantity
func (c *Client) GetPriceForArticle(
	ctx context.Context,
	priceList, articleNumber string,
	fromQuantity int) (*Price, error) {

	resp := &GetPriceForArticleResp{}

	uri := fmt.Sprintf("%s/%s/%s/%d", pricesURI, priceList, articleNumber, fromQuantity)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Price, nil
}

// UpdatePrice does _PUT https://api.fortnox.se/3/prices/{PriceList}/{ArticleNumber}/{FromQuantity}
//
// priceList - identifies the price list
//
// articleNumber - identifies the article
//
// fromQuantity - identifies from quantity
//
// req - price to update
func (c *Client) UpdatePrice(
	ctx context.Context,
	priceList, articleNumber string,
	fromQuantity int,
	p *Price) (*Price, error) {

	req := &UpdatePriceReq{Price: *p}
	resp := &UpdatePriceResp{}

	uri := fmt.Sprintf("%s/%s/%s/%d", pricesURI, priceList, articleNumber, fromQuantity)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Price, nil
}

// DeletePrice does _DELETE https://api.fortnox.se/3/prices/{PriceList}/{ArticleNumber}/{FromQuantity}
//
// priceList - identifies the price list
//
// articleNumber - identifies the article
//
// fromQuantity - identifies from quantity
func (c *Client) DeletePrice(ctx context.Context, priceList, articleNumber string, fromQuantity int) error {
	uri := fmt.Sprintf("%s/%s/%s/%d", pricesURI, priceList, articleNumber, fromQuantity)
	return c._DELETE(ctx, uri)
}

// GetAllArticlesWithPricesInPriceList does _GET https://api.fortnox.se/3/prices/sublist/{PriceList}/{ArticleNumber}
//
// priceList - identifies the price list of the prices
//
// articleNumber - identifies the article number of the prices
func (c *Client) GetAllArticlesWithPricesInPriceList(
	ctx context.Context,
	priceList, articleNumber string) ([]Price, error) {

	resp := &GetAllArticlesWithPricesInPriceListResp{}

	uri := fmt.Sprintf("%s/sublist/%s/%s", pricesURI, priceList, articleNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.Prices, nil
}

// GetFirstPriceForArticle does _GET https://api.fortnox.se/3/prices/{PriceList}/{ArticleNumber}
//
// priceList - identifies the price list of the prices
//
// articleNumber - identifies the article number of the prices
func (c *Client) GetFirstPriceForArticle(
	ctx context.Context,
	priceList, articleNumber string) (*Price, error) {

	resp := &GetFirstPriceForArticleResp{}

	uri := fmt.Sprintf("%s/%s/%s", pricesURI, priceList, articleNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Price, nil
}

// UpdateFirstPriceForArticle does _PUT https://api.fortnox.se/3/prices/{PriceList}/{ArticleNumber}
//
// priceList - identifies the price list of the prices
//
// articleNumber - identifies the article number of the prices
//
// req - price to update
func (c *Client) UpdateFirstPriceForArticle(
	ctx context.Context,
	priceList, articleNumber string,
	p *Price) (*Price, error) {

	req := &UpdateFirstPriceForArticleReq{Price: *p}
	resp := &UpdateFirstPriceForArticleResp{}

	uri := fmt.Sprintf("%s/%s/%s", pricesURI, priceList, articleNumber)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Price, nil
}

// CreatePrice does _POST https://api.fortnox.se/3/prices/
//
// req - price to create
func (c *Client) CreatePrice(ctx context.Context, p *Price) (*Price, error) {
	req := &CreatePriceReq{Price: *p}
	resp := &CreatePriceResp{}

	err := c._POST(ctx, pricesURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Price, nil
}

type Price struct {
	Url           string    `json:"@url,omitempty"`
	ArticleNumber string    `json:"ArticleNumber,omitempty"`
	Date          time.Time `json:"Date,omitempty,omitempty"`
	FromQuantity  int       `json:"FromQuantity,omitempty"`
	Percent       int       `json:"Percent,omitempty,omitempty"`
	Price         int       `json:"Price,omitempty"`
	PriceList     string    `json:"PriceList,omitempty"`
}

type GetPriceForArticleResp struct {
	Price Price `json:"Price"`
}

type UpdatePriceReq struct {
	Price Price `json:"Price"`
}

type UpdatePriceResp struct {
	Price Price `json:"Price"`
}

type GetAllArticlesWithPricesInPriceListResp struct {
	Prices []Price `json:"Prices"`
}

type GetFirstPriceForArticleResp struct {
	Price Price `json:"Price"`
}

type UpdateFirstPriceForArticleReq struct {
	Price Price `json:"Price"`
}

type UpdateFirstPriceForArticleResp struct {
	Price Price `json:"Price"`
}

type CreatePriceReq struct {
	Price Price `json:"Price"`
}

type CreatePriceResp struct {
	Price Price `json:"Price"`
}

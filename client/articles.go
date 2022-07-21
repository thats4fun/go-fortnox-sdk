package client

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

const (
	articlesURI = "articles"
)

const (
	filterParamName = "filter"
)

// GetArticle does _GET https://api.fortnox.se/3/articles/{ArticleNumber}
//
// articleNumber - identifies the article
func (c *Client) GetArticle(ctx context.Context, articleNumber int) (*Article, error) {
	resp := &GetArticleResp{}

	uri := fmt.Sprintf("%s/%d", articlesURI, articleNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Article, nil
}

// UpdateArticle does _PUT https://api.fortnox.se/3/articles/{ArticleNumber}
//
// articleNumber - identifies the article
//
// updateArticle - article to update
func (c *Client) UpdateArticle(ctx context.Context, articleNumber int, updateArticle *Article) (*UpdateArticleResp, error) {
	resp := &UpdateArticleResp{}

	uri := fmt.Sprintf("%s/%d", articlesURI, articleNumber)

	err := c._PUT(ctx, uri, nil, updateArticle, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteArticle does _DELETE https://api.fortnox.se/3/articles/{ArticleNumber}
//
// articleNumber - identifies the article
func (c *Client) DeleteArticle(ctx context.Context, articleNumber int) error {
	uri := fmt.Sprintf("%s/%d", articlesURI, articleNumber)
	return c._DELETE(ctx, uri)
}

// GetArticles does _GET https://api.fortnox.se/3/articles
//
// filter - Enum: {"active", "inactive"}, possibility to filter supplier invoices
func (c *Client) GetArticles(ctx context.Context, filter *ArticleFilter) (*GetArticlesResp, error) {
	resp := &GetArticlesResp{}

	params := filter.urlValues()

	err := c._GET(ctx, articlesURI, params, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateArticle does _POST https://api.fortnox.se/3/articles
//
// article - object to create
func (c *Client) CreateArticle(ctx context.Context, article *Article) (*CreateArticleResp, error) {
	resp := &CreateArticleResp{}

	err := c._POST(ctx, articlesURI, nil, article, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type ArticleFilter string

func (f *ArticleFilter) urlValues() url.Values {
	if f == nil {
		return nil
	}

	params := url.Values{}
	fStr := string(*f)

	if strings.TrimSpace(fStr) != "" {
		params[filterParamName] = []string{fStr}
	}

	return params
}

const (
	ActiveArticle   ArticleFilter = "active"
	InactiveArticle ArticleFilter = "inactive"
)

type Article struct {
	Url                       string `json:"@url"`
	ArticleNumber             string `json:"ArticleNumber"`
	Bulky                     bool   `json:"Bulky"`
	ConstructionAccount       int    `json:"ConstructionAccount"`
	Depth                     int    `json:"Depth"`
	Description               string `json:"Description"`
	DisposableQuantity        int    `json:"DisposableQuantity"`
	EAN                       string `json:"EAN"`
	EUAccount                 int    `json:"EUAccount"`
	EUVATAccount              int    `json:"EUVATAccount"`
	ExportAccount             int    `json:"ExportAccount"`
	Height                    int    `json:"Height"`
	Housework                 bool   `json:"Housework"`
	HouseworkType             string `json:"HouseworkType"`
	Active                    bool   `json:"Active"`
	Manufacturer              string `json:"Manufacturer"`
	ManufacturerArticleNumber string `json:"ManufacturerArticleNumber"`
	Note                      string `json:"Note"`
	PurchaseAccount           int    `json:"PurchaseAccount"`
	PurchasePrice             int    `json:"PurchasePrice"`
	QuantityInStock           int    `json:"QuantityInStock"`
	ReservedQuantity          int    `json:"ReservedQuantity"`
	SalesAccount              int    `json:"SalesAccount"`
	StockGoods                bool   `json:"StockGoods"`
	StockPlace                string `json:"StockPlace"`
	StockValue                int    `json:"StockValue"`
	StockWarning              int    `json:"StockWarning"`
	SupplierName              string `json:"SupplierName"`
	SupplierNumber            string `json:"SupplierNumber"`
	Type                      string `json:"Type"`
	Unit                      string `json:"Unit"`
	VAT                       int    `json:"VAT"`
	WebshopArticle            bool   `json:"WebshopArticle"`
	Weight                    int    `json:"Weight"`
	Width                     int    `json:"Width"`
	Expired                   bool   `json:"Expired"`
	SalesPrice                int    `json:"SalesPrice"`
	CostCalculationMethod     string `json:"CostCalculationMethod"`
	StockAccount              int    `json:"StockAccount"`
	StockChangeAccount        int    `json:"StockChangeAccount"`
	DirectCost                int    `json:"DirectCost"`
	FreightCost               int    `json:"FreightCost"`
	OtherCost                 int    `json:"OtherCost"`
	DefaultStockPoint         string `json:"DefaultStockPoint"`
	DefaultStockLocation      string `json:"DefaultStockLocation"`
}

type GetArticleResp struct {
	Article Article `json:"Article"`
}

type UpdateArticleResp struct {
	Article struct {
		Url                       string `json:"@url"`
		ArticleNumber             string `json:"ArticleNumber"`
		Bulky                     bool   `json:"Bulky"`
		ConstructionAccount       int    `json:"ConstructionAccount"`
		Depth                     int    `json:"Depth"`
		Description               string `json:"Description"`
		DisposableQuantity        int    `json:"DisposableQuantity"`
		EAN                       string `json:"EAN"`
		EUAccount                 int    `json:"EUAccount"`
		EUVATAccount              int    `json:"EUVATAccount"`
		ExportAccount             int    `json:"ExportAccount"`
		Height                    int    `json:"Height"`
		Housework                 bool   `json:"Housework"`
		HouseworkType             string `json:"HouseworkType"`
		Active                    bool   `json:"Active"`
		Manufacturer              string `json:"Manufacturer"`
		ManufacturerArticleNumber string `json:"ManufacturerArticleNumber"`
		Note                      string `json:"Note"`
		PurchaseAccount           int    `json:"PurchaseAccount"`
		PurchasePrice             int    `json:"PurchasePrice"`
		QuantityInStock           int    `json:"QuantityInStock"`
		ReservedQuantity          int    `json:"ReservedQuantity"`
		SalesAccount              int    `json:"SalesAccount"`
		StockGoods                bool   `json:"StockGoods"`
		StockPlace                string `json:"StockPlace"`
		StockValue                int    `json:"StockValue"`
		StockWarning              int    `json:"StockWarning"`
		SupplierName              string `json:"SupplierName"`
		SupplierNumber            string `json:"SupplierNumber"`
		Type                      string `json:"Type"`
		Unit                      string `json:"Unit"`
		VAT                       int    `json:"VAT"`
		WebshopArticle            bool   `json:"WebshopArticle"`
		Weight                    int    `json:"Weight"`
		Width                     int    `json:"Width"`
		Expired                   bool   `json:"Expired"`
		SalesPrice                int    `json:"SalesPrice"`
		CostCalculationMethod     string `json:"CostCalculationMethod"`
		StockAccount              int    `json:"StockAccount"`
		StockChangeAccount        int    `json:"StockChangeAccount"`
		DirectCost                int    `json:"DirectCost"`
		FreightCost               int    `json:"FreightCost"`
		OtherCost                 int    `json:"OtherCost"`
		DefaultStockPoint         string `json:"DefaultStockPoint"`
		DefaultStockLocation      string `json:"DefaultStockLocation"`
	} `json:"Article"`
}

type GetArticlesResp struct {
	Articles []struct {
		Url                string `json:"@url"`
		ArticleNumber      string `json:"ArticleNumber"`
		Description        string `json:"Description"`
		DisposableQuantity string `json:"DisposableQuantity"`
		EAN                string `json:"EAN"`
		Housework          bool   `json:"Housework"`
		PurchasePrice      string `json:"PurchasePrice"`
		SalesPrice         string `json:"SalesPrice"`
		QuantityInStock    int    `json:"QuantityInStock"`
		ReservedQuantity   string `json:"ReservedQuantity"`
		StockPlace         string `json:"StockPlace"`
		StockValue         string `json:"StockValue"`
		Unit               string `json:"Unit"`
		VAT                string `json:"VAT"`
		WebshopArticle     bool   `json:"WebshopArticle"`
	} `json:"Articles"`
}

type CreateArticleResp struct {
	Article struct {
		Url                       string `json:"@url"`
		ArticleNumber             string `json:"ArticleNumber"`
		Bulky                     bool   `json:"Bulky"`
		ConstructionAccount       int    `json:"ConstructionAccount"`
		Depth                     int    `json:"Depth"`
		Description               string `json:"Description"`
		DisposableQuantity        int    `json:"DisposableQuantity"`
		EAN                       string `json:"EAN"`
		EUAccount                 int    `json:"EUAccount"`
		EUVATAccount              int    `json:"EUVATAccount"`
		ExportAccount             int    `json:"ExportAccount"`
		Height                    int    `json:"Height"`
		Housework                 bool   `json:"Housework"`
		HouseworkType             string `json:"HouseworkType"`
		Active                    bool   `json:"Active"`
		Manufacturer              string `json:"Manufacturer"`
		ManufacturerArticleNumber string `json:"ManufacturerArticleNumber"`
		Note                      string `json:"Note"`
		PurchaseAccount           int    `json:"PurchaseAccount"`
		PurchasePrice             int    `json:"PurchasePrice"`
		QuantityInStock           int    `json:"QuantityInStock"`
		ReservedQuantity          int    `json:"ReservedQuantity"`
		SalesAccount              int    `json:"SalesAccount"`
		StockGoods                bool   `json:"StockGoods"`
		StockPlace                string `json:"StockPlace"`
		StockValue                int    `json:"StockValue"`
		StockWarning              int    `json:"StockWarning"`
		SupplierName              string `json:"SupplierName"`
		SupplierNumber            string `json:"SupplierNumber"`
		Type                      string `json:"Type"`
		Unit                      string `json:"Unit"`
		VAT                       int    `json:"VAT"`
		WebshopArticle            bool   `json:"WebshopArticle"`
		Weight                    int    `json:"Weight"`
		Width                     int    `json:"Width"`
		Expired                   bool   `json:"Expired"`
		SalesPrice                int    `json:"SalesPrice"`
		CostCalculationMethod     string `json:"CostCalculationMethod"`
		StockAccount              int    `json:"StockAccount"`
		StockChangeAccount        int    `json:"StockChangeAccount"`
		DirectCost                int    `json:"DirectCost"`
		FreightCost               int    `json:"FreightCost"`
		OtherCost                 int    `json:"OtherCost"`
		DefaultStockPoint         string `json:"DefaultStockPoint"`
		DefaultStockLocation      string `json:"DefaultStockLocation"`
	} `json:"Article"`
}

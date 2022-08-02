package client

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

const (
	offersURI = "offers"
)

var offerStatusToDescription = map[GetAllOffersFilter]string{
	"cancelled":       "Retrieves all offers with the status 'cancelled'",
	"expired":         "Retrieves all offers that has been expired",
	"ordercreated":    "Retrieves all offers where an order has been created",
	"ordernotcreated": "Retrieves all offers where an order has not been created",
}

// GetAllOffers does _GET https://api.fortnox.se/3/offers/
//
// filter - GetAllOffersFilter
func (c Client) GetAllOffers(ctx context.Context, filter *GetAllOffersFilter) ([]Offer, error) {
	resp := &GetAllOffersResp{}

	params := filter.urlValues()

	err := c._GET(ctx, offersURI, params, resp)
	if err != nil {
		return nil, err
	}

	return resp.Offers, nil
}

// CreateOffer does _POST https://api.fortnox.se/3/offers/
//
// Errors that can be raised by this endpoint.
//
// Error 	Code	HTTP Code	Description	Solution
//
// 2004167	400		An account must be provided when using a custom VAT rate and EasyVat has been enabled. Supply each row which has a custom VAT rate with an account.
func (c *Client) CreateOffer(ctx context.Context, o *Offer) (*Offer, error) {
	req := &CreateOfferReq{Offer: *o}
	resp := &CreateOfferResp{}

	err := c._POST(ctx, offersURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Offer, nil
}

// GetOffer does _GET https://api.fortnox.se/3/offers/{DocumentNumber}
//
// documentNumber - identifies the offer
func (c *Client) GetOffer(ctx context.Context, documentNumber string) (*Offer, error) {
	resp := &GetOfferResp{}

	uri := fmt.Sprintf("%s/%s", offersURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Offer, nil
}

// UpdateOffer does _PUT https://api.fortnox.se/3/offers/{DocumentNumber}
//
// documentNumber - identifies the offer
//
// req - offer to update
func (c *Client) UpdateOffer(ctx context.Context, documentNumber string, o *Offer) (*Offer, error) {
	req := &UpdateOfferReq{Offer: *o}
	resp := &UpdateOfferResp{}

	uri := fmt.Sprintf("%s/%s", offersURI, documentNumber)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Offer, nil
}

// PrintOffer does _GET https://api.fortnox.se/3/offers/{DocumentNumber}/print
//
// documentNumber - identifies the offer
func (c *Client) PrintOffer(ctx context.Context, documentNumber string) (*[]byte, error) {
	resp := &[]byte{}

	uri := fmt.Sprintf("%s/%s/print", offersURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SendOfferAsEmail does _GET https://api.fortnox.se/3/offers/{DocumentNumber}/email
//
// documentNumber - identifies the offer
func (c *Client) SendOfferAsEmail(ctx context.Context, documentNumber string) (*Offer, error) {
	resp := &SendOfferAsEmailResp{}

	uri := fmt.Sprintf("%s/%s/email", offersURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Offer, nil
}

// PreviewOffer does _GET https://api.fortnox.se/3/offers/{DocumentNumber}/preview
//
// documentNumber - identifies the offer
//
// The difference between this and the print-endpoint is that property Sent is not set to TRUE.
func (c *Client) PreviewOffer(ctx context.Context, documentNumber string) (*[]byte, error) {
	resp := &[]byte{}

	uri := fmt.Sprintf("%s/%s/preview", offersURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateOrderOutOfOffer does _PUT https://api.fortnox.se/3/offers/{DocumentNumber}/createorder
//
// documentNumber - identifies the offer
func (c *Client) CreateOrderOutOfOffer(ctx context.Context, documentNumber string) (*Order, error) {
	resp := &CreateOrderOutOfOfferResp{}

	uri := fmt.Sprintf("%s/%s/createorder", offersURI, documentNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Order, nil
}

// CancelOffer does _PUT https://api.fortnox.se/3/offers/{DocumentNumber}/cancel
//
// documentNumber - identifies the offer
func (c *Client) CancelOffer(ctx context.Context, documentNumber string) (*Offer, error) {
	resp := &CancelOfferResp{}

	uri := fmt.Sprintf("%s/%s/cancel", offersURI, documentNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Offer, nil
}

// SetOfferAsSent does _PUT https://api.fortnox.se/3/offers/{DocumentNumber}/externalprint
//
// documentNumber - identifies the offer
//
// Use this endpoint to set offer as sent, without generating an offer.
func (c *Client) SetOfferAsSent(ctx context.Context, documentNumber string) (*Offer, error) {
	resp := &SetOfferAsSentResp{}

	uri := fmt.Sprintf("%s/%s/externalprint", offersURI, documentNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Offer, nil
}

type GetAllOffersFilter string

const (
	CancelledGetAllOffersFilter       GetAllOffersFilter = "cancelled"
	ExpiredGetAllOffersFilter         GetAllOffersFilter = "expired"
	OrderCreatedGetAllOffersFilter    GetAllOffersFilter = "ordercreated"
	OrderNotCreatedGetAllOffersFilter GetAllOffersFilter = "ordernotcreated"
)

func (f *GetAllOffersFilter) urlValues() url.Values {
	if f == nil {
		return nil
	}

	params := url.Values{}

	fStr := string(*f)
	if strings.TrimSpace(fStr) != "" {
		params["filter"] = []string{fStr}
	}

	return params
}

type GetAllOffersResp struct {
	Offers []Offer `json:"Offers"`
}

type CreateOfferReq struct {
	Offer Offer `json:"Offer"`
}

type CreateOfferResp struct {
	Offer Offer `json:"Offer"`
}

type GetOfferResp struct {
	Offer Offer `json:"Offer"`
}

type UpdateOfferReq struct {
	Offer Offer `json:"Offer"`
}

type UpdateOfferResp struct {
	Offer Offer `json:"Offer"`
}

type SendOfferAsEmailResp struct {
	Offer Offer `json:"Offer"`
}

type CreateOrderOutOfOfferResp struct {
	Order Order `json:"Order"`
}

type CancelOfferResp struct {
	Offer Offer `json:"Offer"`
}

type SetOfferAsSentResp struct {
	Offer Offer `json:"Offer"`
}

type OfferRow struct {
	AccountNumber          int    `json:"AccountNumber,omitempty"`
	ArticleNumber          string `json:"ArticleNumber,omitempty"`
	ContributionPercent    string `json:"ContributionPercent,omitempty"`
	ContributionValue      string `json:"ContributionValue,omitempty"`
	CostCenter             string `json:"CostCenter,omitempty"`
	Description            string `json:"Description,omitempty"`
	Discount               int    `json:"Discount,omitempty"`
	DiscountType           string `json:"DiscountType,omitempty"`
	HouseWork              bool   `json:"HouseWork,omitempty"`
	HouseWorkHoursToReport int    `json:"HouseWorkHoursToReport,omitempty"`
	HouseWorkType          string `json:"HouseWorkType,omitempty"`
	Price                  int    `json:"Price,omitempty"`
	Project                string `json:"Project,omitempty"`
	Quantity               string `json:"Quantity,omitempty"`
	RowId                  int    `json:"RowId,omitempty"`
	Total                  int    `json:"Total,omitempty"`
	Unit                   string `json:"Unit,omitempty"`
	VAT                    int    `json:"VAT,omitempty"`
}

type Offer struct {
	Url                  string           `json:"@url,omitempty"`
	UrlTaxReductionList  string           `json:"@urlTaxReductionList,omitempty"`
	AdministrationFee    int              `json:"AdministrationFee,omitempty"`
	AdministrationFeeVAT int              `json:"AdministrationFeeVAT,omitempty"`
	Address1             string           `json:"Address1,omitempty"`
	Address2             string           `json:"Address2,omitempty"`
	BasisTaxReduction    int              `json:"BasisTaxReduction,omitempty"`
	Cancelled            bool             `json:"Cancelled,omitempty"`
	City                 string           `json:"City,omitempty"`
	Comments             string           `json:"Comments,omitempty"`
	ContributionPercent  int              `json:"ContributionPercent,omitempty"`
	ContributionValue    int              `json:"ContributionValue,omitempty"`
	CopyRemarks          bool             `json:"CopyRemarks,omitempty"`
	Country              string           `json:"Country,omitempty"`
	CostCenter           string           `json:"CostCenter,omitempty"`
	Currency             string           `json:"Currency,omitempty"`
	CurrencyRate         int              `json:"CurrencyRate,omitempty"`
	CurrencyUnit         int              `json:"CurrencyUnit,omitempty"`
	CustomerName         string           `json:"CustomerName,omitempty"`
	CustomerNumber       string           `json:"CustomerNumber,omitempty"`
	DeliveryAddress1     string           `json:"DeliveryAddress1,omitempty"`
	DeliveryAddress2     string           `json:"DeliveryAddress2,omitempty"`
	DeliveryCity         string           `json:"DeliveryCity,omitempty"`
	DeliveryCountry      string           `json:"DeliveryCountry,omitempty"`
	DeliveryDate         string           `json:"DeliveryDate,omitempty"`
	DeliveryName         string           `json:"DeliveryName,omitempty"`
	DeliveryZipCode      string           `json:"DeliveryZipCode,omitempty"`
	DocumentNumber       string           `json:"DocumentNumber,omitempty"`
	EmailInformation     EmailInformation `json:"EmailInformation,omitempty"`
	ExpireDate           string           `json:"ExpireDate,omitempty"`
	Freight              int              `json:"Freight,omitempty"`
	FreightVAT           int              `json:"FreightVAT,omitempty"`
	Gross                int              `json:"Gross,omitempty"`
	HouseWork            bool             `json:"HouseWork,omitempty"`
	InvoiceReference     string           `json:"InvoiceReference,omitempty"`
	Labels               []Label          `json:"Labels,omitempty"`
	Language             string           `json:"Language,omitempty"`
	Net                  int              `json:"Net,omitempty"`
	NotCompleted         bool             `json:"NotCompleted,omitempty"`
	OfferDate            string           `json:"OfferDate,omitempty"`
	OfferRows            []OfferRow       `json:"OfferRows,omitempty"`
	OrderReference       string           `json:"OrderReference,omitempty"`
	OrganisationNumber   string           `json:"OrganisationNumber,omitempty"`
	OurReference         string           `json:"OurReference,omitempty"`
	Phone1               string           `json:"Phone1,omitempty"`
	Phone2               string           `json:"Phone2,omitempty"`
	PriceList            string           `json:"PriceList,omitempty"`
	PrintTemplate        string           `json:"PrintTemplate,omitempty"`
	Project              string           `json:"Project,omitempty"`
	Remarks              string           `json:"Remarks,omitempty"`
	RoundOff             int              `json:"RoundOff,omitempty"`
	Sent                 bool             `json:"Sent,omitempty"`
	TaxReduction         int              `json:"TaxReduction,omitempty"`
	TermsOfDelivery      string           `json:"TermsOfDelivery,omitempty"`
	TermsOfPayment       string           `json:"TermsOfPayment,omitempty"`
	Total                int              `json:"Total,omitempty"`
	TotalToPay           int              `json:"TotalToPay,omitempty"`
	TotalVAT             int              `json:"TotalVAT,omitempty"`
	VATIncluded          bool             `json:"VATIncluded,omitempty"`
	WayOfDelivery        string           `json:"WayOfDelivery,omitempty"`
	YourReference        string           `json:"YourReference,omitempty"`
	YourReferenceNumber  string           `json:"YourReferenceNumber,omitempty"`
	ZipCode              string           `json:"ZipCode,omitempty"`
	TaxReductionType     string           `json:"TaxReductionType,omitempty"`
}

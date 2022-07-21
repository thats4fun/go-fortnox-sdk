package client

import (
	"context"
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

// GetAllOffers does _GET
//
// filter - GetAllOffersFilter
func (c Client) GetAllOffers(ctx context.Context, filter GetAllOffersFilter) (*GetAllOffersResp, error) {
	resp := &GetAllOffersResp{}

	params := filter.urlValues()

	err := c._GET(ctx, offersURI, params, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetAllOffersFilter string

const (
	CancelledGetAllOffersFilter       GetAllOffersFilter = "cancelled"
	ExpiredGetAllOffersFilter         GetAllOffersFilter = "expired"
	OrderCreatedGetAllOffersFilter    GetAllOffersFilter = "ordercreated"
	OrderNotCreatedGetAllOffersFilter GetAllOffersFilter = "ordernotcreated"
)

func (f GetAllOffersFilter) urlValues() url.Values {
	params := url.Values{}

	fStr := string(f)
	if strings.TrimSpace(fStr) != "" {
		params["filter"] = []string{fStr}
	}

	return params
}

type GetAllOffersResp struct {
	Offers []struct {
		Url            string `json:"@url"`
		Cancelled      bool   `json:"Cancelled"`
		Currency       string `json:"Currency"`
		CustomerName   string `json:"CustomerName"`
		CustomerNumber string `json:"CustomerNumber"`
		DocumentNumber string `json:"DocumentNumber"`
		OfferDate      string `json:"OfferDate"`
		Project        string `json:"Project"`
		Sent           bool   `json:"Sent"`
		Total          int    `json:"Total"`
	} `json:"Offers"`
}

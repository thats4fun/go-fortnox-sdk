package client

import (
	"context"
	"fmt"
)

const (
	customerReferencesURI = "customerreferences"
)

// GetAllCustomerReferences does _GET https://api.fortnox.se/3/customerreferences/
func (c *Client) GetAllCustomerReferences(ctx context.Context) (*GetAllCustomerReferencesResp, error) {
	resp := &GetAllCustomerReferencesResp{}

	err := c._GET(ctx, customerReferencesURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateCustomerReference does _POST https://api.fortnox.se/3/customerreferences/
//
// req - customer reference row to create
func (c *Client) CreateCustomerReference(
	ctx context.Context,
	req *CreateCustomerReferenceReq) (*CreateCustomerReferenceResp, error) {

	resp := &CreateCustomerReferenceResp{}

	err := c._POST(ctx, customerReferencesURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetCustomerReference does _GET https://api.fortnox.se/3/customerreferences/{CustomerReferenceRowId}
//
// customerReferenceRowID - identifies the customer reference row
func (c *Client) GetCustomerReference(
	ctx context.Context,
	customerReferenceRowID string) (*GetCustomerReferenceResp, error) {

	resp := &GetCustomerReferenceResp{}

	uri := fmt.Sprintf("%s/%s", customerReferencesURI, customerReferenceRowID)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateCustomerReference does _PUT https://api.fortnox.se/3/customerreferences/{CustomerReferenceRowId}
//
// customerReferenceRowID - identifies the customer reference row
//
// req - customer reference row to update
func (c *Client) UpdateCustomerReference(
	ctx context.Context,
	customerReferenceRowID string,
	req *UpdateCustomerReferenceReq) (*UpdateCustomerReferenceResp, error) {

	resp := &UpdateCustomerReferenceResp{}

	uri := fmt.Sprintf("%s/%s", customersURI, customerReferenceRowID)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteCustomerReferenceRow does _DELETE https://api.fortnox.se/3/customerreferences/{CustomerReferenceRowId}
//
// customerReferenceRowID - identifies the customer reference row
func (c *Client) DeleteCustomerReferenceRow(ctx context.Context, customerReferenceRowID string) error {
	uri := fmt.Sprintf("%s/%s", customerReferencesURI, customerReferenceRowID)
	return c._DELETE(ctx, uri)
}

type GetAllCustomerReferencesResp struct {
	CustomerReference struct {
		CustomerReferenceRows []struct {
			Reference      string `json:"Reference"`
			CustomerNumber string `json:"CustomerNumber"`
			Id             int    `json:"Id"`
		} `json:"CustomerReferenceRows"`
	} `json:"CustomerReference"`
}

type CreateCustomerReferenceReq struct {
	CustomerReferenceRow struct {
		Reference      string `json:"Reference"`
		CustomerNumber string `json:"CustomerNumber"`
		Id             int    `json:"Id"`
	} `json:"CustomerReferenceRow"`
}

type CreateCustomerReferenceResp struct {
	CustomerReference struct {
		CustomerReferenceRows []struct {
			Reference      string `json:"Reference"`
			CustomerNumber string `json:"CustomerNumber"`
			Id             int    `json:"Id"`
		} `json:"CustomerReferenceRows"`
	} `json:"CustomerReference"`
}

type GetCustomerReferenceResp struct {
	CustomerReference struct {
		CustomerReferenceRows []struct {
			Reference      string `json:"Reference"`
			CustomerNumber string `json:"CustomerNumber"`
			Id             int    `json:"Id"`
		} `json:"CustomerReferenceRows"`
	} `json:"CustomerReference"`
}

type UpdateCustomerReferenceReq struct {
	CustomerReferenceRow struct {
		Reference      string `json:"Reference"`
		CustomerNumber string `json:"CustomerNumber"`
		Id             int    `json:"Id"`
	} `json:"CustomerReferenceRow"`
}

type UpdateCustomerReferenceResp struct {
	Customer struct {
		Url                  string `json:"@url"`
		Address1             string `json:"Address1"`
		Address2             string `json:"Address2"`
		City                 string `json:"City"`
		Country              string `json:"Country"`
		Comments             string `json:"Comments"`
		Currency             string `json:"Currency"`
		CostCenter           string `json:"CostCenter"`
		CountryCode          string `json:"CountryCode"`
		Active               bool   `json:"Active"`
		CustomerNumber       string `json:"CustomerNumber"`
		DefaultDeliveryTypes struct {
			Invoice string `json:"Invoice"`
			Order   string `json:"Order"`
			Offer   string `json:"Offer"`
		} `json:"DefaultDeliveryTypes"`
		DefaultTemplates struct {
			CashInvoice string `json:"CashInvoice"`
			Invoice     string `json:"Invoice"`
			Offer       string `json:"Offer"`
			Order       string `json:"Order"`
		} `json:"DefaultTemplates"`
		DeliveryAddress1         string `json:"DeliveryAddress1"`
		DeliveryAddress2         string `json:"DeliveryAddress2"`
		DeliveryCity             string `json:"DeliveryCity"`
		DeliveryCountry          string `json:"DeliveryCountry"`
		DeliveryCountryCode      string `json:"DeliveryCountryCode"`
		DeliveryFax              string `json:"DeliveryFax"`
		DeliveryName             string `json:"DeliveryName"`
		DeliveryPhone1           string `json:"DeliveryPhone1"`
		DeliveryPhone2           string `json:"DeliveryPhone2"`
		DeliveryZipCode          string `json:"DeliveryZipCode"`
		Email                    string `json:"Email"`
		EmailInvoice             string `json:"EmailInvoice"`
		EmailInvoiceBCC          string `json:"EmailInvoiceBCC"`
		EmailInvoiceCC           string `json:"EmailInvoiceCC"`
		EmailOffer               string `json:"EmailOffer"`
		EmailOfferBCC            string `json:"EmailOfferBCC"`
		EmailOfferCC             string `json:"EmailOfferCC"`
		EmailOrder               string `json:"EmailOrder"`
		EmailOrderBCC            string `json:"EmailOrderBCC"`
		EmailOrderCC             string `json:"EmailOrderCC"`
		ExternalReference        string `json:"ExternalReference"`
		Fax                      string `json:"Fax"`
		GLN                      string `json:"GLN"`
		GLNDelivery              string `json:"GLNDelivery"`
		InvoiceAdministrationFee string `json:"InvoiceAdministrationFee"`
		InvoiceDiscount          int    `json:"InvoiceDiscount"`
		InvoiceFreight           string `json:"InvoiceFreight"`
		InvoiceRemark            string `json:"InvoiceRemark"`
		Name                     string `json:"Name"`
		OrganisationNumber       string `json:"OrganisationNumber"`
		OurReference             string `json:"OurReference"`
		Phone1                   string `json:"Phone1"`
		Phone2                   string `json:"Phone2"`
		PriceList                string `json:"PriceList"`
		Project                  string `json:"Project"`
		SalesAccount             string `json:"SalesAccount"`
		ShowPriceVATIncluded     bool   `json:"ShowPriceVATIncluded"`
		TermsOfDelivery          string `json:"TermsOfDelivery"`
		TermsOfPayment           string `json:"TermsOfPayment"`
		Type                     string `json:"Type"`
		VATNumber                string `json:"VATNumber"`
		VATType                  string `json:"VATType"`
		VisitingAddress          string `json:"VisitingAddress"`
		VisitingCity             string `json:"VisitingCity"`
		VisitingCountry          string `json:"VisitingCountry"`
		VisitingCountryCode      string `json:"VisitingCountryCode"`
		VisitingZipCode          string `json:"VisitingZipCode"`
		WayOfDelivery            string `json:"WayOfDelivery"`
		WWW                      string `json:"WWW"`
		YourReference            string `json:"YourReference"`
		ZipCode                  string `json:"ZipCode"`
	} `json:"Customer"`
}

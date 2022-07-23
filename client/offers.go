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
func (c Client) GetAllOffers(ctx context.Context, filter *GetAllOffersFilter) (*GetAllOffersResp, error) {
	resp := &GetAllOffersResp{}

	params := filter.urlValues()

	err := c._GET(ctx, offersURI, params, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateOffer does _POST https://api.fortnox.se/3/offers/
//
// Errors that can be raised by this endpoint.
//
// Error 	Code	HTTP Code	Description	Solution
//
// 2004167	400		An account must be provided when using a custom VAT rate and EasyVat has been enabled. Supply each row which has a custom VAT rate with an account.
func (c *Client) CreateOffer(ctx context.Context, req *CreateOfferReq) (*CreateOfferResp, error) {
	resp := &CreateOfferResp{}

	err := c._POST(ctx, offersURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetOffer does _GET https://api.fortnox.se/3/offers/{DocumentNumber}
//
// documentNumber - identifies the offer
func (c *Client) GetOffer(ctx context.Context, documentNumber string) (*GetOfferResp, error) {
	resp := &GetOfferResp{}

	uri := fmt.Sprintf("%s/%s", offersURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateOffer does _PUT https://api.fortnox.se/3/offers/{DocumentNumber}
//
// documentNumber - identifies the offer
//
// req - offer to update
func (c *Client) UpdateOffer(ctx context.Context, documentNumber string, req *UpdateOfferReq) (*UpdateOfferResp, error) {
	resp := &UpdateOfferResp{}

	uri := fmt.Sprintf("%s/%s", offersURI, documentNumber)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
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
func (c *Client) SendOfferAsEmail(ctx context.Context, documentNumber string) (*SendOfferAsEmailResp, error) {
	resp := &SendOfferAsEmailResp{}

	uri := fmt.Sprintf("%s/%s/email", offersURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
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
func (c *Client) CreateOrderOutOfOffer(ctx context.Context, documentNumber string) (*CreateOrderOutOfOfferResp, error) {
	resp := &CreateOrderOutOfOfferResp{}

	uri := fmt.Sprintf("%s/%s/createorder", offersURI, documentNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CancelOffer does _PUT https://api.fortnox.se/3/offers/{DocumentNumber}/cancel
//
// documentNumber - identifies the offer
func (c *Client) CancelOffer(ctx context.Context, documentNumber string) (*CancelOfferResp, error) {
	resp := &CancelOfferResp{}

	uri := fmt.Sprintf("%s/%s/cancel", offersURI, documentNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SetOfferAsSent does _PUT https://api.fortnox.se/3/offers/{DocumentNumber}/externalprint
//
// documentNumber - identifies the offer
//
// Use this endpoint to set offer as sent, without generating an offer.
func (c *Client) SetOfferAsSent(ctx context.Context, documentNumber string) (*SetOfferAsSentResp, error) {
	resp := &SetOfferAsSentResp{}

	uri := fmt.Sprintf("%s/%s/externalprint", offersURI, documentNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
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

type CreateOfferReq struct {
	Offer struct {
		Url                  string `json:"@url"`
		UrlTaxReductionList  string `json:"@urlTaxReductionList"`
		AdministrationFee    int    `json:"AdministrationFee"`
		AdministrationFeeVAT int    `json:"AdministrationFeeVAT"`
		Address1             string `json:"Address1"`
		Address2             string `json:"Address2"`
		BasisTaxReduction    int    `json:"BasisTaxReduction"`
		Cancelled            bool   `json:"Cancelled"`
		City                 string `json:"City"`
		Comments             string `json:"Comments"`
		ContributionPercent  int    `json:"ContributionPercent"`
		ContributionValue    int    `json:"ContributionValue"`
		CopyRemarks          bool   `json:"CopyRemarks"`
		Country              string `json:"Country"`
		CostCenter           string `json:"CostCenter"`
		Currency             string `json:"Currency"`
		CurrencyRate         int    `json:"CurrencyRate"`
		CurrencyUnit         int    `json:"CurrencyUnit"`
		CustomerName         string `json:"CustomerName"`
		CustomerNumber       string `json:"CustomerNumber"`
		DeliveryAddress1     string `json:"DeliveryAddress1"`
		DeliveryAddress2     string `json:"DeliveryAddress2"`
		DeliveryCity         string `json:"DeliveryCity"`
		DeliveryCountry      string `json:"DeliveryCountry"`
		DeliveryDate         string `json:"DeliveryDate"`
		DeliveryName         string `json:"DeliveryName"`
		DeliveryZipCode      string `json:"DeliveryZipCode"`
		DocumentNumber       string `json:"DocumentNumber"`
		EmailInformation     struct {
			EmailAddressFrom string `json:"EmailAddressFrom"`
			EmailAddressTo   string `json:"EmailAddressTo"`
			EmailAddressCC   string `json:"EmailAddressCC"`
			EmailAddressBCC  string `json:"EmailAddressBCC"`
			EmailSubject     string `json:"EmailSubject"`
			EmailBody        string `json:"EmailBody"`
		} `json:"EmailInformation"`
		ExpireDate       string `json:"ExpireDate"`
		Freight          int    `json:"Freight"`
		FreightVAT       int    `json:"FreightVAT"`
		Gross            int    `json:"Gross"`
		HouseWork        bool   `json:"HouseWork"`
		InvoiceReference string `json:"InvoiceReference"`
		Labels           []struct {
			Id int `json:"Id"`
		} `json:"Labels"`
		Language     string `json:"Language"`
		Net          int    `json:"Net"`
		NotCompleted bool   `json:"NotCompleted"`
		OfferDate    string `json:"OfferDate"`
		OfferRows    []struct {
			AccountNumber          int    `json:"AccountNumber"`
			ArticleNumber          string `json:"ArticleNumber"`
			ContributionPercent    string `json:"ContributionPercent"`
			ContributionValue      string `json:"ContributionValue"`
			CostCenter             string `json:"CostCenter"`
			Description            string `json:"Description"`
			Discount               int    `json:"Discount"`
			DiscountType           string `json:"DiscountType"`
			HouseWork              bool   `json:"HouseWork"`
			HouseWorkHoursToReport int    `json:"HouseWorkHoursToReport"`
			HouseWorkType          string `json:"HouseWorkType"`
			Price                  int    `json:"Price"`
			Project                string `json:"Project"`
			Quantity               string `json:"Quantity"`
			RowId                  int    `json:"RowId"`
			Total                  int    `json:"Total"`
			Unit                   string `json:"Unit"`
			VAT                    int    `json:"VAT"`
		} `json:"OfferRows"`
		OrderReference      string `json:"OrderReference"`
		OrganisationNumber  string `json:"OrganisationNumber"`
		OurReference        string `json:"OurReference"`
		Phone1              string `json:"Phone1"`
		Phone2              string `json:"Phone2"`
		PriceList           string `json:"PriceList"`
		PrintTemplate       string `json:"PrintTemplate"`
		Project             string `json:"Project"`
		Remarks             string `json:"Remarks"`
		RoundOff            int    `json:"RoundOff"`
		Sent                bool   `json:"Sent"`
		TaxReduction        int    `json:"TaxReduction"`
		TermsOfDelivery     string `json:"TermsOfDelivery"`
		TermsOfPayment      string `json:"TermsOfPayment"`
		Total               int    `json:"Total"`
		TotalToPay          int    `json:"TotalToPay"`
		TotalVAT            int    `json:"TotalVAT"`
		VATIncluded         bool   `json:"VATIncluded"`
		WayOfDelivery       string `json:"WayOfDelivery"`
		YourReference       string `json:"YourReference"`
		YourReferenceNumber string `json:"YourReferenceNumber"`
		ZipCode             string `json:"ZipCode"`
		TaxReductionType    string `json:"TaxReductionType"`
	} `json:"Offer"`
}

type CreateOfferResp struct {
	Offer struct {
		Url                  string `json:"@url"`
		UrlTaxReductionList  string `json:"@urlTaxReductionList"`
		AdministrationFee    int    `json:"AdministrationFee"`
		AdministrationFeeVAT int    `json:"AdministrationFeeVAT"`
		Address1             string `json:"Address1"`
		Address2             string `json:"Address2"`
		BasisTaxReduction    int    `json:"BasisTaxReduction"`
		Cancelled            bool   `json:"Cancelled"`
		City                 string `json:"City"`
		Comments             string `json:"Comments"`
		ContributionPercent  int    `json:"ContributionPercent"`
		ContributionValue    int    `json:"ContributionValue"`
		CopyRemarks          bool   `json:"CopyRemarks"`
		Country              string `json:"Country"`
		CostCenter           string `json:"CostCenter"`
		Currency             string `json:"Currency"`
		CurrencyRate         int    `json:"CurrencyRate"`
		CurrencyUnit         int    `json:"CurrencyUnit"`
		CustomerName         string `json:"CustomerName"`
		CustomerNumber       string `json:"CustomerNumber"`
		DeliveryAddress1     string `json:"DeliveryAddress1"`
		DeliveryAddress2     string `json:"DeliveryAddress2"`
		DeliveryCity         string `json:"DeliveryCity"`
		DeliveryCountry      string `json:"DeliveryCountry"`
		DeliveryDate         string `json:"DeliveryDate"`
		DeliveryName         string `json:"DeliveryName"`
		DeliveryZipCode      string `json:"DeliveryZipCode"`
		DocumentNumber       string `json:"DocumentNumber"`
		EmailInformation     struct {
			EmailAddressFrom string `json:"EmailAddressFrom"`
			EmailAddressTo   string `json:"EmailAddressTo"`
			EmailAddressCC   string `json:"EmailAddressCC"`
			EmailAddressBCC  string `json:"EmailAddressBCC"`
			EmailSubject     string `json:"EmailSubject"`
			EmailBody        string `json:"EmailBody"`
		} `json:"EmailInformation"`
		ExpireDate       string `json:"ExpireDate"`
		Freight          int    `json:"Freight"`
		FreightVAT       int    `json:"FreightVAT"`
		Gross            int    `json:"Gross"`
		HouseWork        bool   `json:"HouseWork"`
		InvoiceReference string `json:"InvoiceReference"`
		Labels           []struct {
			Id int `json:"Id"`
		} `json:"Labels"`
		Language     string `json:"Language"`
		Net          int    `json:"Net"`
		NotCompleted bool   `json:"NotCompleted"`
		OfferDate    string `json:"OfferDate"`
		OfferRows    []struct {
			AccountNumber          int    `json:"AccountNumber"`
			ArticleNumber          string `json:"ArticleNumber"`
			ContributionPercent    string `json:"ContributionPercent"`
			ContributionValue      string `json:"ContributionValue"`
			CostCenter             string `json:"CostCenter"`
			Description            string `json:"Description"`
			Discount               int    `json:"Discount"`
			DiscountType           string `json:"DiscountType"`
			HouseWork              bool   `json:"HouseWork"`
			HouseWorkHoursToReport int    `json:"HouseWorkHoursToReport"`
			HouseWorkType          string `json:"HouseWorkType"`
			Price                  int    `json:"Price"`
			Project                string `json:"Project"`
			Quantity               string `json:"Quantity"`
			RowId                  int    `json:"RowId"`
			Total                  int    `json:"Total"`
			Unit                   string `json:"Unit"`
			VAT                    int    `json:"VAT"`
		} `json:"OfferRows"`
		OrderReference      string `json:"OrderReference"`
		OrganisationNumber  string `json:"OrganisationNumber"`
		OurReference        string `json:"OurReference"`
		Phone1              string `json:"Phone1"`
		Phone2              string `json:"Phone2"`
		PriceList           string `json:"PriceList"`
		PrintTemplate       string `json:"PrintTemplate"`
		Project             string `json:"Project"`
		Remarks             string `json:"Remarks"`
		RoundOff            int    `json:"RoundOff"`
		Sent                bool   `json:"Sent"`
		TaxReduction        int    `json:"TaxReduction"`
		TermsOfDelivery     string `json:"TermsOfDelivery"`
		TermsOfPayment      string `json:"TermsOfPayment"`
		Total               int    `json:"Total"`
		TotalToPay          int    `json:"TotalToPay"`
		TotalVAT            int    `json:"TotalVAT"`
		VATIncluded         bool   `json:"VATIncluded"`
		WayOfDelivery       string `json:"WayOfDelivery"`
		YourReference       string `json:"YourReference"`
		YourReferenceNumber string `json:"YourReferenceNumber"`
		ZipCode             string `json:"ZipCode"`
		TaxReductionType    string `json:"TaxReductionType"`
	} `json:"Offer"`
}

type GetOfferResp struct {
	Offer struct {
		Url                  string `json:"@url"`
		UrlTaxReductionList  string `json:"@urlTaxReductionList"`
		AdministrationFee    int    `json:"AdministrationFee"`
		AdministrationFeeVAT int    `json:"AdministrationFeeVAT"`
		Address1             string `json:"Address1"`
		Address2             string `json:"Address2"`
		BasisTaxReduction    int    `json:"BasisTaxReduction"`
		Cancelled            bool   `json:"Cancelled"`
		City                 string `json:"City"`
		Comments             string `json:"Comments"`
		ContributionPercent  int    `json:"ContributionPercent"`
		ContributionValue    int    `json:"ContributionValue"`
		CopyRemarks          bool   `json:"CopyRemarks"`
		Country              string `json:"Country"`
		CostCenter           string `json:"CostCenter"`
		Currency             string `json:"Currency"`
		CurrencyRate         int    `json:"CurrencyRate"`
		CurrencyUnit         int    `json:"CurrencyUnit"`
		CustomerName         string `json:"CustomerName"`
		CustomerNumber       string `json:"CustomerNumber"`
		DeliveryAddress1     string `json:"DeliveryAddress1"`
		DeliveryAddress2     string `json:"DeliveryAddress2"`
		DeliveryCity         string `json:"DeliveryCity"`
		DeliveryCountry      string `json:"DeliveryCountry"`
		DeliveryDate         string `json:"DeliveryDate"`
		DeliveryName         string `json:"DeliveryName"`
		DeliveryZipCode      string `json:"DeliveryZipCode"`
		DocumentNumber       string `json:"DocumentNumber"`
		EmailInformation     struct {
			EmailAddressFrom string `json:"EmailAddressFrom"`
			EmailAddressTo   string `json:"EmailAddressTo"`
			EmailAddressCC   string `json:"EmailAddressCC"`
			EmailAddressBCC  string `json:"EmailAddressBCC"`
			EmailSubject     string `json:"EmailSubject"`
			EmailBody        string `json:"EmailBody"`
		} `json:"EmailInformation"`
		ExpireDate       string `json:"ExpireDate"`
		Freight          int    `json:"Freight"`
		FreightVAT       int    `json:"FreightVAT"`
		Gross            int    `json:"Gross"`
		HouseWork        bool   `json:"HouseWork"`
		InvoiceReference string `json:"InvoiceReference"`
		Labels           []struct {
			Id int `json:"Id"`
		} `json:"Labels"`
		Language     string `json:"Language"`
		Net          int    `json:"Net"`
		NotCompleted bool   `json:"NotCompleted"`
		OfferDate    string `json:"OfferDate"`
		OfferRows    []struct {
			AccountNumber          int    `json:"AccountNumber"`
			ArticleNumber          string `json:"ArticleNumber"`
			ContributionPercent    string `json:"ContributionPercent"`
			ContributionValue      string `json:"ContributionValue"`
			CostCenter             string `json:"CostCenter"`
			Description            string `json:"Description"`
			Discount               int    `json:"Discount"`
			DiscountType           string `json:"DiscountType"`
			HouseWork              bool   `json:"HouseWork"`
			HouseWorkHoursToReport int    `json:"HouseWorkHoursToReport"`
			HouseWorkType          string `json:"HouseWorkType"`
			Price                  int    `json:"Price"`
			Project                string `json:"Project"`
			Quantity               string `json:"Quantity"`
			RowId                  int    `json:"RowId"`
			Total                  int    `json:"Total"`
			Unit                   string `json:"Unit"`
			VAT                    int    `json:"VAT"`
		} `json:"OfferRows"`
		OrderReference      string `json:"OrderReference"`
		OrganisationNumber  string `json:"OrganisationNumber"`
		OurReference        string `json:"OurReference"`
		Phone1              string `json:"Phone1"`
		Phone2              string `json:"Phone2"`
		PriceList           string `json:"PriceList"`
		PrintTemplate       string `json:"PrintTemplate"`
		Project             string `json:"Project"`
		Remarks             string `json:"Remarks"`
		RoundOff            int    `json:"RoundOff"`
		Sent                bool   `json:"Sent"`
		TaxReduction        int    `json:"TaxReduction"`
		TermsOfDelivery     string `json:"TermsOfDelivery"`
		TermsOfPayment      string `json:"TermsOfPayment"`
		Total               int    `json:"Total"`
		TotalToPay          int    `json:"TotalToPay"`
		TotalVAT            int    `json:"TotalVAT"`
		VATIncluded         bool   `json:"VATIncluded"`
		WayOfDelivery       string `json:"WayOfDelivery"`
		YourReference       string `json:"YourReference"`
		YourReferenceNumber string `json:"YourReferenceNumber"`
		ZipCode             string `json:"ZipCode"`
		TaxReductionType    string `json:"TaxReductionType"`
	} `json:"Offer"`
}

type UpdateOfferReq struct {
	Offer struct {
		Url                  string `json:"@url"`
		UrlTaxReductionList  string `json:"@urlTaxReductionList"`
		AdministrationFee    int    `json:"AdministrationFee"`
		AdministrationFeeVAT int    `json:"AdministrationFeeVAT"`
		Address1             string `json:"Address1"`
		Address2             string `json:"Address2"`
		BasisTaxReduction    int    `json:"BasisTaxReduction"`
		Cancelled            bool   `json:"Cancelled"`
		City                 string `json:"City"`
		Comments             string `json:"Comments"`
		ContributionPercent  int    `json:"ContributionPercent"`
		ContributionValue    int    `json:"ContributionValue"`
		CopyRemarks          bool   `json:"CopyRemarks"`
		Country              string `json:"Country"`
		CostCenter           string `json:"CostCenter"`
		Currency             string `json:"Currency"`
		CurrencyRate         int    `json:"CurrencyRate"`
		CurrencyUnit         int    `json:"CurrencyUnit"`
		CustomerName         string `json:"CustomerName"`
		CustomerNumber       string `json:"CustomerNumber"`
		DeliveryAddress1     string `json:"DeliveryAddress1"`
		DeliveryAddress2     string `json:"DeliveryAddress2"`
		DeliveryCity         string `json:"DeliveryCity"`
		DeliveryCountry      string `json:"DeliveryCountry"`
		DeliveryDate         string `json:"DeliveryDate"`
		DeliveryName         string `json:"DeliveryName"`
		DeliveryZipCode      string `json:"DeliveryZipCode"`
		DocumentNumber       string `json:"DocumentNumber"`
		EmailInformation     struct {
			EmailAddressFrom string `json:"EmailAddressFrom"`
			EmailAddressTo   string `json:"EmailAddressTo"`
			EmailAddressCC   string `json:"EmailAddressCC"`
			EmailAddressBCC  string `json:"EmailAddressBCC"`
			EmailSubject     string `json:"EmailSubject"`
			EmailBody        string `json:"EmailBody"`
		} `json:"EmailInformation"`
		ExpireDate       string `json:"ExpireDate"`
		Freight          int    `json:"Freight"`
		FreightVAT       int    `json:"FreightVAT"`
		Gross            int    `json:"Gross"`
		HouseWork        bool   `json:"HouseWork"`
		InvoiceReference string `json:"InvoiceReference"`
		Labels           []struct {
			Id int `json:"Id"`
		} `json:"Labels"`
		Language     string `json:"Language"`
		Net          int    `json:"Net"`
		NotCompleted bool   `json:"NotCompleted"`
		OfferDate    string `json:"OfferDate"`
		OfferRows    []struct {
			AccountNumber          int    `json:"AccountNumber"`
			ArticleNumber          string `json:"ArticleNumber"`
			ContributionPercent    string `json:"ContributionPercent"`
			ContributionValue      string `json:"ContributionValue"`
			CostCenter             string `json:"CostCenter"`
			Description            string `json:"Description"`
			Discount               int    `json:"Discount"`
			DiscountType           string `json:"DiscountType"`
			HouseWork              bool   `json:"HouseWork"`
			HouseWorkHoursToReport int    `json:"HouseWorkHoursToReport"`
			HouseWorkType          string `json:"HouseWorkType"`
			Price                  int    `json:"Price"`
			Project                string `json:"Project"`
			Quantity               string `json:"Quantity"`
			RowId                  int    `json:"RowId"`
			Total                  int    `json:"Total"`
			Unit                   string `json:"Unit"`
			VAT                    int    `json:"VAT"`
		} `json:"OfferRows"`
		OrderReference      string `json:"OrderReference"`
		OrganisationNumber  string `json:"OrganisationNumber"`
		OurReference        string `json:"OurReference"`
		Phone1              string `json:"Phone1"`
		Phone2              string `json:"Phone2"`
		PriceList           string `json:"PriceList"`
		PrintTemplate       string `json:"PrintTemplate"`
		Project             string `json:"Project"`
		Remarks             string `json:"Remarks"`
		RoundOff            int    `json:"RoundOff"`
		Sent                bool   `json:"Sent"`
		TaxReduction        int    `json:"TaxReduction"`
		TermsOfDelivery     string `json:"TermsOfDelivery"`
		TermsOfPayment      string `json:"TermsOfPayment"`
		Total               int    `json:"Total"`
		TotalToPay          int    `json:"TotalToPay"`
		TotalVAT            int    `json:"TotalVAT"`
		VATIncluded         bool   `json:"VATIncluded"`
		WayOfDelivery       string `json:"WayOfDelivery"`
		YourReference       string `json:"YourReference"`
		YourReferenceNumber string `json:"YourReferenceNumber"`
		ZipCode             string `json:"ZipCode"`
		TaxReductionType    string `json:"TaxReductionType"`
	} `json:"Offer"`
}

type UpdateOfferResp struct {
	Offer struct {
		Url                  string `json:"@url"`
		UrlTaxReductionList  string `json:"@urlTaxReductionList"`
		AdministrationFee    int    `json:"AdministrationFee"`
		AdministrationFeeVAT int    `json:"AdministrationFeeVAT"`
		Address1             string `json:"Address1"`
		Address2             string `json:"Address2"`
		BasisTaxReduction    int    `json:"BasisTaxReduction"`
		Cancelled            bool   `json:"Cancelled"`
		City                 string `json:"City"`
		Comments             string `json:"Comments"`
		ContributionPercent  int    `json:"ContributionPercent"`
		ContributionValue    int    `json:"ContributionValue"`
		CopyRemarks          bool   `json:"CopyRemarks"`
		Country              string `json:"Country"`
		CostCenter           string `json:"CostCenter"`
		Currency             string `json:"Currency"`
		CurrencyRate         int    `json:"CurrencyRate"`
		CurrencyUnit         int    `json:"CurrencyUnit"`
		CustomerName         string `json:"CustomerName"`
		CustomerNumber       string `json:"CustomerNumber"`
		DeliveryAddress1     string `json:"DeliveryAddress1"`
		DeliveryAddress2     string `json:"DeliveryAddress2"`
		DeliveryCity         string `json:"DeliveryCity"`
		DeliveryCountry      string `json:"DeliveryCountry"`
		DeliveryDate         string `json:"DeliveryDate"`
		DeliveryName         string `json:"DeliveryName"`
		DeliveryZipCode      string `json:"DeliveryZipCode"`
		DocumentNumber       string `json:"DocumentNumber"`
		EmailInformation     struct {
			EmailAddressFrom string `json:"EmailAddressFrom"`
			EmailAddressTo   string `json:"EmailAddressTo"`
			EmailAddressCC   string `json:"EmailAddressCC"`
			EmailAddressBCC  string `json:"EmailAddressBCC"`
			EmailSubject     string `json:"EmailSubject"`
			EmailBody        string `json:"EmailBody"`
		} `json:"EmailInformation"`
		ExpireDate       string `json:"ExpireDate"`
		Freight          int    `json:"Freight"`
		FreightVAT       int    `json:"FreightVAT"`
		Gross            int    `json:"Gross"`
		HouseWork        bool   `json:"HouseWork"`
		InvoiceReference string `json:"InvoiceReference"`
		Labels           []struct {
			Id int `json:"Id"`
		} `json:"Labels"`
		Language     string `json:"Language"`
		Net          int    `json:"Net"`
		NotCompleted bool   `json:"NotCompleted"`
		OfferDate    string `json:"OfferDate"`
		OfferRows    []struct {
			AccountNumber          int    `json:"AccountNumber"`
			ArticleNumber          string `json:"ArticleNumber"`
			ContributionPercent    string `json:"ContributionPercent"`
			ContributionValue      string `json:"ContributionValue"`
			CostCenter             string `json:"CostCenter"`
			Description            string `json:"Description"`
			Discount               int    `json:"Discount"`
			DiscountType           string `json:"DiscountType"`
			HouseWork              bool   `json:"HouseWork"`
			HouseWorkHoursToReport int    `json:"HouseWorkHoursToReport"`
			HouseWorkType          string `json:"HouseWorkType"`
			Price                  int    `json:"Price"`
			Project                string `json:"Project"`
			Quantity               string `json:"Quantity"`
			RowId                  int    `json:"RowId"`
			Total                  int    `json:"Total"`
			Unit                   string `json:"Unit"`
			VAT                    int    `json:"VAT"`
		} `json:"OfferRows"`
		OrderReference      string `json:"OrderReference"`
		OrganisationNumber  string `json:"OrganisationNumber"`
		OurReference        string `json:"OurReference"`
		Phone1              string `json:"Phone1"`
		Phone2              string `json:"Phone2"`
		PriceList           string `json:"PriceList"`
		PrintTemplate       string `json:"PrintTemplate"`
		Project             string `json:"Project"`
		Remarks             string `json:"Remarks"`
		RoundOff            int    `json:"RoundOff"`
		Sent                bool   `json:"Sent"`
		TaxReduction        int    `json:"TaxReduction"`
		TermsOfDelivery     string `json:"TermsOfDelivery"`
		TermsOfPayment      string `json:"TermsOfPayment"`
		Total               int    `json:"Total"`
		TotalToPay          int    `json:"TotalToPay"`
		TotalVAT            int    `json:"TotalVAT"`
		VATIncluded         bool   `json:"VATIncluded"`
		WayOfDelivery       string `json:"WayOfDelivery"`
		YourReference       string `json:"YourReference"`
		YourReferenceNumber string `json:"YourReferenceNumber"`
		ZipCode             string `json:"ZipCode"`
		TaxReductionType    string `json:"TaxReductionType"`
	} `json:"Offer"`
}

type SendOfferAsEmailResp struct {
	Offer struct {
		Url                  string `json:"@url"`
		UrlTaxReductionList  string `json:"@urlTaxReductionList"`
		AdministrationFee    int    `json:"AdministrationFee"`
		AdministrationFeeVAT int    `json:"AdministrationFeeVAT"`
		Address1             string `json:"Address1"`
		Address2             string `json:"Address2"`
		BasisTaxReduction    int    `json:"BasisTaxReduction"`
		Cancelled            bool   `json:"Cancelled"`
		City                 string `json:"City"`
		Comments             string `json:"Comments"`
		ContributionPercent  int    `json:"ContributionPercent"`
		ContributionValue    int    `json:"ContributionValue"`
		CopyRemarks          bool   `json:"CopyRemarks"`
		Country              string `json:"Country"`
		CostCenter           string `json:"CostCenter"`
		Currency             string `json:"Currency"`
		CurrencyRate         int    `json:"CurrencyRate"`
		CurrencyUnit         int    `json:"CurrencyUnit"`
		CustomerName         string `json:"CustomerName"`
		CustomerNumber       string `json:"CustomerNumber"`
		DeliveryAddress1     string `json:"DeliveryAddress1"`
		DeliveryAddress2     string `json:"DeliveryAddress2"`
		DeliveryCity         string `json:"DeliveryCity"`
		DeliveryCountry      string `json:"DeliveryCountry"`
		DeliveryDate         string `json:"DeliveryDate"`
		DeliveryName         string `json:"DeliveryName"`
		DeliveryZipCode      string `json:"DeliveryZipCode"`
		DocumentNumber       string `json:"DocumentNumber"`
		EmailInformation     struct {
			EmailAddressFrom string `json:"EmailAddressFrom"`
			EmailAddressTo   string `json:"EmailAddressTo"`
			EmailAddressCC   string `json:"EmailAddressCC"`
			EmailAddressBCC  string `json:"EmailAddressBCC"`
			EmailSubject     string `json:"EmailSubject"`
			EmailBody        string `json:"EmailBody"`
		} `json:"EmailInformation"`
		ExpireDate       string `json:"ExpireDate"`
		Freight          int    `json:"Freight"`
		FreightVAT       int    `json:"FreightVAT"`
		Gross            int    `json:"Gross"`
		HouseWork        bool   `json:"HouseWork"`
		InvoiceReference string `json:"InvoiceReference"`
		Labels           []struct {
			Id int `json:"Id"`
		} `json:"Labels"`
		Language     string `json:"Language"`
		Net          int    `json:"Net"`
		NotCompleted bool   `json:"NotCompleted"`
		OfferDate    string `json:"OfferDate"`
		OfferRows    []struct {
			AccountNumber          int    `json:"AccountNumber"`
			ArticleNumber          string `json:"ArticleNumber"`
			ContributionPercent    string `json:"ContributionPercent"`
			ContributionValue      string `json:"ContributionValue"`
			CostCenter             string `json:"CostCenter"`
			Description            string `json:"Description"`
			Discount               int    `json:"Discount"`
			DiscountType           string `json:"DiscountType"`
			HouseWork              bool   `json:"HouseWork"`
			HouseWorkHoursToReport int    `json:"HouseWorkHoursToReport"`
			HouseWorkType          string `json:"HouseWorkType"`
			Price                  int    `json:"Price"`
			Project                string `json:"Project"`
			Quantity               string `json:"Quantity"`
			RowId                  int    `json:"RowId"`
			Total                  int    `json:"Total"`
			Unit                   string `json:"Unit"`
			VAT                    int    `json:"VAT"`
		} `json:"OfferRows"`
		OrderReference      string `json:"OrderReference"`
		OrganisationNumber  string `json:"OrganisationNumber"`
		OurReference        string `json:"OurReference"`
		Phone1              string `json:"Phone1"`
		Phone2              string `json:"Phone2"`
		PriceList           string `json:"PriceList"`
		PrintTemplate       string `json:"PrintTemplate"`
		Project             string `json:"Project"`
		Remarks             string `json:"Remarks"`
		RoundOff            int    `json:"RoundOff"`
		Sent                bool   `json:"Sent"`
		TaxReduction        int    `json:"TaxReduction"`
		TermsOfDelivery     string `json:"TermsOfDelivery"`
		TermsOfPayment      string `json:"TermsOfPayment"`
		Total               int    `json:"Total"`
		TotalToPay          int    `json:"TotalToPay"`
		TotalVAT            int    `json:"TotalVAT"`
		VATIncluded         bool   `json:"VATIncluded"`
		WayOfDelivery       string `json:"WayOfDelivery"`
		YourReference       string `json:"YourReference"`
		YourReferenceNumber string `json:"YourReferenceNumber"`
		ZipCode             string `json:"ZipCode"`
		TaxReductionType    string `json:"TaxReductionType"`
	} `json:"Offer"`
}

type CreateOrderOutOfOfferResp struct {
	Order struct {
		Url                  string `json:"@url"`
		UrlTaxReductionList  string `json:"@urlTaxReductionList"`
		AdministrationFee    int    `json:"AdministrationFee"`
		AdministrationFeeVAT int    `json:"AdministrationFeeVAT"`
		Address1             string `json:"Address1"`
		Address2             string `json:"Address2"`
		BasisTaxReduction    int    `json:"BasisTaxReduction"`
		Cancelled            bool   `json:"Cancelled"`
		City                 string `json:"City"`
		Comments             string `json:"Comments"`
		ContributionPercent  int    `json:"ContributionPercent"`
		ContributionValue    int    `json:"ContributionValue"`
		CopyRemarks          bool   `json:"CopyRemarks"`
		Country              string `json:"Country"`
		CostCenter           string `json:"CostCenter"`
		Currency             string `json:"Currency"`
		CurrencyRate         int    `json:"CurrencyRate"`
		CurrencyUnit         int    `json:"CurrencyUnit"`
		CustomerName         string `json:"CustomerName"`
		CustomerNumber       string `json:"CustomerNumber"`
		DeliveryState        string `json:"DeliveryState"`
		DeliveryAddress1     string `json:"DeliveryAddress1"`
		DeliveryAddress2     string `json:"DeliveryAddress2"`
		DeliveryCity         string `json:"DeliveryCity"`
		DeliveryCountry      string `json:"DeliveryCountry"`
		DeliveryDate         string `json:"DeliveryDate"`
		DeliveryName         string `json:"DeliveryName"`
		DeliveryZipCode      string `json:"DeliveryZipCode"`
		DocumentNumber       string `json:"DocumentNumber"`
		EmailInformation     struct {
			EmailAddressFrom string `json:"EmailAddressFrom"`
			EmailAddressTo   string `json:"EmailAddressTo"`
			EmailAddressCC   string `json:"EmailAddressCC"`
			EmailAddressBCC  string `json:"EmailAddressBCC"`
			EmailSubject     string `json:"EmailSubject"`
			EmailBody        string `json:"EmailBody"`
		} `json:"EmailInformation"`
		ExternalInvoiceReference1 string `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string `json:"ExternalInvoiceReference2"`
		Freight                   int    `json:"Freight"`
		FreightVAT                int    `json:"FreightVAT"`
		Gross                     int    `json:"Gross"`
		HouseWork                 bool   `json:"HouseWork"`
		InvoiceReference          string `json:"InvoiceReference"`
		Labels                    []struct {
			Id int `json:"Id"`
		} `json:"Labels"`
		Language       string `json:"Language"`
		Net            int    `json:"Net"`
		NotCompleted   bool   `json:"NotCompleted"`
		OfferReference string `json:"OfferReference"`
		OrderDate      string `json:"OrderDate"`
		OrderRows      []struct {
			AccountNumber          int    `json:"AccountNumber"`
			ArticleNumber          string `json:"ArticleNumber"`
			ContributionPercent    string `json:"ContributionPercent"`
			ContributionValue      string `json:"ContributionValue"`
			CostCenter             string `json:"CostCenter"`
			DeliveredQuantity      string `json:"DeliveredQuantity"`
			Description            string `json:"Description"`
			Discount               int    `json:"Discount"`
			DiscountType           string `json:"DiscountType"`
			HouseWork              bool   `json:"HouseWork"`
			HouseWorkHoursToReport int    `json:"HouseWorkHoursToReport"`
			HouseWorkType          string `json:"HouseWorkType"`
			OrderedQuantity        string `json:"OrderedQuantity"`
			Price                  int    `json:"Price"`
			Project                string `json:"Project"`
			ReservedQuantity       string `json:"ReservedQuantity"`
			RowId                  int    `json:"RowId"`
			StockPointCode         string `json:"StockPointCode"`
			StockPointId           string `json:"StockPointId"`
			Total                  int    `json:"Total"`
			Unit                   string `json:"Unit"`
			VAT                    int    `json:"VAT"`
		} `json:"OrderRows"`
		OrderType          string `json:"OrderType"`
		OrganisationNumber string `json:"OrganisationNumber"`
		OurReference       string `json:"OurReference"`
		Phone1             string `json:"Phone1"`
		Phone2             string `json:"Phone2"`
		PriceList          string `json:"PriceList"`
		PrintTemplate      string `json:"PrintTemplate"`
		Project            string `json:"Project"`
		WarehouseReady     bool   `json:"WarehouseReady"`
		OutboundDate       string `json:"OutboundDate"`
		Remarks            string `json:"Remarks"`
		RoundOff           int    `json:"RoundOff"`
		Sent               bool   `json:"Sent"`
		TaxReduction       int    `json:"TaxReduction"`
		TermsOfDelivery    string `json:"TermsOfDelivery"`
		TermsOfPayment     string `json:"TermsOfPayment"`
		TimeBasisReference int    `json:"TimeBasisReference"`
		Total              int    `json:"Total"`
		TotalToPay         int    `json:"TotalToPay"`
		TotalVAT           int    `json:"TotalVAT"`
		VATIncluded        bool   `json:"VATIncluded"`
		WayOfDelivery      string `json:"WayOfDelivery"`
		YourReference      string `json:"YourReference"`
		YourOrderNumber    string `json:"YourOrderNumber"`
		ZipCode            string `json:"ZipCode"`
		StockPointCode     string `json:"StockPointCode"`
		StockPointId       string `json:"StockPointId"`
		TaxReductionType   string `json:"TaxReductionType"`
	} `json:"Order"`
}

type CancelOfferResp struct {
	Offer struct {
		Url                  string `json:"@url"`
		UrlTaxReductionList  string `json:"@urlTaxReductionList"`
		AdministrationFee    int    `json:"AdministrationFee"`
		AdministrationFeeVAT int    `json:"AdministrationFeeVAT"`
		Address1             string `json:"Address1"`
		Address2             string `json:"Address2"`
		BasisTaxReduction    int    `json:"BasisTaxReduction"`
		Cancelled            bool   `json:"Cancelled"`
		City                 string `json:"City"`
		Comments             string `json:"Comments"`
		ContributionPercent  int    `json:"ContributionPercent"`
		ContributionValue    int    `json:"ContributionValue"`
		CopyRemarks          bool   `json:"CopyRemarks"`
		Country              string `json:"Country"`
		CostCenter           string `json:"CostCenter"`
		Currency             string `json:"Currency"`
		CurrencyRate         int    `json:"CurrencyRate"`
		CurrencyUnit         int    `json:"CurrencyUnit"`
		CustomerName         string `json:"CustomerName"`
		CustomerNumber       string `json:"CustomerNumber"`
		DeliveryAddress1     string `json:"DeliveryAddress1"`
		DeliveryAddress2     string `json:"DeliveryAddress2"`
		DeliveryCity         string `json:"DeliveryCity"`
		DeliveryCountry      string `json:"DeliveryCountry"`
		DeliveryDate         string `json:"DeliveryDate"`
		DeliveryName         string `json:"DeliveryName"`
		DeliveryZipCode      string `json:"DeliveryZipCode"`
		DocumentNumber       string `json:"DocumentNumber"`
		EmailInformation     struct {
			EmailAddressFrom string `json:"EmailAddressFrom"`
			EmailAddressTo   string `json:"EmailAddressTo"`
			EmailAddressCC   string `json:"EmailAddressCC"`
			EmailAddressBCC  string `json:"EmailAddressBCC"`
			EmailSubject     string `json:"EmailSubject"`
			EmailBody        string `json:"EmailBody"`
		} `json:"EmailInformation"`
		ExpireDate       string `json:"ExpireDate"`
		Freight          int    `json:"Freight"`
		FreightVAT       int    `json:"FreightVAT"`
		Gross            int    `json:"Gross"`
		HouseWork        bool   `json:"HouseWork"`
		InvoiceReference string `json:"InvoiceReference"`
		Labels           []struct {
			Id int `json:"Id"`
		} `json:"Labels"`
		Language     string `json:"Language"`
		Net          int    `json:"Net"`
		NotCompleted bool   `json:"NotCompleted"`
		OfferDate    string `json:"OfferDate"`
		OfferRows    []struct {
			AccountNumber          int    `json:"AccountNumber"`
			ArticleNumber          string `json:"ArticleNumber"`
			ContributionPercent    string `json:"ContributionPercent"`
			ContributionValue      string `json:"ContributionValue"`
			CostCenter             string `json:"CostCenter"`
			Description            string `json:"Description"`
			Discount               int    `json:"Discount"`
			DiscountType           string `json:"DiscountType"`
			HouseWork              bool   `json:"HouseWork"`
			HouseWorkHoursToReport int    `json:"HouseWorkHoursToReport"`
			HouseWorkType          string `json:"HouseWorkType"`
			Price                  int    `json:"Price"`
			Project                string `json:"Project"`
			Quantity               string `json:"Quantity"`
			RowId                  int    `json:"RowId"`
			Total                  int    `json:"Total"`
			Unit                   string `json:"Unit"`
			VAT                    int    `json:"VAT"`
		} `json:"OfferRows"`
		OrderReference      string `json:"OrderReference"`
		OrganisationNumber  string `json:"OrganisationNumber"`
		OurReference        string `json:"OurReference"`
		Phone1              string `json:"Phone1"`
		Phone2              string `json:"Phone2"`
		PriceList           string `json:"PriceList"`
		PrintTemplate       string `json:"PrintTemplate"`
		Project             string `json:"Project"`
		Remarks             string `json:"Remarks"`
		RoundOff            int    `json:"RoundOff"`
		Sent                bool   `json:"Sent"`
		TaxReduction        int    `json:"TaxReduction"`
		TermsOfDelivery     string `json:"TermsOfDelivery"`
		TermsOfPayment      string `json:"TermsOfPayment"`
		Total               int    `json:"Total"`
		TotalToPay          int    `json:"TotalToPay"`
		TotalVAT            int    `json:"TotalVAT"`
		VATIncluded         bool   `json:"VATIncluded"`
		WayOfDelivery       string `json:"WayOfDelivery"`
		YourReference       string `json:"YourReference"`
		YourReferenceNumber string `json:"YourReferenceNumber"`
		ZipCode             string `json:"ZipCode"`
		TaxReductionType    string `json:"TaxReductionType"`
	} `json:"Offer"`
}

type SetOfferAsSentResp struct {
	Offer struct {
		Url                  string `json:"@url"`
		UrlTaxReductionList  string `json:"@urlTaxReductionList"`
		AdministrationFee    int    `json:"AdministrationFee"`
		AdministrationFeeVAT int    `json:"AdministrationFeeVAT"`
		Address1             string `json:"Address1"`
		Address2             string `json:"Address2"`
		BasisTaxReduction    int    `json:"BasisTaxReduction"`
		Cancelled            bool   `json:"Cancelled"`
		City                 string `json:"City"`
		Comments             string `json:"Comments"`
		ContributionPercent  int    `json:"ContributionPercent"`
		ContributionValue    int    `json:"ContributionValue"`
		CopyRemarks          bool   `json:"CopyRemarks"`
		Country              string `json:"Country"`
		CostCenter           string `json:"CostCenter"`
		Currency             string `json:"Currency"`
		CurrencyRate         int    `json:"CurrencyRate"`
		CurrencyUnit         int    `json:"CurrencyUnit"`
		CustomerName         string `json:"CustomerName"`
		CustomerNumber       string `json:"CustomerNumber"`
		DeliveryAddress1     string `json:"DeliveryAddress1"`
		DeliveryAddress2     string `json:"DeliveryAddress2"`
		DeliveryCity         string `json:"DeliveryCity"`
		DeliveryCountry      string `json:"DeliveryCountry"`
		DeliveryDate         string `json:"DeliveryDate"`
		DeliveryName         string `json:"DeliveryName"`
		DeliveryZipCode      string `json:"DeliveryZipCode"`
		DocumentNumber       string `json:"DocumentNumber"`
		EmailInformation     struct {
			EmailAddressFrom string `json:"EmailAddressFrom"`
			EmailAddressTo   string `json:"EmailAddressTo"`
			EmailAddressCC   string `json:"EmailAddressCC"`
			EmailAddressBCC  string `json:"EmailAddressBCC"`
			EmailSubject     string `json:"EmailSubject"`
			EmailBody        string `json:"EmailBody"`
		} `json:"EmailInformation"`
		ExpireDate       string `json:"ExpireDate"`
		Freight          int    `json:"Freight"`
		FreightVAT       int    `json:"FreightVAT"`
		Gross            int    `json:"Gross"`
		HouseWork        bool   `json:"HouseWork"`
		InvoiceReference string `json:"InvoiceReference"`
		Labels           []struct {
			Id int `json:"Id"`
		} `json:"Labels"`
		Language     string `json:"Language"`
		Net          int    `json:"Net"`
		NotCompleted bool   `json:"NotCompleted"`
		OfferDate    string `json:"OfferDate"`
		OfferRows    []struct {
			AccountNumber          int    `json:"AccountNumber"`
			ArticleNumber          string `json:"ArticleNumber"`
			ContributionPercent    string `json:"ContributionPercent"`
			ContributionValue      string `json:"ContributionValue"`
			CostCenter             string `json:"CostCenter"`
			Description            string `json:"Description"`
			Discount               int    `json:"Discount"`
			DiscountType           string `json:"DiscountType"`
			HouseWork              bool   `json:"HouseWork"`
			HouseWorkHoursToReport int    `json:"HouseWorkHoursToReport"`
			HouseWorkType          string `json:"HouseWorkType"`
			Price                  int    `json:"Price"`
			Project                string `json:"Project"`
			Quantity               string `json:"Quantity"`
			RowId                  int    `json:"RowId"`
			Total                  int    `json:"Total"`
			Unit                   string `json:"Unit"`
			VAT                    int    `json:"VAT"`
		} `json:"OfferRows"`
		OrderReference      string `json:"OrderReference"`
		OrganisationNumber  string `json:"OrganisationNumber"`
		OurReference        string `json:"OurReference"`
		Phone1              string `json:"Phone1"`
		Phone2              string `json:"Phone2"`
		PriceList           string `json:"PriceList"`
		PrintTemplate       string `json:"PrintTemplate"`
		Project             string `json:"Project"`
		Remarks             string `json:"Remarks"`
		RoundOff            int    `json:"RoundOff"`
		Sent                bool   `json:"Sent"`
		TaxReduction        int    `json:"TaxReduction"`
		TermsOfDelivery     string `json:"TermsOfDelivery"`
		TermsOfPayment      string `json:"TermsOfPayment"`
		Total               int    `json:"Total"`
		TotalToPay          int    `json:"TotalToPay"`
		TotalVAT            int    `json:"TotalVAT"`
		VATIncluded         bool   `json:"VATIncluded"`
		WayOfDelivery       string `json:"WayOfDelivery"`
		YourReference       string `json:"YourReference"`
		YourReferenceNumber string `json:"YourReferenceNumber"`
		ZipCode             string `json:"ZipCode"`
		TaxReductionType    string `json:"TaxReductionType"`
	} `json:"Offer"`
}

package client

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

const (
	ordersURI = "orders"
)

// GetAllOrders does _GET https://api.fortnox.se/3/orders/
//
// filter - GetAllOffersFilter
func (c Client) GetAllOrders(ctx context.Context, filter *GetAllOrdersFilter) (*GetAllOrdersResp, error) {
	resp := &GetAllOrdersResp{}

	params := filter.urlValues()

	err := c._GET(ctx, ordersURI, params, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateOrder does _POST https://api.fortnox.se/3/orders/
//
// req - to create
//
// Should you have EasyVat enabled, it is mandatory to provide an account in the request should you use a custom VAT rate.
//
// This endpoint can produce errors, some of which may only be relevant for EasyVat. Refer to the table below.
//
// Errors that can be raised by this endpoint:
//
// | Error Code |	HTTP Code  | Description	| Solution |
// |------------|--------------|----------------|----------|
// | 2004167	| 400	       | An account must be provided when using a custom VAT rate and EasyVat has been enabled.	| Supply each row which has a custom VAT rate with an account. |
func (c Client) CreateOrder(ctx context.Context, req *CreateOrderReq) (*CreateOrderResp, error) {
	resp := &CreateOrderResp{}

	err := c._POST(ctx, ordersURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetOrder does _GET https://api.fortnox.se/3/orders/{DocumentNumber}
//
// documentNumber - identifies the order
func (c Client) GetOrder(ctx context.Context, documentNumber string) (*GetOrderResp, error) {
	resp := &GetOrderResp{}

	uri := fmt.Sprintf("%s/%s", ordersURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateOrder does _PUT https://api.fortnox.se/3/orders/{DocumentNumber}
//
// documentNumber - identifies the order
//
// req - to update
//
// Note that there are two approaches for updating the rows on an order.
//
// If RowId is not specified on any row, the rows will be mapped and updated in the order in which they are set in the array. All rows that should remain on the order needs to be provided.
//
// If RowId is specified on one or more rows the following goes: Corresponding row with that id will be updated. The rows without RowId will be interpreted as new rows. If a row should not be updated but remain on the order then specify only RowId like { "RowId": 123 }, otherwise it will be removed. Note that new RowIds are generated for all rows every time an order is updated.
func (c Client) UpdateOrder(ctx context.Context, documentNumber string, req *UpdateOrderReq) (*UpdateOrderResp, error) {
	resp := &UpdateOrderResp{}

	uri := fmt.Sprintf("%s/%s", ordersURI, documentNumber)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// PrintOrder does _GET https://api.fortnox.se/3/orders/{DocumentNumber}/print
//
// documentNumber - identifies the order
func (c Client) PrintOrder(ctx context.Context, documentNumber string) (*[]byte, error) {
	resp := &[]byte{}

	uri := fmt.Sprintf("%s/%s/print", ordersURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SendOrderAsEmail does _GET https://api.fortnox.se/3/orders/{DocumentNumber}/email
//
// documentNumber - identifies the order
//
// You can use the properties in the EmailInformation to customize the e-mail message on each order.
func (c Client) SendOrderAsEmail(ctx context.Context, documentNumber string) (*SendOrderAsEmailResp, error) {
	resp := &SendOrderAsEmailResp{}

	uri := fmt.Sprintf("%s/%s/email", ordersURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// PreviewOrder does _GET https://api.fortnox.se/3/orders/{DocumentNumber}/preview
// documentNumber - identifies the order
//
// The difference between this and the print-endpoint is that property Sent is not set to TRUE.
func (c Client) PreviewOrder(ctx context.Context, documentNumber string) (*[]byte, error) {
	resp := &[]byte{}

	uri := fmt.Sprintf("%s/%s/preview", ordersURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateInvoiceOutOfGivenOrder does _PUT https://api.fortnox.se/3/orders/{DocumentNumber}/createinvoice
//
// documentNumber - identifies the order
func (c Client) CreateInvoiceOutOfGivenOrder(
	ctx context.Context,
	documentNumber string) (*CreateInvoiceOutOfGivenOrderResp, error) {

	resp := &CreateInvoiceOutOfGivenOrderResp{}

	uri := fmt.Sprintf("%s/%s/createinvoice", ordersURI, documentNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CancelGivenOrder does _PUT https://api.fortnox.se/3/orders/{DocumentNumber}/cancel
//
// documentNumber - identifies the order
func (c Client) CancelGivenOrder(ctx context.Context, documentNumber string) (*CancelGivenOrderResp, error) {
	resp := &CancelGivenOrderResp{}

	uri := fmt.Sprintf("%s/%s/cancel", ordersURI, documentNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SetGivenOrderAsSent does _PUT https://api.fortnox.se/3/orders/{DocumentNumber}/externalprint
//
// documentNumber - identifies the order
//
// Use this endpoint to set order as sent, without generating an order.
func (c Client) SetGivenOrderAsSent(ctx context.Context, documentNumber string) (*SetGivenOrderAsSentResp, error) {
	resp := &SetGivenOrderAsSentResp{}

	uri := fmt.Sprintf("%s/%s/externalprint", ordersURI, documentNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetAllOrdersFilter string

const (
	CancelledGetAllOrdersFilter         GetAllOrdersFilter = "cancelled"
	ExpiredGetAllOrdersFilter           GetAllOrdersFilter = "expired"
	InvoiceCreatedGetAllOrdersFilter    GetAllOrdersFilter = "invoicecreated"
	invoiceNotCreatedGetAllOrdersFilter GetAllOrdersFilter = "invoicenotcreated"
)

func (f *GetAllOrdersFilter) urlValues() url.Values {
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

type GetAllOrdersResp struct {
	Orders []struct {
		Url                       string `json:"@url"`
		Cancelled                 bool   `json:"Cancelled"`
		Currency                  string `json:"Currency"`
		CustomerName              string `json:"CustomerName"`
		CustomerNumber            string `json:"CustomerNumber"`
		DeliveryDate              string `json:"DeliveryDate"`
		DocumentNumber            string `json:"DocumentNumber"`
		ExternalInvoiceReference1 string `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string `json:"ExternalInvoiceReference2"`
		OrderDate                 string `json:"OrderDate"`
		OrderType                 string `json:"OrderType"`
		Project                   string `json:"Project"`
		Sent                      bool   `json:"Sent"`
		Total                     int    `json:"Total"`
	} `json:"Orders"`
}

type CreateOrderReq struct {
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

type CreateOrderResp struct {
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

type GetOrderResp struct {
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

type UpdateOrderReq struct {
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

type UpdateOrderResp struct {
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

type SendOrderAsEmailResp struct {
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

type CreateInvoiceOutOfGivenOrderResp struct {
	Invoice struct {
		Url                    string `json:"@url"`
		UrlTaxReductionList    string `json:"@urlTaxReductionList"`
		AdministrationFee      int    `json:"AdministrationFee"`
		AdministrationFeeVAT   int    `json:"AdministrationFeeVAT"`
		Address1               string `json:"Address1"`
		Address2               string `json:"Address2"`
		Balance                int    `json:"Balance"`
		BasisTaxReduction      int    `json:"BasisTaxReduction"`
		Booked                 bool   `json:"Booked"`
		Cancelled              bool   `json:"Cancelled"`
		City                   string `json:"City"`
		Comments               string `json:"Comments"`
		ContractReference      int    `json:"ContractReference"`
		ContributionPercent    int    `json:"ContributionPercent"`
		ContributionValue      int    `json:"ContributionValue"`
		Country                string `json:"Country"`
		CostCenter             string `json:"CostCenter"`
		Credit                 string `json:"Credit"`
		CreditInvoiceReference string `json:"CreditInvoiceReference"`
		Currency               string `json:"Currency"`
		CurrencyRate           int    `json:"CurrencyRate"`
		CurrencyUnit           int    `json:"CurrencyUnit"`
		CustomerName           string `json:"CustomerName"`
		CustomerNumber         string `json:"CustomerNumber"`
		DeliveryAddress1       string `json:"DeliveryAddress1"`
		DeliveryAddress2       string `json:"DeliveryAddress2"`
		DeliveryCity           string `json:"DeliveryCity"`
		DeliveryCountry        string `json:"DeliveryCountry"`
		DeliveryDate           string `json:"DeliveryDate"`
		DeliveryName           string `json:"DeliveryName"`
		DeliveryZipCode        string `json:"DeliveryZipCode"`
		DocumentNumber         string `json:"DocumentNumber"`
		DueDate                string `json:"DueDate"`
		EDIInformation         struct {
			EDIGlobalLocationNumber         string `json:"EDIGlobalLocationNumber"`
			EDIGlobalLocationNumberDelivery string `json:"EDIGlobalLocationNumberDelivery"`
			EDIInvoiceExtra1                string `json:"EDIInvoiceExtra1"`
			EDIInvoiceExtra2                string `json:"EDIInvoiceExtra2"`
			EDIOurElectronicReference       string `json:"EDIOurElectronicReference"`
			EDIYourElectronicReference      string `json:"EDIYourElectronicReference"`
			EDIStatus                       string `json:"EDIStatus"`
		} `json:"EDIInformation"`
		EmailInformation struct {
			EmailAddressFrom string `json:"EmailAddressFrom"`
			EmailAddressTo   string `json:"EmailAddressTo"`
			EmailAddressCC   string `json:"EmailAddressCC"`
			EmailAddressBCC  string `json:"EmailAddressBCC"`
			EmailSubject     string `json:"EmailSubject"`
			EmailBody        string `json:"EmailBody"`
		} `json:"EmailInformation"`
		EUQuarterlyReport         bool   `json:"EUQuarterlyReport"`
		ExternalInvoiceReference1 string `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string `json:"ExternalInvoiceReference2"`
		Freight                   int    `json:"Freight"`
		FreightVAT                int    `json:"FreightVAT"`
		Gross                     int    `json:"Gross"`
		HouseWork                 bool   `json:"HouseWork"`
		InvoiceDate               string `json:"InvoiceDate"`
		InvoicePeriodStart        string `json:"InvoicePeriodStart"`
		InvoicePeriodEnd          string `json:"InvoicePeriodEnd"`
		InvoicePeriodReference    string `json:"InvoicePeriodReference"`
		InvoiceRows               []struct {
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
			Price                  int    `json:"Price"`
			PriceExcludingVAT      int    `json:"PriceExcludingVAT"`
			Project                string `json:"Project"`
			RowId                  int    `json:"RowId"`
			StockPointCode         string `json:"StockPointCode"`
			Total                  int    `json:"Total"`
			TotalExcludingVAT      int    `json:"TotalExcludingVAT"`
			Unit                   string `json:"Unit"`
			VAT                    int    `json:"VAT"`
		} `json:"InvoiceRows"`
		InvoiceType string `json:"InvoiceType"`
		Labels      []struct {
			Id int `json:"Id"`
		} `json:"Labels"`
		Language           string `json:"Language"`
		LastRemindDate     string `json:"LastRemindDate"`
		Net                int    `json:"Net"`
		NotCompleted       bool   `json:"NotCompleted"`
		NoxFinans          bool   `json:"NoxFinans"`
		OCR                string `json:"OCR"`
		OfferReference     string `json:"OfferReference"`
		OrderReference     string `json:"OrderReference"`
		OrganisationNumber string `json:"OrganisationNumber"`
		OurReference       string `json:"OurReference"`
		PaymentWay         string `json:"PaymentWay"`
		Phone1             string `json:"Phone1"`
		Phone2             string `json:"Phone2"`
		PriceList          string `json:"PriceList"`
		PrintTemplate      string `json:"PrintTemplate"`
		Project            string `json:"Project"`
		WarehouseReady     bool   `json:"WarehouseReady"`
		OutboundDate       string `json:"OutboundDate"`
		Remarks            string `json:"Remarks"`
		Reminders          int    `json:"Reminders"`
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
		VoucherNumber      int    `json:"VoucherNumber"`
		VoucherSeries      string `json:"VoucherSeries"`
		VoucherYear        int    `json:"VoucherYear"`
		WayOfDelivery      string `json:"WayOfDelivery"`
		YourOrderNumber    string `json:"YourOrderNumber"`
		YourReference      string `json:"YourReference"`
		ZipCode            string `json:"ZipCode"`
		AccountingMethod   string `json:"AccountingMethod"`
		TaxReductionType   string `json:"TaxReductionType"`
		FinalPayDate       string `json:"FinalPayDate"`
	} `json:"Invoice"`
}

type CancelGivenOrderResp struct {
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

type SetGivenOrderAsSentResp struct {
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

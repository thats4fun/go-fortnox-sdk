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
func (c Client) GetAllOrders(ctx context.Context, filter *GetAllOrdersFilter) ([]Order, error) {
	resp := &GetAllOrdersResp{}

	params := filter.urlValues()

	err := c._GET(ctx, ordersURI, params, resp)
	if err != nil {
		return nil, err
	}

	return resp.Orders, nil
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
func (c Client) CreateOrder(ctx context.Context, o *Order) (*Order, error) {
	req := &CreateOrderReq{Order: *o}
	resp := &CreateOrderResp{}

	err := c._POST(ctx, ordersURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Order, nil
}

// GetOrder does _GET https://api.fortnox.se/3/orders/{DocumentNumber}
//
// documentNumber - identifies the order
func (c Client) GetOrder(ctx context.Context, documentNumber string) (*Order, error) {
	resp := &GetOrderResp{}

	uri := fmt.Sprintf("%s/%s", ordersURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Order, nil
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
func (c Client) UpdateOrder(ctx context.Context, documentNumber string, o *Order) (*Order, error) {
	req := &UpdateOrderReq{Order: *o}
	resp := &UpdateOrderResp{}

	uri := fmt.Sprintf("%s/%s", ordersURI, documentNumber)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Order, nil
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
func (c Client) SendOrderAsEmail(ctx context.Context, documentNumber string) (*Order, error) {
	resp := &SendOrderAsEmailResp{}

	uri := fmt.Sprintf("%s/%s/email", ordersURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Order, nil
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
func (c Client) CreateInvoiceOutOfGivenOrder(ctx context.Context, documentNumber string) (*Invoice, error) {
	resp := &CreateInvoiceOutOfGivenOrderResp{}

	uri := fmt.Sprintf("%s/%s/createinvoice", ordersURI, documentNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Invoice, nil
}

// CancelGivenOrder does _PUT https://api.fortnox.se/3/orders/{DocumentNumber}/cancel
//
// documentNumber - identifies the order
func (c Client) CancelGivenOrder(ctx context.Context, documentNumber string) (*Order, error) {
	resp := &CancelGivenOrderResp{}

	uri := fmt.Sprintf("%s/%s/cancel", ordersURI, documentNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Order, nil
}

// SetGivenOrderAsSent does _PUT https://api.fortnox.se/3/orders/{DocumentNumber}/externalprint
//
// documentNumber - identifies the order
//
// Use this endpoint to set order as sent, without generating an order.
func (c Client) SetGivenOrderAsSent(ctx context.Context, documentNumber string) (*Order, error) {
	resp := &SetGivenOrderAsSentResp{}

	uri := fmt.Sprintf("%s/%s/externalprint", ordersURI, documentNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Order, nil
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
	Orders []Order `json:"Orders"`
}

type CreateOrderReq struct {
	Order Order `json:"Order"`
}

type CreateOrderResp struct {
	Order Order `json:"Order"`
}

type GetOrderResp struct {
	Order Order `json:"Order"`
}

type UpdateOrderReq struct {
	Order Order `json:"Order"`
}

type UpdateOrderResp struct {
	Order Order `json:"Order"`
}

type SendOrderAsEmailResp struct {
	Order Order `json:"Order"`
}

type CreateInvoiceOutOfGivenOrderResp struct {
	Invoice Invoice `json:"Invoice"`
}

type CancelGivenOrderResp struct {
	Order Order `json:"Order"`
}

type SetGivenOrderAsSentResp struct {
	Order Order `json:"Order"`
}

type Order struct {
	Url                       string           `json:"@url,omitempty"`
	UrlTaxReductionList       string           `json:"@urlTaxReductionList,omitempty"`
	AdministrationFee         int              `json:"AdministrationFee,omitempty"`
	AdministrationFeeVAT      int              `json:"AdministrationFeeVAT,omitempty"`
	Address1                  string           `json:"Address1,omitempty"`
	Address2                  string           `json:"Address2,omitempty"`
	BasisTaxReduction         int              `json:"BasisTaxReduction,omitempty"`
	Cancelled                 bool             `json:"Cancelled,omitempty"`
	City                      string           `json:"City,omitempty"`
	Comments                  string           `json:"Comments,omitempty"`
	ContributionPercent       int              `json:"ContributionPercent,omitempty"`
	ContributionValue         int              `json:"ContributionValue,omitempty"`
	CopyRemarks               bool             `json:"CopyRemarks,omitempty"`
	Country                   string           `json:"Country,omitempty"`
	CostCenter                string           `json:"CostCenter,omitempty"`
	Currency                  string           `json:"Currency,omitempty"`
	CurrencyRate              int              `json:"CurrencyRate,omitempty"`
	CurrencyUnit              int              `json:"CurrencyUnit,omitempty"`
	CustomerName              string           `json:"CustomerName,omitempty"`
	CustomerNumber            string           `json:"CustomerNumber,omitempty"`
	DeliveryState             string           `json:"DeliveryState,omitempty"`
	DeliveryAddress1          string           `json:"DeliveryAddress1,omitempty"`
	DeliveryAddress2          string           `json:"DeliveryAddress2,omitempty"`
	DeliveryCity              string           `json:"DeliveryCity,omitempty"`
	DeliveryCountry           string           `json:"DeliveryCountry,omitempty"`
	DeliveryDate              string           `json:"DeliveryDate,omitempty"`
	DeliveryName              string           `json:"DeliveryName,omitempty"`
	DeliveryZipCode           string           `json:"DeliveryZipCode,omitempty"`
	DocumentNumber            string           `json:"DocumentNumber,omitempty"`
	EmailInformation          EmailInformation `json:"EmailInformation,omitempty"`
	ExternalInvoiceReference1 string           `json:"ExternalInvoiceReference1,omitempty"`
	ExternalInvoiceReference2 string           `json:"ExternalInvoiceReference2,omitempty"`
	Freight                   int              `json:"Freight,omitempty"`
	FreightVAT                int              `json:"FreightVAT,omitempty"`
	Gross                     int              `json:"Gross,omitempty"`
	HouseWork                 bool             `json:"HouseWork,omitempty"`
	InvoiceReference          string           `json:"InvoiceReference,omitempty"`
	Labels                    []Label          `json:"Labels,omitempty"`
	Language                  string           `json:"Language,omitempty"`
	Net                       int              `json:"Net,omitempty"`
	NotCompleted              bool             `json:"NotCompleted,omitempty"`
	OfferReference            string           `json:"OfferReference,omitempty"`
	OrderDate                 string           `json:"OrderDate,omitempty"`
	OrderRows                 []OrderRow       `json:"OrderRows,omitempty"`
	OrderType                 string           `json:"OrderType,omitempty"`
	OrganisationNumber        string           `json:"OrganisationNumber,omitempty"`
	OurReference              string           `json:"OurReference,omitempty"`
	Phone1                    string           `json:"Phone1,omitempty"`
	Phone2                    string           `json:"Phone2,omitempty"`
	PriceList                 string           `json:"PriceList,omitempty"`
	PrintTemplate             string           `json:"PrintTemplate,omitempty"`
	Project                   string           `json:"Project,omitempty"`
	WarehouseReady            bool             `json:"WarehouseReady,omitempty"`
	OutboundDate              string           `json:"OutboundDate,omitempty"`
	Remarks                   string           `json:"Remarks,omitempty"`
	RoundOff                  int              `json:"RoundOff,omitempty"`
	Sent                      bool             `json:"Sent,omitempty"`
	TaxReduction              int              `json:"TaxReduction,omitempty"`
	TermsOfDelivery           string           `json:"TermsOfDelivery,omitempty"`
	TermsOfPayment            string           `json:"TermsOfPayment,omitempty"`
	TimeBasisReference        int              `json:"TimeBasisReference,omitempty"`
	Total                     int              `json:"Total,omitempty"`
	TotalToPay                int              `json:"TotalToPay,omitempty"`
	TotalVAT                  int              `json:"TotalVAT,omitempty"`
	VATIncluded               bool             `json:"VATIncluded,omitempty"`
	WayOfDelivery             string           `json:"WayOfDelivery,omitempty"`
	YourReference             string           `json:"YourReference,omitempty"`
	YourOrderNumber           string           `json:"YourOrderNumber,omitempty"`
	ZipCode                   string           `json:"ZipCode,omitempty"`
	StockPointCode            string           `json:"StockPointCode,omitempty"`
	StockPointId              string           `json:"StockPointId,omitempty"`
	TaxReductionType          string           `json:"TaxReductionType,omitempty"`
}

type OrderRow struct {
	AccountNumber          int    `json:"AccountNumber,omitempty"`
	ArticleNumber          string `json:"ArticleNumber,omitempty"`
	ContributionPercent    string `json:"ContributionPercent,omitempty"`
	ContributionValue      string `json:"ContributionValue,omitempty"`
	CostCenter             string `json:"CostCenter,omitempty"`
	DeliveredQuantity      string `json:"DeliveredQuantity,omitempty"`
	Description            string `json:"Description,omitempty"`
	Discount               int    `json:"Discount,omitempty"`
	DiscountType           string `json:"DiscountType,omitempty"`
	HouseWork              bool   `json:"HouseWork,omitempty"`
	HouseWorkHoursToReport int    `json:"HouseWorkHoursToReport,omitempty"`
	HouseWorkType          string `json:"HouseWorkType,omitempty"`
	OrderedQuantity        string `json:"OrderedQuantity,omitempty"`
	Price                  int    `json:"Price,omitempty"`
	Project                string `json:"Project,omitempty"`
	ReservedQuantity       string `json:"ReservedQuantity,omitempty"`
	RowId                  int    `json:"RowId,omitempty"`
	StockPointCode         string `json:"StockPointCode,omitempty"`
	StockPointId           string `json:"StockPointId,omitempty"`
	Total                  int    `json:"Total,omitempty"`
	Unit                   string `json:"Unit,omitempty"`
	VAT                    int    `json:"VAT,omitempty"`
}

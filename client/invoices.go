package client

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

const (
	invoicesURI = "invoices"
)

// GetInvoice does _GET https://api.fortnox.se/3/invoices/{DocumentNumber}
//
// documentNumber - identifies the invoice
func (c *Client) GetInvoice(ctx context.Context, documentNumber string) (*Invoice, error) {
	resp := &GetInvoiceResp{}

	uri := fmt.Sprintf("%s/%s", invoicesURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Invoice, nil
}

// UpdateInvoice odes _PUT https://api.fortnox.se/3/invoices/{DocumentNumber}
//
// documentNumber - identifies the invoice
//
// req - payload
func (c *Client) UpdateInvoice(
	ctx context.Context,
	documentNumber string,
	i *Invoice) (*Invoice, error) {

	req := &UpdateInvoiceReq{*i}
	resp := &UpdateInvoiceResp{}

	uri := fmt.Sprintf("%s/%s", invoicesURI, documentNumber)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Invoice, nil
}

// GetAllInvoices does _GET https://api.fortnox.se/3/invoices
//
// queryParams - filters
func (c *Client) GetAllInvoices(ctx context.Context, queryParams *GetAllInvoicesQueryParams) ([]Invoice, error) {
	resp := &GetAllInvoicesResp{}

	var params url.Values
	if queryParams != nil {
		params = queryParams.urlValues()
	}

	err := c._GET(ctx, invoicesURI, params, resp)
	if err != nil {
		return nil, err
	}

	return resp.Invoices, nil
}

// CreateInvoice does _POST https://api.fortnox.se/3/invoices
//
// req - payload
func (c *Client) CreateInvoice(ctx context.Context, i *Invoice) (*Invoice, error) {
	req := &CreateInvoiceReq{Invoice: *i}
	resp := &CreateInvoiceResp{}

	err := c._POST(ctx, invoicesURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Invoice, nil
}

// BookKeepInvoice does _PUT https://api.fortnox.se/3/invoices/{DocumentNumber}/bookkeep
//
// documentNumber - identifies the invoice
func (c *Client) BookKeepInvoice(ctx context.Context, documentNumber string) (*Invoice, error) {
	resp := &BookKeepInvoiceResp{}

	uri := fmt.Sprintf("%s/%s/bookkeep", invoicesURI, documentNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Invoice, nil
}

// CancelInvoice does _PUT https://api.fortnox.se/3/invoices/{DocumentNumber}/cancel
//
// documentNumber - identifies the invoice
func (c *Client) CancelInvoice(ctx context.Context, documentNumber string) (*Invoice, error) {
	resp := &CancelInvoiceResp{}

	uri := fmt.Sprintf("%s/%s/cancel", invoicesURI, documentNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Invoice, nil
}

// CreditInvoice does _PUT https://api.fortnox.se/3/invoices/{DocumentNumber}/credit
//
// documentNumber - identifies the invoice
func (c *Client) CreditInvoice(ctx context.Context, documentNumber string) (*Invoice, error) {
	resp := &CreditInvoiceResp{}

	uri := fmt.Sprintf("%s/%s/credit", invoicesURI, documentNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Invoice, nil
}

// SetInvoiceAsSent does _PUT https://api.fortnox.se/3/invoices/{DocumentNumber}/externalprint
//
// documentNumber - identifies the invoice
func (c *Client) SetInvoiceAsSent(ctx context.Context, documentNumber string) (*Invoice, error) {
	resp := &SetInvoiceAsSentResp{}

	uri := fmt.Sprintf("%s/%s/externalprint", invoicesURI, documentNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Invoice, nil
}

// SetInvoiceAsDone does _PUT https://api.fortnox.se/3/invoices/{DocumentNumber}/warehouseready
//
// documentNumber - identifies the invoice
func (c *Client) SetInvoiceAsDone(ctx context.Context, documentNumber string) (*Invoice, error) {
	resp := &SetInvoiceAsDoneResp{}

	uri := fmt.Sprintf("%s/%s/warehouseready", invoicesURI, documentNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Invoice, nil
}

// PrintInvoice does _PUT https://api.fortnox.se/3/invoices/{DocumentNumber}/print
//
// documentNumber - identifies the invoice
func (c *Client) PrintInvoice(ctx context.Context, documentNumber string) (*[]byte, error) {
	resp := &[]byte{}

	uri := fmt.Sprintf("%s/%s/print", invoicesURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SendInvoiceAsEmail does _PUT https://api.fortnox.se/3/invoices/{DocumentNumber}/email
//
// documentNumber - identifies the invoice
func (c *Client) SendInvoiceAsEmail(ctx context.Context, documentNumber string) (*Invoice, error) {
	resp := &SendInvoiceAsEmailResp{}

	uri := fmt.Sprintf("%s/%s/email", invoicesURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Invoice, nil
}

// SendInvoiceAsReminder does _GET https://api.fortnox.se/3/invoices/{DocumentNumber}/printreminder
//
// documentNumber - identifies the invoice
func (c *Client) SendInvoiceAsReminder(ctx context.Context, documentNumber string) (*[]byte, error) {
	resp := &[]byte{}

	uri := fmt.Sprintf("%s/%s/printreminder", invoicesURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// PreviewInvoice does _GET https://api.fortnox.se/3/invoices/{DocumentNumber}/preview
//
// documentNumber - identifies the invoice
func (c *Client) PreviewInvoice(ctx context.Context, documentNumber string) (*[]byte, error) {
	resp := &[]byte{}

	uri := fmt.Sprintf("%s/%s/preview", invoicesURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SendInvoiceAsEPrint does _GET https://api.fortnox.se/3/invoices/{DocumentNumber}/eprint
//
// documentNumber - identifies the invoice
func (c *Client) SendInvoiceAsEPrint(ctx context.Context, documentNumber string) (*Invoice, error) {
	resp := &SendInvoiceAsEPrintResp{}

	uri := fmt.Sprintf("%s/%s/eprint", invoicesURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Invoice, nil
}

// SendInvoiceAsEInvoice does _GET https://api.fortnox.se/3/invoices/{DocumentNumber}/einvoice
//
// documentNumber - identifies the invoice
func (c *Client) SendInvoiceAsEInvoice(ctx context.Context, documentNumber string) (*Invoice, error) {
	resp := &SendInvoiceAsEInvoiceResp{}

	uri := fmt.Sprintf("%s/%s/einvoice", invoicesURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Invoice, nil
}

type GetAllInvoicesQueryParams struct {
	Filter                    GetAllInvoicesFilter
	CostCenter                string
	CustomerName              string
	CustomerNumber            string
	Label                     string
	DocumentNumber            string
	FromDate                  string
	ToDate                    string
	FromFinalPayDate          string
	ToFinalPayDate            string
	LastModified              string
	NotCompleted              string
	Ocr                       string
	OurReference              string
	Project                   string
	Sent                      string
	ExternalInvoiceReference1 string
	ExternalInvoiceReference2 string
	YourReference             string
	InvoiceType               string
	ArticleNumber             string
	ArticleDescription        string
	Currency                  string
	AccountNumberFrom         string
	AccountNumberTo           string
	YourOrderNumber           string
	Credit                    string
	SortBy                    GetAllInvoicesSortBy
}

func (p GetAllInvoicesQueryParams) urlValues() url.Values {
	params := url.Values{}

	filterStr := string(p.Filter)
	if strings.TrimSpace(filterStr) != "" {
		params["filter"] = []string{filterStr}
	}
	if strings.TrimSpace(p.CostCenter) != "" {
		params["costcenter"] = []string{p.CostCenter}
	}
	if strings.TrimSpace(p.CustomerName) != "" {
		params["customername"] = []string{p.CostCenter}
	}
	if strings.TrimSpace(p.CustomerNumber) != "" {
		params["customername"] = []string{p.CustomerNumber}
	}
	if strings.TrimSpace(p.Label) != "" {
		params["customername"] = []string{p.Label}
	}
	if strings.TrimSpace(p.DocumentNumber) != "" {
		params["customername"] = []string{p.DocumentNumber}
	}
	if strings.TrimSpace(p.FromDate) != "" {
		params["customername"] = []string{p.FromDate}
	}
	return params
}

type GetAllInvoicesSortBy string

const (
	CustomerName   = "customername"
	CustomerNumber = "customernumber"
	DocumentNumber = "documentnumber"
	InvoiceDate    = "invoicedate"
	Ocr            = "ocr"
	Total          = "total"
)

type GetAllInvoicesFilter string

const (
	Cancelled     GetAllInvoicesFilter = "cancelled"
	FullyPaid     GetAllInvoicesFilter = "fullypaid"
	Unpaid        GetAllInvoicesFilter = "unpaid"
	UnpaidOverDue GetAllInvoicesFilter = "unpaidoverdue"
	Unbooked      GetAllInvoicesFilter = "unbooked"
)

type GetInvoiceResp struct {
	Invoice Invoice `json:"Invoice"`
}

type UpdateInvoiceReq struct {
	Invoice Invoice `json:"Invoice"`
}

type UpdateInvoiceResp struct {
	Invoice Invoice `json:"Invoice"`
}

type GetAllInvoicesResp struct {
	Invoices []Invoice `json:"Invoices"`
}

type CreateInvoiceReq struct {
	Invoice Invoice `json:"Invoice"`
}

type CreateInvoiceResp struct {
	Invoice Invoice `json:"Invoice"`
}

type BookKeepInvoiceResp struct {
	Invoice Invoice `json:"Invoice"`
}

type CancelInvoiceResp struct {
	Invoice Invoice `json:"Invoice"`
}

type CreditInvoiceResp struct {
	Invoice Invoice `json:"Invoice"`
}

type SetInvoiceAsSentResp struct {
	Invoice Invoice `json:"Invoice"`
}

type SetInvoiceAsDoneResp struct {
	Invoice Invoice `json:"Invoice"`
}

type SendInvoiceAsEmailResp struct {
	Invoice Invoice `json:"Invoice"`
}

type SendInvoiceAsEPrintResp struct {
	Invoice Invoice `json:"Invoice"`
}

type SendInvoiceAsEInvoiceResp struct {
	Invoice Invoice `json:"Invoice"`
}

type Invoice struct {
	Url                       string           `json:"@url,omitempty"`
	UrlTaxReductionList       string           `json:"@urlTaxReductionList,omitempty"`
	AdministrationFee         int              `json:"AdministrationFee,omitempty"`
	AdministrationFeeVAT      int              `json:"AdministrationFeeVAT,omitempty"`
	Address1                  string           `json:"Address1,omitempty"`
	Address2                  string           `json:"Address2,omitempty"`
	Balance                   int              `json:"Balance,omitempty"`
	BasisTaxReduction         int              `json:"BasisTaxReduction,omitempty"`
	Booked                    bool             `json:"Booked,omitempty"`
	Cancelled                 bool             `json:"Cancelled,omitempty"`
	City                      string           `json:"City,omitempty"`
	Comments                  string           `json:"Comments,omitempty"`
	ContractReference         int              `json:"ContractReference,omitempty"`
	ContributionPercent       int              `json:"ContributionPercent,omitempty"`
	ContributionValue         int              `json:"ContributionValue,omitempty"`
	Country                   string           `json:"Country,omitempty"`
	CostCenter                string           `json:"CostCenter,omitempty"`
	Credit                    string           `json:"Credit,omitempty"`
	CreditInvoiceReference    string           `json:"CreditInvoiceReference,omitempty"`
	Currency                  string           `json:"Currency,omitempty"`
	CurrencyRate              int              `json:"CurrencyRate,omitempty"`
	CurrencyUnit              int              `json:"CurrencyUnit,omitempty"`
	CustomerName              string           `json:"CustomerName,omitempty"`
	CustomerNumber            string           `json:"CustomerNumber,omitempty"`
	DeliveryAddress1          string           `json:"DeliveryAddress1,omitempty"`
	DeliveryAddress2          string           `json:"DeliveryAddress2,omitempty"`
	DeliveryCity              string           `json:"DeliveryCity,omitempty"`
	DeliveryCountry           string           `json:"DeliveryCountry,omitempty"`
	DeliveryDate              string           `json:"DeliveryDate,omitempty"`
	DeliveryName              string           `json:"DeliveryName,omitempty"`
	DeliveryZipCode           string           `json:"DeliveryZipCode,omitempty"`
	DocumentNumber            string           `json:"DocumentNumber,omitempty"`
	DueDate                   string           `json:"DueDate,omitempty"`
	EDIInformation            EDIInformation   `json:"EDIInformation,omitempty"`
	EmailInformation          EmailInformation `json:"EmailInformation,omitempty"`
	EUQuarterlyReport         bool             `json:"EUQuarterlyReport,omitempty"`
	ExternalInvoiceReference1 string           `json:"ExternalInvoiceReference1,omitempty"`
	ExternalInvoiceReference2 string           `json:"ExternalInvoiceReference2,omitempty"`
	Freight                   int              `json:"Freight,omitempty"`
	FreightVAT                int              `json:"FreightVAT,omitempty"`
	Gross                     int              `json:"Gross,omitempty"`
	HouseWork                 bool             `json:"HouseWork,omitempty"`
	InvoiceDate               string           `json:"InvoiceDate,omitempty"`
	InvoicePeriodStart        string           `json:"InvoicePeriodStart,omitempty"`
	InvoicePeriodEnd          string           `json:"InvoicePeriodEnd,omitempty"`
	InvoicePeriodReference    string           `json:"InvoicePeriodReference,omitempty"`
	InvoiceRows               []InvoiceRow     `json:"InvoiceRows,omitempty"`
	InvoiceType               string           `json:"InvoiceType,omitempty"`
	Labels                    []Label          `json:"Labels,omitempty"`
	Language                  string           `json:"Language,omitempty"`
	LastRemindDate            string           `json:"LastRemindDate,omitempty"`
	Net                       int              `json:"Net,omitempty"`
	NotCompleted              bool             `json:"NotCompleted,omitempty"`
	NoxFinans                 bool             `json:"NoxFinans,omitempty"`
	OCR                       string           `json:"OCR,omitempty"`
	OfferReference            string           `json:"OfferReference,omitempty"`
	OrderReference            string           `json:"OrderReference,omitempty"`
	OrganisationNumber        string           `json:"OrganisationNumber,omitempty"`
	OurReference              string           `json:"OurReference,omitempty"`
	PaymentWay                string           `json:"PaymentWay,omitempty"`
	Phone1                    string           `json:"Phone1,omitempty"`
	Phone2                    string           `json:"Phone2,omitempty"`
	PriceList                 string           `json:"PriceList,omitempty"`
	PrintTemplate             string           `json:"PrintTemplate,omitempty"`
	Project                   string           `json:"Project,omitempty"`
	WarehouseReady            bool             `json:"WarehouseReady,omitempty"`
	OutboundDate              string           `json:"OutboundDate,omitempty"`
	Remarks                   string           `json:"Remarks,omitempty"`
	Reminders                 int              `json:"Reminders,omitempty"`
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
	VoucherNumber             int              `json:"VoucherNumber,omitempty"`
	VoucherSeries             string           `json:"VoucherSeries,omitempty"`
	VoucherYear               int              `json:"VoucherYear,omitempty"`
	WayOfDelivery             string           `json:"WayOfDelivery,omitempty"`
	YourOrderNumber           string           `json:"YourOrderNumber,omitempty"`
	YourReference             string           `json:"YourReference,omitempty"`
	ZipCode                   string           `json:"ZipCode,omitempty"`
	AccountingMethod          string           `json:"AccountingMethod,omitempty"`
	TaxReductionType          string           `json:"TaxReductionType,omitempty"`
	FinalPayDate              string           `json:"FinalPayDate,omitempty"`
}

type EDIInformation struct {
	EDIGlobalLocationNumber         string `json:"EDIGlobalLocationNumber,omitempty"`
	EDIGlobalLocationNumberDelivery string `json:"EDIGlobalLocationNumberDelivery,omitempty"`
	EDIInvoiceExtra1                string `json:"EDIInvoiceExtra1,omitempty"`
	EDIInvoiceExtra2                string `json:"EDIInvoiceExtra2,omitempty"`
	EDIOurElectronicReference       string `json:"EDIOurElectronicReference,omitempty"`
	EDIYourElectronicReference      string `json:"EDIYourElectronicReference,omitempty"`
	EDIStatus                       string `json:"EDIStatus,omitempty"`
}

type EmailInformation struct {
	EmailAddressFrom string `json:"EmailAddressFrom,omitempty"`
	EmailAddressTo   string `json:"EmailAddressTo,omitempty"`
	EmailAddressCC   string `json:"EmailAddressCC,omitempty"`
	EmailAddressBCC  string `json:"EmailAddressBCC,omitempty"`
	EmailSubject     string `json:"EmailSubject,omitempty"`
	EmailBody        string `json:"EmailBody,omitempty"`
}

type InvoiceRow struct {
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
	Price                  int    `json:"Price,omitempty"`
	PriceExcludingVAT      int    `json:"PriceExcludingVAT,omitempty"`
	Project                string `json:"Project,omitempty"`
	RowId                  int    `json:"RowId,omitempty"`
	StockPointCode         string `json:"StockPointCode,omitempty"`
	Total                  int    `json:"Total,omitempty"`
	TotalExcludingVAT      int    `json:"TotalExcludingVAT,omitempty"`
	Unit                   string `json:"Unit,omitempty"`
	VAT                    int    `json:"VAT,omitempty"`
}

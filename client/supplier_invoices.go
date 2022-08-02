package client

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

const (
	supplierInvoiceURI = "supplierinvoices"
)

const (
	getAllSupplierInvoicesFilterParamName = "filter"
)

// GetAllSupplierInvoices does _GET https://api.fortnox.se/3/supplierinvoices/
//
// filter - Enum: "cancelled" "fullypaid" "unpaid" "unpaidoverdue" "unbooked" "pendingpayment" "authorizepending"
// possibility to filter supplier invoices
func (c *Client) GetAllSupplierInvoices(ctx context.Context, filter *GetAllSupplierInvoicesFilter) ([]SupplierInvoice, error) {
	resp := &GetAllSupplierInvoicesResp{}

	params := filter.urlValues()

	err := c._GET(ctx, supplierInvoiceURI, params, resp)
	if err != nil {
		return nil, err
	}

	return resp.SupplierInvoices, nil
}

// CreateSupplierInvoice does _POST https://api.fortnox.se/3/supplierinvoices/
//
// req - supplier invoice to create
func (c *Client) CreateSupplierInvoice(ctx context.Context, si *SupplierInvoice) (*SupplierInvoice, error) {
	req := &CreateSupplierInvoiceReq{SupplierInvoice: *si}
	resp := &CreateSupplierInvoiceResp{}

	err := c._POST(ctx, supplierInvoiceURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.SupplierInvoice, nil
}

// GetSupplierInvoice does _GET https://api.fortnox.se/3/supplierinvoices/{GivenNumber}
//
// givenNumber - identifies the invoice
func (c *Client) GetSupplierInvoice(ctx context.Context, givenNumber int) (*SupplierInvoice, error) {
	resp := &GetSupplierInvoiceResp{}

	uri := fmt.Sprintf("%s/%d", supplierInvoiceURI, givenNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.SupplierInvoice, nil
}

// UpdateSupplierInvoice does _PUT https://api.fortnox.se/3/supplierinvoices/{GivenNumber}
//
// givenNumber - identifies the invoice
//
// si - supplier invoice to update
func (c *Client) UpdateSupplierInvoice(ctx context.Context, givenNumber int, si *SupplierInvoice) (*SupplierInvoice, error) {
	req := &UpdateSupplierInvoiceReq{SupplierInvoice: *si}
	resp := &UpdateSupplierInvoiceResp{}

	uri := fmt.Sprintf("%s/%d", supplierInvoiceURI, givenNumber)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.SupplierInvoice, nil
}

// BookKeepSupplierInvoice does _PUT https://api.fortnox.se/3/supplierinvoices/{GivenNumber}/bookkeep
//
// givenNumber - identifies the invoice
func (c *Client) BookKeepSupplierInvoice(ctx context.Context, givenNumber int) (*SupplierInvoice, error) {
	resp := &BookKeepSupplierInvoiceResp{}

	uri := fmt.Sprintf("%s/%d/bookkeep", supplierInvoiceURI, givenNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.SupplierInvoice, nil
}

// CancelSupplierInvoice does _PUT https://api.fortnox.se/3/supplierinvoices/{GivenNumber}/cancel
//
// givenNumber - identifies the invoice
func (c *Client) CancelSupplierInvoice(ctx context.Context, givenNumber int) (*SupplierInvoice, error) {
	resp := &CancelSupplierInvoiceResp{}

	uri := fmt.Sprintf("%s/%d/cancel", supplierInvoiceURI, givenNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.SupplierInvoice, nil
}

// CreditSupplierInvoicePayment does _PUT https://api.fortnox.se/3/supplierinvoices/{GivenNumber}/credit
//
// givenNumber - identifies the invoice
func (c *Client) CreditSupplierInvoicePayment(ctx context.Context, givenNumber int) (*SupplierInvoice, error) {
	resp := &CreditSupplierInvoicePaymentResp{}

	uri := fmt.Sprintf("%s/%d/credit", supplierInvoiceURI, givenNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.SupplierInvoice, nil
}

// ApprovalSupplierInvoicePayment does _PUT https://api.fortnox.se/3/supplierinvoices/{GivenNumber}/approvalpayment
func (c *Client) ApprovalSupplierInvoicePayment(ctx context.Context, givenNumber int) (*SupplierInvoice, error) {
	resp := &ApprovalSupplierInvoicePaymentResp{}

	uri := fmt.Sprintf("%s/%d/approvalpayment", supplierInvoiceURI, givenNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.SupplierInvoice, nil
}

// ApprovalSupplierInvoiceBookKeep does _PUT https://api.fortnox.se/3/supplierinvoices/{GivenNumber}/approvalbookkeep
//
// givenNumber - identifies the invoice
func (c *Client) ApprovalSupplierInvoiceBookKeep(ctx context.Context, givenNumber int) (*SupplierInvoice, error) {

	resp := &ApprovalSupplierInvoiceBookKeepResp{}

	uri := fmt.Sprintf("%s/%d/approvalbookkeep", supplierInvoiceURI, givenNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.SupplierInvoice, nil
}

type GetAllSupplierInvoicesFilter string

const (
	CancelledSupplierInvoiceFilter        GetAllSupplierInvoicesFilter = "cancelled"
	fullyPaidSupplierInvoiceFilter        GetAllSupplierInvoicesFilter = "fullypaid"
	UnpaidSupplierInvoiceFilter           GetAllSupplierInvoicesFilter = "unpaid"
	UnPaidOverdueSupplierInvoiceFilter    GetAllSupplierInvoicesFilter = "unpaidoverdue"
	UnbookedSupplierInvoiceFilter         GetAllSupplierInvoicesFilter = "unbooked"
	PendingPaymentSupplierInvoiceFilter   GetAllSupplierInvoicesFilter = "pendingpayment"
	AuthorizePendingSupplierInvoiceFilter GetAllSupplierInvoicesFilter = "authorizepending"
)

func (f *GetAllSupplierInvoicesFilter) urlValues() url.Values {
	if f == nil {
		return nil
	}

	params := url.Values{}

	fStr := string(*f)
	if strings.TrimSpace(fStr) != "" {
		params[getAllSupplierInvoicesFilterParamName] = []string{fStr}
	}

	return params
}

type GetAllSupplierInvoicesResp struct {
	SupplierInvoices []SupplierInvoice `json:"SupplierInvoices"`
}

type CreateSupplierInvoiceReq struct {
	SupplierInvoice SupplierInvoice `json:"SupplierInvoice"`
}

type CreateSupplierInvoiceResp struct {
	SupplierInvoice SupplierInvoice `json:"SupplierInvoice"`
}

type GetSupplierInvoiceResp struct {
	SupplierInvoice SupplierInvoice `json:"SupplierInvoice"`
}

type UpdateSupplierInvoiceReq struct {
	SupplierInvoice SupplierInvoice `json:"SupplierInvoice"`
}

type UpdateSupplierInvoiceResp struct {
	SupplierInvoice SupplierInvoice `json:"SupplierInvoice"`
}

type BookKeepSupplierInvoiceResp struct {
	SupplierInvoice SupplierInvoice `json:"SupplierInvoice"`
}

type CancelSupplierInvoiceResp struct {
	SupplierInvoice SupplierInvoice `json:"SupplierInvoice"`
}

type CreditSupplierInvoicePaymentResp struct {
	SupplierInvoice SupplierInvoice `json:"SupplierInvoice"`
}

type ApprovalSupplierInvoicePaymentResp struct {
	SupplierInvoice SupplierInvoice `json:"SupplierInvoice"`
}

type ApprovalSupplierInvoiceBookKeepResp struct {
	SupplierInvoice SupplierInvoice `json:"SupplierInvoice"`
}

type SupplierInvoiceRow struct {
	Account                int    `json:"Account,omitempty"`
	ArticleNumber          string `json:"ArticleNumber,omitempty"`
	Code                   string `json:"Code,omitempty"`
	CostCenter             string `json:"CostCenter,omitempty"`
	AccountDescription     string `json:"AccountDescription,omitempty"`
	ItemDescription        string `json:"ItemDescription,omitempty"`
	Debit                  int    `json:"Debit,omitempty"`
	DebitCurrency          int    `json:"DebitCurrency,omitempty"`
	Credit                 int    `json:"Credit,omitempty"`
	CreditCurrency         int    `json:"CreditCurrency,omitempty"`
	Project                string `json:"Project,omitempty"`
	TransactionInformation string `json:"TransactionInformation,omitempty"`
	Price                  int    `json:"Price,omitempty"`
	Quantity               int    `json:"Quantity,omitempty"`
	Total                  int    `json:"Total,omitempty"`
	Unit                   string `json:"Unit,omitempty"`
	StockPointCode         string `json:"StockPointCode,omitempty"`
	StockLocationCode      string `json:"StockLocationCode,omitempty"`
}

type SupplierInvoice struct {
	Url                   string               `json:"@url,omitempty"`
	AdministrationFee     string               `json:"AdministrationFee,omitempty"`
	Balance               string               `json:"Balance,omitempty"`
	Booked                bool                 `json:"Booked,omitempty"`
	Cancelled             bool                 `json:"Cancelled,omitempty"`
	Comments              string               `json:"Comments,omitempty"`
	CostCenter            string               `json:"CostCenter,omitempty"`
	Credit                bool                 `json:"Credit,omitempty"`
	CreditReference       int                  `json:"CreditReference,omitempty"`
	Currency              string               `json:"Currency,omitempty"`
	CurrencyRate          string               `json:"CurrencyRate,omitempty"`
	CurrencyUnit          int                  `json:"CurrencyUnit,omitempty"`
	DisablePaymentFile    bool                 `json:"DisablePaymentFile,omitempty"`
	DueDate               string               `json:"DueDate,omitempty"`
	ExternalInvoiceNumber string               `json:"ExternalInvoiceNumber,omitempty"`
	ExternalInvoiceSeries string               `json:"ExternalInvoiceSeries,omitempty"`
	Freight               string               `json:"Freight,omitempty"`
	GivenNumber           string               `json:"GivenNumber,omitempty"`
	InvoiceDate           string               `json:"InvoiceDate,omitempty"`
	InvoiceNumber         string               `json:"InvoiceNumber,omitempty"`
	OCR                   string               `json:"OCR,omitempty"`
	OurReference          string               `json:"OurReference,omitempty"`
	PaymentPending        bool                 `json:"PaymentPending,omitempty"`
	Project               string               `json:"Project,omitempty"`
	RoundOffValue         string               `json:"RoundOffValue,omitempty"`
	SupplierInvoiceRows   []SupplierInvoiceRow `json:"SupplierInvoiceRows,omitempty"`
	SupplierNumber        string               `json:"SupplierNumber,omitempty"`
	SupplierName          string               `json:"SupplierName,omitempty"`
	Total                 string               `json:"Total,omitempty"`
	VAT                   string               `json:"VAT,omitempty"`
	YourReference         string               `json:"YourReference,omitempty"`
	VoucherNumber         int                  `json:"VoucherNumber,omitempty"`
	VoucherSeries         string               `json:"VoucherSeries,omitempty"`
	VoucherYear           int                  `json:"VoucherYear,omitempty"`
	VATType               string               `json:"VATType,omitempty"`
	SalesType             string               `json:"SalesType,omitempty"`
	AccountingMethod      string               `json:"AccountingMethod,omitempty"`
	Vouchers              []Voucher            `json:"Vouchers,omitempty"`
	FinalPayDate          string               `json:"FinalPayDate,omitempty"`
}

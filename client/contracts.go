package client

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

const (
	contractsURI = "contracts"
)

// GetContract does _GET https://api.fortnox.se/3/contracts/{DocumentNumber}
//
// documentNumber - identifies the contract
func (c *Client) GetContract(ctx context.Context, documentNumber string) (*GetContractResp, error) {
	resp := &GetContractResp{}

	uri := fmt.Sprintf("%s/%s", contractsURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateContract does _PUT https://api.fortnox.se/3/contracts/{DocumentNumber}
//
// documentNumber - identifies the contract
func (c *Client) UpdateContract(
	ctx context.Context,
	documentNumber int,
	req *UpdateContractReq) (*UpdateContractResp, error) {

	resp := &UpdateContractResp{}

	uri := fmt.Sprintf("%s/%d", contractsURI, documentNumber)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetAllContract does _GET https://api.fortnox.se/3/contracts/
//
// filter - Enum: "active" "inactive" "finished"  possibility to filter contracts
func (c *Client) GetAllContract(ctx context.Context, filter GetAllContractFilter) (*GetContractResp, error) {
	resp := &GetContractResp{}

	params := filter.urlValues()

	err := c._GET(ctx, contractsURI, params, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateContract does _POST https://api.fortnox.se/3/contracts/
//
// req - contract to create
func (c *Client) CreateContract(ctx context.Context, req *CreateContractReq) (*CreateContractResp, error) {
	resp := &CreateContractResp{}

	err := c._POST(ctx, contractsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SetContractAsFinished does _PUT https://api.fortnox.se/3/contracts/{DocumentNumber}/finish
//
// documentNumber - identifies the contract
func (c *Client) SetContractAsFinished(ctx context.Context, documentNumber string) (*SetContractAsFinishedResp, error) {
	resp := &SetContractAsFinishedResp{}

	uri := fmt.Sprintf("%s/%s/finish", contractsURI, documentNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateInvoiceFromContract does _PUT https://api.fortnox.se/3/contracts/{DocumentNumber}/createinvoice
//
// identifies the contract
func (c *Client) CreateInvoiceFromContract(ctx context.Context, documentNumber string) (*CreateInvoiceFromContractResp, error) {
	resp := &CreateInvoiceFromContractResp{}

	uri := fmt.Sprintf("%s/%s/createinvoice", contractsURI, documentNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// IncreaseInvoiceCount does _PUT https://api.fortnox.se/3/contracts/{DocumentNumber}/increaseinvoicecount
//
// documentNumber - identifies the contract
func (c *Client) IncreaseInvoiceCount(ctx context.Context, documentNumber string) (*CreateInvoiceFromContractResp, error) {
	resp := &CreateInvoiceFromContractResp{}

	uri := fmt.Sprintf("%s/%s/increaseinvoicecount", contractsURI, documentNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetAllContractFilter string

const (
	ActiveGetAllContractFilter   GetAllContractFilter = "active"
	InactiveGetAllContractFilter GetAllContractFilter = "inactive"
	FinishedGetAllContractFilter GetAllContractFilter = "finished"
)

func (f GetAllContractFilter) urlValues() url.Values {
	urlValues := url.Values{}

	fStr := string(f)

	if strings.TrimSpace(fStr) != "" {
		urlValues["filter"] = []string{fStr}
	}

	return urlValues
}

type GetContractResp struct {
	Contract struct {
		Url                       string           `json:"@url"`
		UrlTaxReductionList       string           `json:"@urlTaxReductionList"`
		Active                    bool             `json:"Active"`
		AdministrationFee         int              `json:"AdministrationFee"`
		BasisTaxReduction         int              `json:"BasisTaxReduction"`
		Comments                  string           `json:"Comments"`
		Continuous                bool             `json:"Continuous"`
		ContractDate              string           `json:"ContractDate"`
		ContractLength            int              `json:"ContractLength"`
		ContributionPercent       int              `json:"ContributionPercent"`
		ContributionValue         int              `json:"ContributionValue"`
		CostCenter                string           `json:"CostCenter"`
		Currency                  string           `json:"Currency"`
		CustomerName              string           `json:"CustomerName"`
		CustomerNumber            string           `json:"CustomerNumber"`
		DocumentNumber            string           `json:"DocumentNumber"`
		EmailInformation          EmailInformation `json:"EmailInformation"`
		ExternalInvoiceReference1 string           `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string           `json:"ExternalInvoiceReference2"`
		Freight                   int              `json:"Freight"`
		Gross                     int              `json:"Gross"`
		HouseWork                 bool             `json:"HouseWork"`
		InvoiceDiscount           int              `json:"InvoiceDiscount"`
		InvoiceInterval           int              `json:"InvoiceInterval"`
		InvoicesRemaining         string           `json:"InvoicesRemaining"`
		InvoiceRows               []InvoiceRow     `json:"InvoiceRows"`
		Language                  string           `json:"Language"`
		LastInvoiceDate           string           `json:"LastInvoiceDate"`
		Net                       int              `json:"Net"`
		OurReference              string           `json:"OurReference"`
		PeriodEnd                 string           `json:"PeriodEnd"`
		PeriodStart               string           `json:"PeriodStart"`
		PriceList                 string           `json:"PriceList"`
		PrintTemplate             string           `json:"PrintTemplate"`
		Project                   string           `json:"Project"`
		Remarks                   string           `json:"Remarks"`
		RoundOff                  int              `json:"RoundOff"`
		TaxReduction              int              `json:"TaxReduction"`
		TemplateName              string           `json:"TemplateName"`
		TemplateNumber            int              `json:"TemplateNumber"`
		TermsOfDelivery           string           `json:"TermsOfDelivery"`
		TermsOfPayment            string           `json:"TermsOfPayment"`
		Total                     int              `json:"Total"`
		TotalToPay                int              `json:"TotalToPay"`
		TotalVAT                  int              `json:"TotalVAT"`
		VatIncluded               bool             `json:"VatIncluded"`
		WayOfDelivery             string           `json:"WayOfDelivery"`
		YourOrderNumber           string           `json:"YourOrderNumber"`
		YourReference             string           `json:"YourReference"`
		TaxReductionType          string           `json:"TaxReductionType"`
	} `json:"Contract"`
}

type UpdateContractReq struct {
	Contract struct {
		Url                       string           `json:"@url"`
		UrlTaxReductionList       string           `json:"@urlTaxReductionList"`
		Active                    bool             `json:"Active"`
		AdministrationFee         int              `json:"AdministrationFee"`
		BasisTaxReduction         int              `json:"BasisTaxReduction"`
		Comments                  string           `json:"Comments"`
		Continuous                bool             `json:"Continuous"`
		ContractDate              string           `json:"ContractDate"`
		ContractLength            int              `json:"ContractLength"`
		ContributionPercent       int              `json:"ContributionPercent"`
		ContributionValue         int              `json:"ContributionValue"`
		CostCenter                string           `json:"CostCenter"`
		Currency                  string           `json:"Currency"`
		CustomerName              string           `json:"CustomerName"`
		CustomerNumber            string           `json:"CustomerNumber"`
		DocumentNumber            string           `json:"DocumentNumber"`
		EmailInformation          EmailInformation `json:"EmailInformation"`
		ExternalInvoiceReference1 string           `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string           `json:"ExternalInvoiceReference2"`
		Freight                   int              `json:"Freight"`
		Gross                     int              `json:"Gross"`
		HouseWork                 bool             `json:"HouseWork"`
		InvoiceDiscount           int              `json:"InvoiceDiscount"`
		InvoiceInterval           int              `json:"InvoiceInterval"`
		InvoicesRemaining         string           `json:"InvoicesRemaining"`
		InvoiceRows               []InvoiceRow     `json:"InvoiceRows"`
		Language                  string           `json:"Language"`
		LastInvoiceDate           string           `json:"LastInvoiceDate"`
		Net                       int              `json:"Net"`
		OurReference              string           `json:"OurReference"`
		PeriodEnd                 string           `json:"PeriodEnd"`
		PeriodStart               string           `json:"PeriodStart"`
		PriceList                 string           `json:"PriceList"`
		PrintTemplate             string           `json:"PrintTemplate"`
		Project                   string           `json:"Project"`
		Remarks                   string           `json:"Remarks"`
		RoundOff                  int              `json:"RoundOff"`
		TaxReduction              int              `json:"TaxReduction"`
		TemplateName              string           `json:"TemplateName"`
		TemplateNumber            int              `json:"TemplateNumber"`
		TermsOfDelivery           string           `json:"TermsOfDelivery"`
		TermsOfPayment            string           `json:"TermsOfPayment"`
		Total                     int              `json:"Total"`
		TotalToPay                int              `json:"TotalToPay"`
		TotalVAT                  int              `json:"TotalVAT"`
		VatIncluded               bool             `json:"VatIncluded"`
		WayOfDelivery             string           `json:"WayOfDelivery"`
		YourOrderNumber           string           `json:"YourOrderNumber"`
		YourReference             string           `json:"YourReference"`
		TaxReductionType          string           `json:"TaxReductionType"`
	} `json:"Contract"`
}

type UpdateContractResp struct {
	Contract struct {
		Url                       string           `json:"@url"`
		UrlTaxReductionList       string           `json:"@urlTaxReductionList"`
		Active                    bool             `json:"Active"`
		AdministrationFee         int              `json:"AdministrationFee"`
		BasisTaxReduction         int              `json:"BasisTaxReduction"`
		Comments                  string           `json:"Comments"`
		Continuous                bool             `json:"Continuous"`
		ContractDate              string           `json:"ContractDate"`
		ContractLength            int              `json:"ContractLength"`
		ContributionPercent       int              `json:"ContributionPercent"`
		ContributionValue         int              `json:"ContributionValue"`
		CostCenter                string           `json:"CostCenter"`
		Currency                  string           `json:"Currency"`
		CustomerName              string           `json:"CustomerName"`
		CustomerNumber            string           `json:"CustomerNumber"`
		DocumentNumber            string           `json:"DocumentNumber"`
		EmailInformation          EmailInformation `json:"EmailInformation"`
		ExternalInvoiceReference1 string           `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string           `json:"ExternalInvoiceReference2"`
		Freight                   int              `json:"Freight"`
		Gross                     int              `json:"Gross"`
		HouseWork                 bool             `json:"HouseWork"`
		InvoiceDiscount           int              `json:"InvoiceDiscount"`
		InvoiceInterval           int              `json:"InvoiceInterval"`
		InvoicesRemaining         string           `json:"InvoicesRemaining"`
		InvoiceRows               []InvoiceRow     `json:"InvoiceRows"`
		Language                  string           `json:"Language"`
		LastInvoiceDate           string           `json:"LastInvoiceDate"`
		Net                       int              `json:"Net"`
		OurReference              string           `json:"OurReference"`
		PeriodEnd                 string           `json:"PeriodEnd"`
		PeriodStart               string           `json:"PeriodStart"`
		PriceList                 string           `json:"PriceList"`
		PrintTemplate             string           `json:"PrintTemplate"`
		Project                   string           `json:"Project"`
		Remarks                   string           `json:"Remarks"`
		RoundOff                  int              `json:"RoundOff"`
		TaxReduction              int              `json:"TaxReduction"`
		TemplateName              string           `json:"TemplateName"`
		TemplateNumber            int              `json:"TemplateNumber"`
		TermsOfDelivery           string           `json:"TermsOfDelivery"`
		TermsOfPayment            string           `json:"TermsOfPayment"`
		Total                     int              `json:"Total"`
		TotalToPay                int              `json:"TotalToPay"`
		TotalVAT                  int              `json:"TotalVAT"`
		VatIncluded               bool             `json:"VatIncluded"`
		WayOfDelivery             string           `json:"WayOfDelivery"`
		YourOrderNumber           string           `json:"YourOrderNumber"`
		YourReference             string           `json:"YourReference"`
		TaxReductionType          string           `json:"TaxReductionType"`
	} `json:"Contract"`
}

type CreateContractReq struct {
	Contract struct {
		Url                       string           `json:"@url"`
		UrlTaxReductionList       string           `json:"@urlTaxReductionList"`
		Active                    bool             `json:"Active"`
		AdministrationFee         int              `json:"AdministrationFee"`
		BasisTaxReduction         int              `json:"BasisTaxReduction"`
		Comments                  string           `json:"Comments"`
		Continuous                bool             `json:"Continuous"`
		ContractDate              string           `json:"ContractDate"`
		ContractLength            int              `json:"ContractLength"`
		ContributionPercent       int              `json:"ContributionPercent"`
		ContributionValue         int              `json:"ContributionValue"`
		CostCenter                string           `json:"CostCenter"`
		Currency                  string           `json:"Currency"`
		CustomerName              string           `json:"CustomerName"`
		CustomerNumber            string           `json:"CustomerNumber"`
		DocumentNumber            string           `json:"DocumentNumber"`
		EmailInformation          EmailInformation `json:"EmailInformation"`
		ExternalInvoiceReference1 string           `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string           `json:"ExternalInvoiceReference2"`
		Freight                   int              `json:"Freight"`
		Gross                     int              `json:"Gross"`
		HouseWork                 bool             `json:"HouseWork"`
		InvoiceDiscount           int              `json:"InvoiceDiscount"`
		InvoiceInterval           int              `json:"InvoiceInterval"`
		InvoicesRemaining         string           `json:"InvoicesRemaining"`
		InvoiceRows               []InvoiceRow     `json:"InvoiceRows"`
		Language                  string           `json:"Language"`
		LastInvoiceDate           string           `json:"LastInvoiceDate"`
		Net                       int              `json:"Net"`
		OurReference              string           `json:"OurReference"`
		PeriodEnd                 string           `json:"PeriodEnd"`
		PeriodStart               string           `json:"PeriodStart"`
		PriceList                 string           `json:"PriceList"`
		PrintTemplate             string           `json:"PrintTemplate"`
		Project                   string           `json:"Project"`
		Remarks                   string           `json:"Remarks"`
		RoundOff                  int              `json:"RoundOff"`
		TaxReduction              int              `json:"TaxReduction"`
		TemplateName              string           `json:"TemplateName"`
		TemplateNumber            int              `json:"TemplateNumber"`
		TermsOfDelivery           string           `json:"TermsOfDelivery"`
		TermsOfPayment            string           `json:"TermsOfPayment"`
		Total                     int              `json:"Total"`
		TotalToPay                int              `json:"TotalToPay"`
		TotalVAT                  int              `json:"TotalVAT"`
		VatIncluded               bool             `json:"VatIncluded"`
		WayOfDelivery             string           `json:"WayOfDelivery"`
		YourOrderNumber           string           `json:"YourOrderNumber"`
		YourReference             string           `json:"YourReference"`
		TaxReductionType          string           `json:"TaxReductionType"`
	} `json:"Contract"`
}

type CreateContractResp struct {
	Contract struct {
		Url                       string           `json:"@url"`
		UrlTaxReductionList       string           `json:"@urlTaxReductionList"`
		Active                    bool             `json:"Active"`
		AdministrationFee         int              `json:"AdministrationFee"`
		BasisTaxReduction         int              `json:"BasisTaxReduction"`
		Comments                  string           `json:"Comments"`
		Continuous                bool             `json:"Continuous"`
		ContractDate              string           `json:"ContractDate"`
		ContractLength            int              `json:"ContractLength"`
		ContributionPercent       int              `json:"ContributionPercent"`
		ContributionValue         int              `json:"ContributionValue"`
		CostCenter                string           `json:"CostCenter"`
		Currency                  string           `json:"Currency"`
		CustomerName              string           `json:"CustomerName"`
		CustomerNumber            string           `json:"CustomerNumber"`
		DocumentNumber            string           `json:"DocumentNumber"`
		EmailInformation          EmailInformation `json:"EmailInformation"`
		ExternalInvoiceReference1 string           `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string           `json:"ExternalInvoiceReference2"`
		Freight                   int              `json:"Freight"`
		Gross                     int              `json:"Gross"`
		HouseWork                 bool             `json:"HouseWork"`
		InvoiceDiscount           int              `json:"InvoiceDiscount"`
		InvoiceInterval           int              `json:"InvoiceInterval"`
		InvoicesRemaining         string           `json:"InvoicesRemaining"`
		InvoiceRows               []InvoiceRow     `json:"InvoiceRows"`
		Language                  string           `json:"Language"`
		LastInvoiceDate           string           `json:"LastInvoiceDate"`
		Net                       int              `json:"Net"`
		OurReference              string           `json:"OurReference"`
		PeriodEnd                 string           `json:"PeriodEnd"`
		PeriodStart               string           `json:"PeriodStart"`
		PriceList                 string           `json:"PriceList"`
		PrintTemplate             string           `json:"PrintTemplate"`
		Project                   string           `json:"Project"`
		Remarks                   string           `json:"Remarks"`
		RoundOff                  int              `json:"RoundOff"`
		TaxReduction              int              `json:"TaxReduction"`
		TemplateName              string           `json:"TemplateName"`
		TemplateNumber            int              `json:"TemplateNumber"`
		TermsOfDelivery           string           `json:"TermsOfDelivery"`
		TermsOfPayment            string           `json:"TermsOfPayment"`
		Total                     int              `json:"Total"`
		TotalToPay                int              `json:"TotalToPay"`
		TotalVAT                  int              `json:"TotalVAT"`
		VatIncluded               bool             `json:"VatIncluded"`
		WayOfDelivery             string           `json:"WayOfDelivery"`
		YourOrderNumber           string           `json:"YourOrderNumber"`
		YourReference             string           `json:"YourReference"`
		TaxReductionType          string           `json:"TaxReductionType"`
	} `json:"Contract"`
}

type SetContractAsFinishedResp struct {
	Contract struct {
		Url                       string           `json:"@url"`
		UrlTaxReductionList       string           `json:"@urlTaxReductionList"`
		Active                    bool             `json:"Active"`
		AdministrationFee         int              `json:"AdministrationFee"`
		BasisTaxReduction         int              `json:"BasisTaxReduction"`
		Comments                  string           `json:"Comments"`
		Continuous                bool             `json:"Continuous"`
		ContractDate              string           `json:"ContractDate"`
		ContractLength            int              `json:"ContractLength"`
		ContributionPercent       int              `json:"ContributionPercent"`
		ContributionValue         int              `json:"ContributionValue"`
		CostCenter                string           `json:"CostCenter"`
		Currency                  string           `json:"Currency"`
		CustomerName              string           `json:"CustomerName"`
		CustomerNumber            string           `json:"CustomerNumber"`
		DocumentNumber            string           `json:"DocumentNumber"`
		EmailInformation          EmailInformation `json:"EmailInformation"`
		ExternalInvoiceReference1 string           `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string           `json:"ExternalInvoiceReference2"`
		Freight                   int              `json:"Freight"`
		Gross                     int              `json:"Gross"`
		HouseWork                 bool             `json:"HouseWork"`
		InvoiceDiscount           int              `json:"InvoiceDiscount"`
		InvoiceInterval           int              `json:"InvoiceInterval"`
		InvoicesRemaining         string           `json:"InvoicesRemaining"`
		InvoiceRows               []InvoiceRow     `json:"InvoiceRows"`
		Language                  string           `json:"Language"`
		LastInvoiceDate           string           `json:"LastInvoiceDate"`
		Net                       int              `json:"Net"`
		OurReference              string           `json:"OurReference"`
		PeriodEnd                 string           `json:"PeriodEnd"`
		PeriodStart               string           `json:"PeriodStart"`
		PriceList                 string           `json:"PriceList"`
		PrintTemplate             string           `json:"PrintTemplate"`
		Project                   string           `json:"Project"`
		Remarks                   string           `json:"Remarks"`
		RoundOff                  int              `json:"RoundOff"`
		TaxReduction              int              `json:"TaxReduction"`
		TemplateName              string           `json:"TemplateName"`
		TemplateNumber            int              `json:"TemplateNumber"`
		TermsOfDelivery           string           `json:"TermsOfDelivery"`
		TermsOfPayment            string           `json:"TermsOfPayment"`
		Total                     int              `json:"Total"`
		TotalToPay                int              `json:"TotalToPay"`
		TotalVAT                  int              `json:"TotalVAT"`
		VatIncluded               bool             `json:"VatIncluded"`
		WayOfDelivery             string           `json:"WayOfDelivery"`
		YourOrderNumber           string           `json:"YourOrderNumber"`
		YourReference             string           `json:"YourReference"`
		TaxReductionType          string           `json:"TaxReductionType"`
	} `json:"Contract"`
}

type CreateInvoiceFromContractResp struct {
	Invoice struct {
		Url                    string         `json:"@url"`
		UrlTaxReductionList    string         `json:"@urlTaxReductionList"`
		AdministrationFee      int            `json:"AdministrationFee"`
		AdministrationFeeVAT   int            `json:"AdministrationFeeVAT"`
		Address1               string         `json:"Address1"`
		Address2               string         `json:"Address2"`
		Balance                int            `json:"Balance"`
		BasisTaxReduction      int            `json:"BasisTaxReduction"`
		Booked                 bool           `json:"Booked"`
		Cancelled              bool           `json:"Cancelled"`
		City                   string         `json:"City"`
		Comments               string         `json:"Comments"`
		ContractReference      int            `json:"ContractReference"`
		ContributionPercent    int            `json:"ContributionPercent"`
		ContributionValue      int            `json:"ContributionValue"`
		Country                string         `json:"Country"`
		CostCenter             string         `json:"CostCenter"`
		Credit                 string         `json:"Credit"`
		CreditInvoiceReference string         `json:"CreditInvoiceReference"`
		Currency               string         `json:"Currency"`
		CurrencyRate           int            `json:"CurrencyRate"`
		CurrencyUnit           int            `json:"CurrencyUnit"`
		CustomerName           string         `json:"CustomerName"`
		CustomerNumber         string         `json:"CustomerNumber"`
		DeliveryAddress1       string         `json:"DeliveryAddress1"`
		DeliveryAddress2       string         `json:"DeliveryAddress2"`
		DeliveryCity           string         `json:"DeliveryCity"`
		DeliveryCountry        string         `json:"DeliveryCountry"`
		DeliveryDate           string         `json:"DeliveryDate"`
		DeliveryName           string         `json:"DeliveryName"`
		DeliveryZipCode        string         `json:"DeliveryZipCode"`
		DocumentNumber         string         `json:"DocumentNumber"`
		DueDate                string         `json:"DueDate"`
		EDIInformation         EDIInformation `json:"EDIInformation"`
		EmailInformation       struct {
			EmailAddressFrom string `json:"EmailAddressFrom"`
			EmailAddressTo   string `json:"EmailAddressTo"`
			EmailAddressCC   string `json:"EmailAddressCC"`
			EmailAddressBCC  string `json:"EmailAddressBCC"`
			EmailSubject     string `json:"EmailSubject"`
			EmailBody        string `json:"EmailBody"`
		} `json:"EmailInformation"`
		EUQuarterlyReport         bool         `json:"EUQuarterlyReport"`
		ExternalInvoiceReference1 string       `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string       `json:"ExternalInvoiceReference2"`
		Freight                   int          `json:"Freight"`
		FreightVAT                int          `json:"FreightVAT"`
		Gross                     int          `json:"Gross"`
		HouseWork                 bool         `json:"HouseWork"`
		InvoiceDate               string       `json:"InvoiceDate"`
		InvoicePeriodStart        string       `json:"InvoicePeriodStart"`
		InvoicePeriodEnd          string       `json:"InvoicePeriodEnd"`
		InvoicePeriodReference    string       `json:"InvoicePeriodReference"`
		InvoiceRows               []InvoiceRow `json:"InvoiceRows"`
		InvoiceType               string       `json:"InvoiceType"`
		Labels                    []Label      `json:"Labels"`
		Language                  string       `json:"Language"`
		LastRemindDate            string       `json:"LastRemindDate"`
		Net                       int          `json:"Net"`
		NotCompleted              bool         `json:"NotCompleted"`
		NoxFinans                 bool         `json:"NoxFinans"`
		OCR                       string       `json:"OCR"`
		OfferReference            string       `json:"OfferReference"`
		OrderReference            string       `json:"OrderReference"`
		OrganisationNumber        string       `json:"OrganisationNumber"`
		OurReference              string       `json:"OurReference"`
		PaymentWay                string       `json:"PaymentWay"`
		Phone1                    string       `json:"Phone1"`
		Phone2                    string       `json:"Phone2"`
		PriceList                 string       `json:"PriceList"`
		PrintTemplate             string       `json:"PrintTemplate"`
		Project                   string       `json:"Project"`
		WarehouseReady            bool         `json:"WarehouseReady"`
		OutboundDate              string       `json:"OutboundDate"`
		Remarks                   string       `json:"Remarks"`
		Reminders                 int          `json:"Reminders"`
		RoundOff                  int          `json:"RoundOff"`
		Sent                      bool         `json:"Sent"`
		TaxReduction              int          `json:"TaxReduction"`
		TermsOfDelivery           string       `json:"TermsOfDelivery"`
		TermsOfPayment            string       `json:"TermsOfPayment"`
		TimeBasisReference        int          `json:"TimeBasisReference"`
		Total                     int          `json:"Total"`
		TotalToPay                int          `json:"TotalToPay"`
		TotalVAT                  int          `json:"TotalVAT"`
		VATIncluded               bool         `json:"VATIncluded"`
		VoucherNumber             int          `json:"VoucherNumber"`
		VoucherSeries             string       `json:"VoucherSeries"`
		VoucherYear               int          `json:"VoucherYear"`
		WayOfDelivery             string       `json:"WayOfDelivery"`
		YourOrderNumber           string       `json:"YourOrderNumber"`
		YourReference             string       `json:"YourReference"`
		ZipCode                   string       `json:"ZipCode"`
		AccountingMethod          string       `json:"AccountingMethod"`
		TaxReductionType          string       `json:"TaxReductionType"`
		FinalPayDate              string       `json:"FinalPayDate"`
	} `json:"Invoice"`
}

type IncreaseInvoiceCountResp struct {
	Contract struct {
		Url                       string           `json:"@url"`
		UrlTaxReductionList       string           `json:"@urlTaxReductionList"`
		Active                    bool             `json:"Active"`
		AdministrationFee         int              `json:"AdministrationFee"`
		BasisTaxReduction         int              `json:"BasisTaxReduction"`
		Comments                  string           `json:"Comments"`
		Continuous                bool             `json:"Continuous"`
		ContractDate              string           `json:"ContractDate"`
		ContractLength            int              `json:"ContractLength"`
		ContributionPercent       int              `json:"ContributionPercent"`
		ContributionValue         int              `json:"ContributionValue"`
		CostCenter                string           `json:"CostCenter"`
		Currency                  string           `json:"Currency"`
		CustomerName              string           `json:"CustomerName"`
		CustomerNumber            string           `json:"CustomerNumber"`
		DocumentNumber            string           `json:"DocumentNumber"`
		EmailInformation          EmailInformation `json:"EmailInformation"`
		ExternalInvoiceReference1 string           `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string           `json:"ExternalInvoiceReference2"`
		Freight                   int              `json:"Freight"`
		Gross                     int              `json:"Gross"`
		HouseWork                 bool             `json:"HouseWork"`
		InvoiceDiscount           int              `json:"InvoiceDiscount"`
		InvoiceInterval           int              `json:"InvoiceInterval"`
		InvoicesRemaining         string           `json:"InvoicesRemaining"`
		InvoiceRows               []InvoiceRow     `json:"InvoiceRows"`
		Language                  string           `json:"Language"`
		LastInvoiceDate           string           `json:"LastInvoiceDate"`
		Net                       int              `json:"Net"`
		OurReference              string           `json:"OurReference"`
		PeriodEnd                 string           `json:"PeriodEnd"`
		PeriodStart               string           `json:"PeriodStart"`
		PriceList                 string           `json:"PriceList"`
		PrintTemplate             string           `json:"PrintTemplate"`
		Project                   string           `json:"Project"`
		Remarks                   string           `json:"Remarks"`
		RoundOff                  int              `json:"RoundOff"`
		TaxReduction              int              `json:"TaxReduction"`
		TemplateName              string           `json:"TemplateName"`
		TemplateNumber            int              `json:"TemplateNumber"`
		TermsOfDelivery           string           `json:"TermsOfDelivery"`
		TermsOfPayment            string           `json:"TermsOfPayment"`
		Total                     int              `json:"Total"`
		TotalToPay                int              `json:"TotalToPay"`
		TotalVAT                  int              `json:"TotalVAT"`
		VatIncluded               bool             `json:"VatIncluded"`
		WayOfDelivery             string           `json:"WayOfDelivery"`
		YourOrderNumber           string           `json:"YourOrderNumber"`
		YourReference             string           `json:"YourReference"`
		TaxReductionType          string           `json:"TaxReductionType"`
	} `json:"Contract"`
}

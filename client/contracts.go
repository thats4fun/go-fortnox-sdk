package client

import (
	"context"
	"fmt"
)

const (
	contactsURI = "contacts"
)

// GetContract does _GET https://api.fortnox.se/3/contracts/{DocumentNumber}
//
// documentNumber - identifies the contract accrual
func (c *Client) GetContract(ctx context.Context, documentNumber string) (*GetContractResp, error) {
	resp := &GetContractResp{}

	uri := fmt.Sprintf("%s/%s", contactsURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// RemoveContactAccrual does _DELETE
//
// documentNumber - identifies the contract accrual
func (c *Client) RemoveContactAccrual(ctx context.Context, documentNumber int) error {
	uri := fmt.Sprintf("%s/%d", contactsURI, documentNumber)
	return c._DELETE(ctx, uri)
}

type GetContractResp struct {
	Contract struct {
		Url                 string `json:"@url"`
		UrlTaxReductionList string `json:"@urlTaxReductionList"`
		Active              bool   `json:"Active"`
		AdministrationFee   int    `json:"AdministrationFee"`
		BasisTaxReduction   int    `json:"BasisTaxReduction"`
		Comments            string `json:"Comments"`
		Continuous          bool   `json:"Continuous"`
		ContractDate        string `json:"ContractDate"`
		ContractLength      int    `json:"ContractLength"`
		ContributionPercent int    `json:"ContributionPercent"`
		ContributionValue   int    `json:"ContributionValue"`
		CostCenter          string `json:"CostCenter"`
		Currency            string `json:"Currency"`
		CustomerName        string `json:"CustomerName"`
		CustomerNumber      string `json:"CustomerNumber"`
		DocumentNumber      string `json:"DocumentNumber"`
		EmailInformation    struct {
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
		Gross                     int    `json:"Gross"`
		HouseWork                 bool   `json:"HouseWork"`
		InvoiceDiscount           int    `json:"InvoiceDiscount"`
		InvoiceInterval           int    `json:"InvoiceInterval"`
		InvoicesRemaining         string `json:"InvoicesRemaining"`
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
			Total                  int    `json:"Total"`
			TotalExcludingVAT      int    `json:"TotalExcludingVAT"`
			Unit                   string `json:"Unit"`
			VAT                    int    `json:"VAT"`
		} `json:"InvoiceRows"`
		Language         string `json:"Language"`
		LastInvoiceDate  string `json:"LastInvoiceDate"`
		Net              int    `json:"Net"`
		OurReference     string `json:"OurReference"`
		PeriodEnd        string `json:"PeriodEnd"`
		PeriodStart      string `json:"PeriodStart"`
		PriceList        string `json:"PriceList"`
		PrintTemplate    string `json:"PrintTemplate"`
		Project          string `json:"Project"`
		Remarks          string `json:"Remarks"`
		RoundOff         int    `json:"RoundOff"`
		TaxReduction     int    `json:"TaxReduction"`
		TemplateName     string `json:"TemplateName"`
		TemplateNumber   int    `json:"TemplateNumber"`
		TermsOfDelivery  string `json:"TermsOfDelivery"`
		TermsOfPayment   string `json:"TermsOfPayment"`
		Total            int    `json:"Total"`
		TotalToPay       int    `json:"TotalToPay"`
		TotalVAT         int    `json:"TotalVAT"`
		VatIncluded      bool   `json:"VatIncluded"`
		WayOfDelivery    string `json:"WayOfDelivery"`
		YourOrderNumber  string `json:"YourOrderNumber"`
		YourReference    string `json:"YourReference"`
		TaxReductionType string `json:"TaxReductionType"`
	} `json:"Contract"`
}

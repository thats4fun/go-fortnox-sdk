package client

import "context"

const (
	companySettingsURI = "settings/company"
)

// GetCompanySettings does _GET https://api.fortnox.se/3/settings/company/
func (c *Client) GetCompanySettings(ctx context.Context) (*GetCompanySettingResp, error) {
	resp := &GetCompanySettingResp{}

	err := c._GET(ctx, companySettingsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type CompanySettings struct {
	Address            string `json:"Address"`
	BG                 string `json:"BG"`
	BIC                string `json:"BIC"`
	BranchCode         string `json:"BranchCode"`
	City               string `json:"City"`
	ContactFirstName   string `json:"ContactFirstName"`
	ContactLastName    string `json:"ContactLastName"`
	Country            string `json:"Country"`
	CountryCode        string `json:"CountryCode"`
	DatabaseNumber     string `json:"DatabaseNumber"`
	Domicile           string `json:"Domicile"`
	Email              string `json:"Email"`
	Fax                string `json:"Fax"`
	IBAN               string `json:"IBAN"`
	Name               string `json:"Name"`
	OrganizationNumber string `json:"OrganizationNumber"`
	PG                 string `json:"PG"`
	Phone1             string `json:"Phone1"`
	Phone2             string `json:"Phone2"`
	TaxEnabled         bool   `json:"TaxEnabled"`
	VATNumber          string `json:"VATNumber"`
	VisitAddress       string `json:"VisitAddress"`
	VisitCity          string `json:"VisitCity"`
	VisitCountry       string `json:"VisitCountry"`
	VisitCountryCode   string `json:"VisitCountryCode"`
	VisitName          string `json:"VisitName"`
	VisitZipCode       string `json:"VisitZipCode"`
	WWW                string `json:"WWW"`
	ZipCode            string `json:"ZipCode"`
}

type GetCompanySettingResp struct {
	CompanySettings CompanySettings `json:"CompanySettings"`
}

package client

import "context"

const (
	companySettingsURI = "settings/company"
)

// GetCompanySettings does _GET https://api.fortnox.se/3/settings/company/
func (c *Client) GetCompanySettings(ctx context.Context) (*CompanySettings, error) {
	resp := &GetCompanySettingResp{}

	err := c._GET(ctx, companySettingsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.CompanySettings, nil
}

type CompanySettings struct {
	Address            string `json:"Address,omitempty"`
	BG                 string `json:"BG,omitempty"`
	BIC                string `json:"BIC,omitempty"`
	BranchCode         string `json:"BranchCode,omitempty"`
	City               string `json:"City,omitempty"`
	ContactFirstName   string `json:"ContactFirstName,omitempty"`
	ContactLastName    string `json:"ContactLastName,omitempty"`
	Country            string `json:"Country,omitempty"`
	CountryCode        string `json:"CountryCode,omitempty"`
	DatabaseNumber     string `json:"DatabaseNumber,omitempty"`
	Domicile           string `json:"Domicile,omitempty"`
	Email              string `json:"Email,omitempty"`
	Fax                string `json:"Fax,omitempty"`
	IBAN               string `json:"IBAN,omitempty"`
	Name               string `json:"Name,omitempty"`
	OrganizationNumber string `json:"OrganizationNumber,omitempty"`
	PG                 string `json:"PG,omitempty"`
	Phone1             string `json:"Phone1,omitempty"`
	Phone2             string `json:"Phone2,omitempty"`
	TaxEnabled         bool   `json:"TaxEnabled,omitempty"`
	VATNumber          string `json:"VATNumber,omitempty"`
	VisitAddress       string `json:"VisitAddress,omitempty"`
	VisitCity          string `json:"VisitCity,omitempty"`
	VisitCountry       string `json:"VisitCountry,omitempty"`
	VisitCountryCode   string `json:"VisitCountryCode,omitempty"`
	VisitName          string `json:"VisitName,omitempty"`
	VisitZipCode       string `json:"VisitZipCode,omitempty"`
	WWW                string `json:"WWW,omitempty"`
	ZipCode            string `json:"ZipCode,omitempty"`
}

type GetCompanySettingResp struct {
	CompanySettings CompanySettings `json:"CompanySettings"`
}

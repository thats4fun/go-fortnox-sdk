package client

import "context"

const (
	companyInformationURI = "companyinformation"
)

// GetCompanyInformation does _GET https://api.fortnox.se/3/companyinformation
func (c Client) GetCompanyInformation(ctx context.Context) (*CompanyInformation, error) {
	resp := &GetCompanyInformationResp{}

	err := c._GET(ctx, companyInformationURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.CompanyInformation, nil
}

type CompanyInformation struct {
	Address            string `json:"Address,omitempty"`
	City               string `json:"City,omitempty"`
	CountryCode        string `json:"CountryCode,omitempty"`
	DatabaseNumber     int    `json:"DatabaseNumber,omitempty"`
	CompanyName        string `json:"CompanyName,omitempty"`
	OrganizationNumber string `json:"OrganizationNumber,omitempty"`
	VisitAddress       string `json:"VisitAddress,omitempty"`
	VisitCity          string `json:"VisitCity,omitempty"`
	VisitCountryCode   string `json:"VisitCountryCode,omitempty"`
	VisitZipCode       string `json:"VisitZipCode,omitempty"`
	ZipCode            string `json:"ZipCode,omitempty"`
}

type GetCompanyInformationResp struct {
	CompanyInformation CompanyInformation `json:"CompanyInformation"`
}

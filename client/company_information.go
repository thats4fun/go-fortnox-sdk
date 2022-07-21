package client

import "context"

const (
	companyInformationURI = "companyinformation"
)

// GetCompanyInformation does _GET https://api.fortnox.se/3/companyinformation
func (c Client) GetCompanyInformation(ctx context.Context) (*GetCompanyInformationResp, error) {
	resp := &GetCompanyInformationResp{}

	err := c._GET(ctx, companyInformationURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetCompanyInformationResp struct {
	CompanyInformation struct {
		Address            string `json:"Address"`
		City               string `json:"City"`
		CountryCode        string `json:"CountryCode"`
		DatabaseNumber     int    `json:"DatabaseNumber"`
		CompanyName        string `json:"CompanyName"`
		OrganizationNumber string `json:"OrganizationNumber"`
		VisitAddress       string `json:"VisitAddress"`
		VisitCity          string `json:"VisitCity"`
		VisitCountryCode   string `json:"VisitCountryCode"`
		VisitZipCode       string `json:"VisitZipCode"`
		ZipCode            string `json:"ZipCode"`
	} `json:"CompanyInformation"`
}

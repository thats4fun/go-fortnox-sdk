package client

import "context"

const (
	euVatLimitRegulationURI = "euvatlimitregulation/"
)

// GetEUVATLimitDetails does _GET https://api.fortnox.se/3/euvatlimitregulation/
func (c Client) GetEUVATLimitDetails(ctx context.Context) (*GetEUVATLimitDetailsResp, error) {
	resp := &GetEUVATLimitDetailsResp{}

	err := c._GET(ctx, euVatLimitRegulationURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetEUVATLimitDetailsResp struct {
	EUVatLimitRegulation struct {
		TotalExclVat int    `json:"TotalExclVat"`
		IsOverLimit  bool   `json:"IsOverLimit"`
		Limit        int    `json:"Limit"`
		Year         string `json:"Year"`
	} `json:"EUVatLimitRegulation"`
}

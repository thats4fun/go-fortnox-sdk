package client

import "context"

const (
	euVatLimitRegulationURI = "euvatlimitregulation/"
)

// GetEUVATLimitDetails does _GET https://api.fortnox.se/3/euvatlimitregulation/
func (c Client) GetEUVATLimitDetails(ctx context.Context) (*EUVatLimitRegulation, error) {
	resp := &GetEUVATLimitDetailsResp{}

	err := c._GET(ctx, euVatLimitRegulationURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.EUVatLimitRegulation, nil
}

type EUVatLimitRegulation struct {
	TotalExclVat int    `json:"TotalExclVat,omitempty"`
	IsOverLimit  bool   `json:"IsOverLimit,omitempty"`
	Limit        int    `json:"Limit,omitempty"`
	Year         string `json:"Year,omitempty"`
}

type GetEUVATLimitDetailsResp struct {
	EUVatLimitRegulation EUVatLimitRegulation `json:"EUVatLimitRegulation"`
}

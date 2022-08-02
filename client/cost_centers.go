package client

import (
	"context"
	"fmt"
)

const (
	costCentersURI = "costcenters"
)

// GetAllCostCenters does _GET https://api.fortnox.se/3/costcenters
func (c *Client) GetAllCostCenters(ctx context.Context) ([]CostCenter, error) {
	resp := &GetAllCostCentersResp{}

	err := c._GET(ctx, costCentersURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.CostCenters, nil
}

// CreateCostCenter does _POST https://api.fortnox.se/3/costcenters
//
// cc - cost center to create
func (c *Client) CreateCostCenter(ctx context.Context, cc *CostCenter) (*CostCenter, error) {
	req := &CreateCostCenterReq{CostCenter: *cc}
	resp := &CreateCostCenterResp{}

	err := c._POST(ctx, costCentersURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.CostCenter, nil
}

// GetCostCenter does _GET https://api.fortnox.se/3/costcenters/{Code}
//
// code - identifies the cost center
func (c Client) GetCostCenter(ctx context.Context, code string) (*CostCenter, error) {
	resp := &GetCostCenterResp{}

	uri := fmt.Sprintf("%s/%s", costCentersURI, code)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.CostCenter, nil
}

// UpdateCostCenter does _PUT https://api.fortnox.se/3/costcenters/{Code}
//
// code - identifies the cost center
//
// cc - cost center to update
func (c Client) UpdateCostCenter(ctx context.Context, code string, cc *CostCenter) (*CostCenter, error) {
	req := UpdateCostCenterReq{CostCenter: *cc}
	resp := &UpdateCostCenterResp{}

	uri := fmt.Sprintf("%s/%s", costCentersURI, code)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.CostCenter, nil
}

// RemoveCostCenter does _DELETE https://api.fortnox.se/3/costcenters/{Code}
//
// code - identifies the cost center to remove
func (c *Client) RemoveCostCenter(ctx context.Context, code string) error {
	uri := fmt.Sprintf("%s/%s", currenciesURI, code)
	return c._DELETE(ctx, uri)
}

type CostCenter struct {
	Url         string `json:"@url,omitempty"`
	Code        string `json:"Code,omitempty"`
	Description string `json:"Description,omitempty"`
	Note        string `json:"Note,omitempty"`
	Active      bool   `json:"Active,omitempty"`
}

type GetAllCostCentersResp struct {
	CostCenters []CostCenter `json:"CostCenters"`
}

type CreateCostCenterReq struct {
	CostCenter CostCenter `json:"CostCenter"`
}

type CreateCostCenterResp struct {
	CostCenter CostCenter `json:"CostCenter"`
}

type GetCostCenterResp struct {
	CostCenter CostCenter `json:"CostCenter"`
}

type UpdateCostCenterReq struct {
	CostCenter CostCenter `json:"CostCenter"`
}

type UpdateCostCenterResp struct {
	CostCenter CostCenter `json:"CostCenter"`
}

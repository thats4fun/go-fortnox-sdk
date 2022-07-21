package client

import (
	"context"
	"fmt"
)

const (
	costCentersURI = "costcenters"
)

// GetAllCostCenters does _GET https://api.fortnox.se/3/costcenters
func (c *Client) GetAllCostCenters(ctx context.Context) (*GetAllCostCentersResp, error) {
	resp := &GetAllCostCentersResp{}

	err := c._GET(ctx, costCentersURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateCostCenter does _POST https://api.fortnox.se/3/costcenters
//
// req - cost center to create
func (c *Client) CreateCostCenter(ctx context.Context, req *CreateCostCenterReq) (*CreateCostCenterResp, error) {
	resp := &CreateCostCenterResp{}

	err := c._POST(ctx, costCentersURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetCostCenter does _GET https://api.fortnox.se/3/costcenters/{Code}
//
// code - identifies the cost center
func (c Client) GetCostCenter(ctx context.Context, code string) (*GetCostCenterResp, error) {
	resp := &GetCostCenterResp{}

	uri := fmt.Sprintf("%s/%s", costCentersURI, code)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateCostCenter does _PUT https://api.fortnox.se/3/costcenters/{Code}
//
// code - identifies the cost center
//
// req - cost center to update
func (c Client) UpdateCostCenter(
	ctx context.Context, code string, req *UpdateCostCenterReq) (*UpdateCostCenterResp, error) {

	resp := &UpdateCostCenterResp{}

	uri := fmt.Sprintf("%s/%s", costCentersURI, code)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// RemoveCostCenter does _DELETE https://api.fortnox.se/3/costcenters/{Code}
//
// code - identifies the cost center to remove
func (c *Client) RemoveCostCenter(ctx context.Context, code string) error {
	uri := fmt.Sprintf("%s/%s", currenciesURI, code)
	return c._DELETE(ctx, uri)
}

type GetAllCostCentersResp struct {
	CostCenters []struct {
		Url         string `json:"@url"`
		Code        string `json:"Code"`
		Description string `json:"Description"`
		Note        string `json:"Note"`
		Active      bool   `json:"Active"`
	} `json:"CostCenters"`
}

type CreateCostCenterReq struct {
	CostCenter struct {
		Url         string `json:"@url"`
		Code        string `json:"Code"`
		Description string `json:"Description"`
		Note        string `json:"Note"`
		Active      bool   `json:"Active"`
	} `json:"CostCenter"`
}

type CreateCostCenterResp struct {
	CostCenter struct {
		Url         string `json:"@url"`
		Code        string `json:"Code"`
		Description string `json:"Description"`
		Note        string `json:"Note"`
		Active      bool   `json:"Active"`
	} `json:"CostCenter"`
}

type GetCostCenterResp struct {
	CostCenter struct {
		Url         string `json:"@url"`
		Code        string `json:"Code"`
		Description string `json:"Description"`
		Note        string `json:"Note"`
		Active      bool   `json:"Active"`
	} `json:"CostCenter"`
}

type UpdateCostCenterReq struct {
	CostCenter struct {
		Url         string `json:"@url"`
		Code        string `json:"Code"`
		Description string `json:"Description"`
		Note        string `json:"Note"`
		Active      bool   `json:"Active"`
	} `json:"CostCenter"`
}

type UpdateCostCenterResp struct {
	CostCenter struct {
		Url         string `json:"@url"`
		Code        string `json:"Code"`
		Description string `json:"Description"`
		Note        string `json:"Note"`
		Active      bool   `json:"Active"`
	} `json:"CostCenter"`
}

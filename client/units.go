package client

import (
	"context"
	"fmt"
)

const (
	unitsURI = "units"
)

// GetAllUnits does _GET https://api.fortnox.se/3/units
func (c *Client) GetAllUnits(ctx context.Context) ([]*Unit, error) {
	resp := GetAllUnitsResp{}

	err := c._GET(ctx, unitsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.Units, nil
}

// CreateUnit does _POST https://api.fortnox.se/3/units
//
// req - unit to create
func (c *Client) CreateUnit(ctx context.Context, req *CreateUnitReq) (*CreateUnitResp, error) {
	resp := &CreateUnitResp{}

	err := c._POST(ctx, unitsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetUnit does _GET https://api.fortnox.se/3/units/{Code}
//
// code - identifies the unit
func (c *Client) GetUnit(ctx context.Context, code string) (*GetUnitResp, error) {
	resp := &GetUnitResp{}

	uri := fmt.Sprintf("%s/%s", unitsURI, code)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateUnit does _PUT https://api.fortnox.se/3/units/{Code}
//
// code - identifies the unit
//
// req - unit to update
func (c *Client) UpdateUnit(ctx context.Context, code string, req *UpdateUnitReq) (*UpdateUnitResp, error) {
	resp := &UpdateUnitResp{}

	uri := fmt.Sprintf("%s/%s", unitsURI, code)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// RemoveUnit does _DELETE
func (c *Client) RemoveUnit(ctx context.Context, code string) error {
	uri := fmt.Sprintf("%s/%s", unitsURI, code)
	return c._DELETE(ctx, uri)
}

type GetAllUnitsResp struct {
	Units []*Unit `json:"Units"`
}

type Unit struct {
	Url         string `json:"@url"`
	Code        string `json:"Code"`
	Description string `json:"Description"`
	CodeEnglish string `json:"CodeEnglish"`
}

type CreateUnitReq struct {
	Unit struct {
		Url         string `json:"@url"`
		Code        string `json:"Code"`
		Description string `json:"Description"`
		CodeEnglish string `json:"CodeEnglish"`
	} `json:"Unit"`
}

type CreateUnitResp struct {
	Unit struct {
		Url         string `json:"@url"`
		Code        string `json:"Code"`
		Description string `json:"Description"`
		CodeEnglish string `json:"CodeEnglish"`
	} `json:"Unit"`
}

type GetUnitResp struct {
	Unit struct {
		Url         string `json:"@url"`
		Code        string `json:"Code"`
		Description string `json:"Description"`
		CodeEnglish string `json:"CodeEnglish"`
	} `json:"Unit"`
}

type UpdateUnitReq struct {
	Unit struct {
		Url         string `json:"@url"`
		Code        string `json:"Code"`
		Description string `json:"Description"`
		CodeEnglish string `json:"CodeEnglish"`
	} `json:"Unit"`
}

type UpdateUnitResp struct {
	Unit struct {
		Url         string `json:"@url"`
		Code        string `json:"Code"`
		Description string `json:"Description"`
		CodeEnglish string `json:"CodeEnglish"`
	} `json:"Unit"`
}

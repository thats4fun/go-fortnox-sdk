package client

import (
	"context"
	"fmt"
)

const labelsURI = "labels"

// GetAllLabels does _GET https://api.fortnox.se/3/labels
func (c Client) GetAllLabels(ctx context.Context) ([]Label, error) {
	resp := &GetAllLabelsResp{}

	err := c._GET(ctx, labelsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.Labels, nil
}

// CreateLabels does _POST https://api.fortnox.se/3/labels
//
// name - label to create
func (c *Client) CreateLabels(ctx context.Context, name string) (*Label, error) {
	resp := &CreateLabelResp{}

	req := &CreateLabelReq{
		Label: Label{
			Description: name,
		},
	}

	err := c._POST(ctx, labelsURI, nil, &req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Label, nil
}

// UpdateLabel does _PUT https://api.fortnox.se/3/labels/{Id}
//
// id - identifies the label
//
// name - to update
func (c *Client) UpdateLabel(ctx context.Context, id int, name string) (*Label, error) {
	resp := &UpdateLabelResp{}

	req := &UpdateLabelReq{
		Label: Label{
			Description: name,
		},
	}

	uri := fmt.Sprintf("%s/%d", labelsURI, id)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Label, nil
}

// DeleteLabel does _DELETE https://api.fortnox.se/3/labels/{Id}
//
// id - identifies the label
func (c *Client) DeleteLabel(ctx context.Context, id int) error {
	uri := fmt.Sprintf("%s/%d", labelsURI, id)
	return c._DELETE(ctx, uri)
}

type Label struct {
	ID          int    `json:"Id,omitempty"`
	Description string `json:"Description,omitempty"`
}

type GetAllLabelsResp struct {
	Labels []Label `json:"Labels"`
}

type CreateLabelReq struct {
	Label Label `json:"Label"`
}

type CreateLabelResp CreateLabelReq

type UpdateLabelReq CreateLabelReq

type UpdateLabelResp CreateLabelResp

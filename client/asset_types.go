package client

import (
	"context"
	"fmt"
)

const (
	assetTypesURI = "assets/types"
)

// GetAssetType does _GET https://api.fortnox.se/3/assets/types/{id}
//
// id - id
func (c *Client) GetAssetType(ctx context.Context, id int) (*GetAssetTypeResp, error) {
	resp := &GetAssetTypeResp{}

	uri := fmt.Sprintf("%s/%d", assetTypesURI, id)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateAssetType does _POST https://api.fortnox.se/3/assets/types/{id}
//
// id - id
//
// req - request
func (c *Client) CreateAssetType(ctx context.Context, id int, req *CreateAssetTypeReq) (*CreateAssetTypeResp, error) {
	resp := &CreateAssetTypeResp{}

	uri := fmt.Sprintf("%s/%d", assetTypesURI, id)

	err := c._POST(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateAssetType does _PUT https://api.fortnox.se/3/assets/types/{id}
//
// id - id
//
// req - request
func (c *Client) UpdateAssetType(ctx context.Context, id int, req *UpdateAssetTypeReq) (*UpdateAssetTypeResp, error) {
	resp := &UpdateAssetTypeResp{}

	uri := fmt.Sprintf("%s/%d", assetTypesURI, id)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteAssetType does _DELETE https://api.fortnox.se/3/assets/types/{id}
//
// id - id
func (c *Client) DeleteAssetType(ctx context.Context, id int) error {
	uri := fmt.Sprintf("%s/%d", assetTypesURI, id)
	return c._DELETE(ctx, uri)
}

// GetAllAssetTypes does _GET https://api.fortnox.se/3/assets/types/
func (c *Client) GetAllAssetTypes(ctx context.Context) (*GetAllAssetTypesResp, error) {
	resp := &GetAllAssetTypesResp{}

	err := c._GET(ctx, assetTypesURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetAssetTypeResp struct {
	Type struct {
		Url                   string `json:"@url"`
		Id                    int    `json:"Id"`
		Number                string `json:"Number"`
		Description           string `json:"Description"`
		Notes                 string `json:"Notes"`
		Type                  int    `json:"Type"`
		InUse                 bool   `json:"InUse"`
		AccountAssetId        int    `json:"AccountAssetId"`
		AccountValueLossId    int    `json:"AccountValueLossId"`
		AccountSaleLossId     int    `json:"AccountSaleLossId"`
		AccountSaleWinId      int    `json:"AccountSaleWinId"`
		AccountRevaluationId  int    `json:"AccountRevaluationId"`
		AccountWriteDownAckId int    `json:"AccountWriteDownAckId"`
		AccountWriteDownId    int    `json:"AccountWriteDownId"`
		AccountDepreciationId int    `json:"AccountDepreciationId"`
		AccountAsset          int    `json:"AccountAsset"`
		AccountValueLoss      int    `json:"AccountValueLoss"`
		AccountSaleLoss       int    `json:"AccountSaleLoss"`
		AccountSaleWin        int    `json:"AccountSaleWin"`
		AccountRevaluation    int    `json:"AccountRevaluation"`
		AccountWriteDownAck   int    `json:"AccountWriteDownAck"`
		AccountWriteDown      int    `json:"AccountWriteDown"`
		AccountDepreciation   int    `json:"AccountDepreciation"`
	} `json:"Type"`
}

type AssetTypeShort struct {
	Number                string `json:"Number"`
	Description           string `json:"Description"`
	Notes                 string `json:"Notes"`
	Type                  int    `json:"Type"`
	AccountAssetId        int    `json:"AccountAssetId"`
	AccountDepreciationId int    `json:"AccountDepreciationId"`
	AccountValueLossId    int    `json:"AccountValueLossId"`
}

type CreateAssetTypeReq struct {
	AssetType AssetTypeShort `json:"AssetType"`
}

type CreateAssetTypeResp struct {
	Type struct {
		Url                   string `json:"@url"`
		Id                    int    `json:"Id"`
		Number                string `json:"Number"`
		Description           string `json:"Description"`
		Notes                 string `json:"Notes"`
		Type                  int    `json:"Type"`
		InUse                 bool   `json:"InUse"`
		AccountAssetId        int    `json:"AccountAssetId"`
		AccountValueLossId    int    `json:"AccountValueLossId"`
		AccountSaleLossId     int    `json:"AccountSaleLossId"`
		AccountSaleWinId      int    `json:"AccountSaleWinId"`
		AccountRevaluationId  int    `json:"AccountRevaluationId"`
		AccountWriteDownAckId int    `json:"AccountWriteDownAckId"`
		AccountWriteDownId    int    `json:"AccountWriteDownId"`
		AccountDepreciationId int    `json:"AccountDepreciationId"`
		AccountAsset          int    `json:"AccountAsset"`
		AccountValueLoss      int    `json:"AccountValueLoss"`
		AccountSaleLoss       int    `json:"AccountSaleLoss"`
		AccountSaleWin        int    `json:"AccountSaleWin"`
		AccountRevaluation    int    `json:"AccountRevaluation"`
		AccountWriteDownAck   int    `json:"AccountWriteDownAck"`
		AccountWriteDown      int    `json:"AccountWriteDown"`
		AccountDepreciation   int    `json:"AccountDepreciation"`
	} `json:"Type"`
}

type UpdateAssetType struct {
	Description string `json:"Description"`
	Notes       string `json:"Notes"`
}
type UpdateAssetTypeReq struct {
	AssetType UpdateAssetType `json:"AssetType"`
}
type UpdateAssetTypeResp struct {
	Type struct {
		Url                   string `json:"@url"`
		Id                    int    `json:"Id"`
		Number                string `json:"Number"`
		Description           string `json:"Description"`
		Notes                 string `json:"Notes"`
		Type                  int    `json:"Type"`
		InUse                 bool   `json:"InUse"`
		AccountAssetId        int    `json:"AccountAssetId"`
		AccountValueLossId    int    `json:"AccountValueLossId"`
		AccountSaleLossId     int    `json:"AccountSaleLossId"`
		AccountSaleWinId      int    `json:"AccountSaleWinId"`
		AccountRevaluationId  int    `json:"AccountRevaluationId"`
		AccountWriteDownAckId int    `json:"AccountWriteDownAckId"`
		AccountWriteDownId    int    `json:"AccountWriteDownId"`
		AccountDepreciationId int    `json:"AccountDepreciationId"`
		AccountAsset          int    `json:"AccountAsset"`
		AccountValueLoss      int    `json:"AccountValueLoss"`
		AccountSaleLoss       int    `json:"AccountSaleLoss"`
		AccountSaleWin        int    `json:"AccountSaleWin"`
		AccountRevaluation    int    `json:"AccountRevaluation"`
		AccountWriteDownAck   int    `json:"AccountWriteDownAck"`
		AccountWriteDown      int    `json:"AccountWriteDown"`
		AccountDepreciation   int    `json:"AccountDepreciation"`
	} `json:"Type"`
}

type MetaInformation struct {
	TotalResources int `json:"@TotalResources"`
	TotalPages     int `json:"@TotalPages"`
	CurrentPage    int `json:"@CurrentPage"`
}

type GetAllAssetTypesResp struct {
	MetaInformation MetaInformation `json:"MetaInformation"`
	Types           []struct {
		Url                   string `json:"@url"`
		Id                    int    `json:"Id"`
		Number                string `json:"Number"`
		Description           string `json:"Description"`
		Notes                 string `json:"Notes"`
		Type                  int    `json:"Type"`
		InUse                 bool   `json:"InUse"`
		AccountAssetId        int    `json:"AccountAssetId"`
		AccountValueLossId    int    `json:"AccountValueLossId"`
		AccountSaleLossId     int    `json:"AccountSaleLossId"`
		AccountSaleWinId      int    `json:"AccountSaleWinId"`
		AccountRevaluationId  int    `json:"AccountRevaluationId"`
		AccountWriteDownAckId int    `json:"AccountWriteDownAckId"`
		AccountWriteDownId    int    `json:"AccountWriteDownId"`
		AccountDepreciationId int    `json:"AccountDepreciationId"`
		AccountAsset          int    `json:"AccountAsset"`
		AccountValueLoss      int    `json:"AccountValueLoss"`
		AccountSaleLoss       int    `json:"AccountSaleLoss"`
		AccountSaleWin        int    `json:"AccountSaleWin"`
		AccountRevaluation    int    `json:"AccountRevaluation"`
		AccountWriteDownAck   int    `json:"AccountWriteDownAck"`
		AccountWriteDown      int    `json:"AccountWriteDown"`
		AccountDepreciation   int    `json:"AccountDepreciation"`
	} `json:"Types"`
}

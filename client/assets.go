package client

import (
	"context"
	"fmt"
)

const (
	assetsURI = "assets"
)

// GetAllAssets does _GET https://api.fortnox.se/3/assets/
func (c *Client) GetAllAssets(ctx context.Context) (*GetAllAssetsResp, error) {
	resp := &GetAllAssetsResp{}

	err := c._GET(ctx, assetsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateAsset does _POST https://api.fortnox.se/3/assets/
//
// req - request
func (c *Client) CreateAsset(ctx context.Context, req *CreateAssetTypeReq) (*CreateAssetResp, error) {
	resp := &CreateAssetResp{}

	err := c._POST(ctx, assetsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetAsset does _GET https://api.fortnox.se/3/assets/{GivenNumber}
//
// givenNumber - asset number
func (c *Client) GetAsset(ctx context.Context, givenNumber string) (*GetAssetResp, error) {
	resp := &GetAssetResp{}

	uri := fmt.Sprintf("%s/%s", assetsURI, givenNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ChangeManualAssetOBValue does _PUT https://api.fortnox.se/3/assets/{GivenNumber}
//
// givenNumber - Asset number
//
// req - request
func (c *Client) ChangeManualAssetOBValue(
	ctx context.Context,
	givenNumber int,
	req *ChangeManualAssetOBValueReq) (*ChangeManualAssetOBValueResp, error) {

	resp := &ChangeManualAssetOBValueResp{}

	uri := fmt.Sprintf("%s/%d", assetsURI, givenNumber)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteOrVoidAsset does _DELETE https://api.fortnox.se/3/assets/{GivenNumber}
//
// givenNumber - asset number
//
// req - request
//
// TODO: pass to _DELETE req
func (c *Client) DeleteOrVoidAsset(ctx context.Context, givenNumber int, req *DeleteOrVoidAssetReq) error {
	uri := fmt.Sprintf("%s/%d", assetsURI, givenNumber)

	return c._DELETE(ctx, uri)
}

// GetAssetsDepreciationList does _GET https://api.fortnox.se/3/assets/depreciations/{ToDate}
//
// toDate - toDate
func (c *Client) GetAssetsDepreciationList(ctx context.Context, toDate string) (*GetAssetsDepreciationListResp, error) {
	resp := &GetAssetsDepreciationListResp{}

	uri := fmt.Sprintf("%s/depreciations/%s", assetsURI, toDate)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// WriteUpAsset does _PUT https://api.fortnox.se/3/assets/writeup/{GivenNumber}
//
// givenNumber - asset number
//
// req - request
func (c *Client) WriteUpAsset(ctx context.Context, givenNumber string, req *WriteUpAssetReq) (*WriteUpAssetResp, error) {
	resp := &WriteUpAssetResp{}

	uri := fmt.Sprintf("%s/writeup/%s", assetsURI, givenNumber)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// WriteDownAsset does _PUT https://api.fortnox.se/3/assets/writedown/{GivenNumber}
//
// givenNumber - asset number
//
// req - request
func (c *Client) WriteDownAsset(
	ctx context.Context,
	givenNumber string,
	req *WriteDownAssetReq) (*WriteDownAssetResp, error) {

	resp := &WriteDownAssetResp{}

	uri := fmt.Sprintf("%s/writedown/%s", assetsURI, givenNumber)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ScrapAsset does _PUT https://api.fortnox.se/3/assets/scrap/{GivenNumber}
//
// givenNumber - asset number
//
// req - request
func (c *Client) ScrapAsset(
	ctx context.Context,
	givenNumber string,
	req *ScrapAssetReq) (*ScrapAssetResp, error) {

	resp := &ScrapAssetResp{}

	uri := fmt.Sprintf("%s/scrap/%s", assetsURI, givenNumber)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SellAsset does _PUT https://api.fortnox.se/3/assets/sell/{GivenNumber}
//
// givenNumber - asset number
//
// req - request
func (c *Client) SellAsset(
	ctx context.Context,
	givenNumber string,
	req *SellAssetReq) (*SellAssetResp, error) {

	resp := &SellAssetResp{}

	uri := fmt.Sprintf("%s/sell/%s", assetsURI, givenNumber)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// PerformAssetDepreciation does _POST https://api.fortnox.se/3/assets/depreciate
//
// req - request
func (c *Client) PerformAssetDepreciation(
	ctx context.Context,
	req *PerformAssetDepreciationReq) (*PerformAssetDepreciationResp, error) {

	resp := &PerformAssetDepreciationResp{}

	uri := fmt.Sprintf("%s/depreciate", assetsURI)

	err := c._POST(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetAllAssetsResp struct {
	Assets []struct {
		Url               string `json:"@url"`
		Id                int    `json:"Id"`
		Number            string `json:"Number"`
		Description       string `json:"Description"`
		Status            string `json:"Status"`
		StatusId          string `json:"StatusId"`
		Type              string `json:"Type"`
		TypeId            int    `json:"TypeId"`
		AcquisitionValue  int    `json:"AcquisitionValue"`
		AcquisitionDate   string `json:"AcquisitionDate"`
		DepreciationFinal string `json:"DepreciationFinal"`
		DepreciatedTo     string `json:"DepreciatedTo"`
	} `json:"Assets"`
}

type CreateAssetReq struct {
	AssetType struct {
		Number                string `json:"Number"`
		Description           string `json:"Description"`
		Notes                 string `json:"Notes"`
		Type                  int    `json:"Type"`
		AccountAssetId        int    `json:"AccountAssetId"`
		AccountDepreciationId int    `json:"AccountDepreciationId"`
		AccountValueLossId    int    `json:"AccountValueLossId"`
	} `json:"AssetType"`
}

type CreateAssetResp struct {
	Assets struct {
		Url                       string `json:"@url"`
		Id                        int    `json:"Id"`
		Number                    string `json:"Number"`
		Description               string `json:"Description"`
		Status                    string `json:"Status"`
		StatusId                  string `json:"StatusId"`
		CostCenter                string `json:"CostCenter"`
		Project                   string `json:"Project"`
		Type                      string `json:"Type"`
		TypeId                    int    `json:"TypeId"`
		DepreciationMethod        int    `json:"DepreciationMethod"`
		AcquisitionValue          int    `json:"AcquisitionValue"`
		DepreciateToResidualValue int    `json:"DepreciateToResidualValue"`
		AcquisitionDate           string `json:"AcquisitionDate"`
		AcquisitionStart          string `json:"AcquisitionStart"`
		DepreciationFinal         string `json:"DepreciationFinal"`
		DepreciatedTo             string `json:"DepreciatedTo"`
		ManualOb                  int    `json:"ManualOb"`
		Notes                     string `json:"Notes"`
		Reference                 string `json:"Reference"`
		Brand                     string `json:"Brand"`
		InsuredNumber             string `json:"InsuredNumber"`
		InsuredWith               string `json:"InsuredWith"`
		Group                     string `json:"Group"`
		Room                      string `json:"Room"`
		Placement                 string `json:"Placement"`
		Department                string `json:"Department"`
		History                   []struct {
			Id              int    `json:"Id"`
			Date            string `json:"Date"`
			EventId         int    `json:"EventId"`
			Amount          string `json:"Amount"`
			UserId          int    `json:"UserId"`
			UserName        string `json:"UserName"`
			Notes           string `json:"Notes"`
			VoucherNumber   int    `json:"VoucherNumber"`
			VoucherSeries   string `json:"VoucherSeries"`
			VoucherYear     int    `json:"VoucherYear"`
			SupplierInvoice int    `json:"SupplierInvoice"`
		} `json:"History"`
	} `json:"Assets"`
}

type GetAssetResp struct {
	Assets struct {
		Url                       string `json:"@url"`
		Id                        int    `json:"Id"`
		Number                    string `json:"Number"`
		Description               string `json:"Description"`
		Status                    string `json:"Status"`
		StatusId                  string `json:"StatusId"`
		CostCenter                string `json:"CostCenter"`
		Project                   string `json:"Project"`
		Type                      string `json:"Type"`
		TypeId                    int    `json:"TypeId"`
		DepreciationMethod        int    `json:"DepreciationMethod"`
		AcquisitionValue          int    `json:"AcquisitionValue"`
		DepreciateToResidualValue int    `json:"DepreciateToResidualValue"`
		AcquisitionDate           string `json:"AcquisitionDate"`
		AcquisitionStart          string `json:"AcquisitionStart"`
		DepreciationFinal         string `json:"DepreciationFinal"`
		DepreciatedTo             string `json:"DepreciatedTo"`
		ManualOb                  int    `json:"ManualOb"`
		Notes                     string `json:"Notes"`
		Reference                 string `json:"Reference"`
		Brand                     string `json:"Brand"`
		InsuredNumber             string `json:"InsuredNumber"`
		InsuredWith               string `json:"InsuredWith"`
		Group                     string `json:"Group"`
		Room                      string `json:"Room"`
		Placement                 string `json:"Placement"`
		Department                string `json:"Department"`
		History                   []struct {
			Id              int    `json:"Id"`
			Date            string `json:"Date"`
			EventId         int    `json:"EventId"`
			Amount          string `json:"Amount"`
			UserId          int    `json:"UserId"`
			UserName        string `json:"UserName"`
			Notes           string `json:"Notes"`
			VoucherNumber   int    `json:"VoucherNumber"`
			VoucherSeries   string `json:"VoucherSeries"`
			VoucherYear     int    `json:"VoucherYear"`
			SupplierInvoice int    `json:"SupplierInvoice"`
		} `json:"History"`
	} `json:"Assets"`
}

type ChangeManualAssetOBValueReq struct {
	Amount  int    `json:"Amount"`
	Comment string `json:"Comment"`
}

type ChangeManualAssetOBValueResp struct {
	Assets struct {
		Url                       string `json:"@url"`
		Id                        int    `json:"Id"`
		Number                    string `json:"Number"`
		Description               string `json:"Description"`
		Status                    string `json:"Status"`
		StatusId                  string `json:"StatusId"`
		CostCenter                string `json:"CostCenter"`
		Project                   string `json:"Project"`
		Type                      string `json:"Type"`
		TypeId                    int    `json:"TypeId"`
		DepreciationMethod        int    `json:"DepreciationMethod"`
		AcquisitionValue          int    `json:"AcquisitionValue"`
		DepreciateToResidualValue int    `json:"DepreciateToResidualValue"`
		AcquisitionDate           string `json:"AcquisitionDate"`
		AcquisitionStart          string `json:"AcquisitionStart"`
		DepreciationFinal         string `json:"DepreciationFinal"`
		DepreciatedTo             string `json:"DepreciatedTo"`
		ManualOb                  int    `json:"ManualOb"`
		Notes                     string `json:"Notes"`
		Reference                 string `json:"Reference"`
		Brand                     string `json:"Brand"`
		InsuredNumber             string `json:"InsuredNumber"`
		InsuredWith               string `json:"InsuredWith"`
		Group                     string `json:"Group"`
		Room                      string `json:"Room"`
		Placement                 string `json:"Placement"`
		Department                string `json:"Department"`
		History                   []struct {
			Id              int    `json:"Id"`
			Date            string `json:"Date"`
			EventId         int    `json:"EventId"`
			Amount          string `json:"Amount"`
			UserId          int    `json:"UserId"`
			UserName        string `json:"UserName"`
			Notes           string `json:"Notes"`
			VoucherNumber   int    `json:"VoucherNumber"`
			VoucherSeries   string `json:"VoucherSeries"`
			VoucherYear     int    `json:"VoucherYear"`
			SupplierInvoice int    `json:"SupplierInvoice"`
		} `json:"History"`
	} `json:"Assets"`
}

type DeleteOrVoidAssetReq struct {
	Asset struct {
		Date string `json:"Date"`
	} `json:"Asset"`
}

type GetAssetsDepreciationListResp struct {
	Assets []struct {
		Url               string `json:"@url"`
		Id                int    `json:"Id"`
		Number            string `json:"Number"`
		Description       string `json:"Description"`
		Status            string `json:"Status"`
		StatusId          string `json:"StatusId"`
		Type              string `json:"Type"`
		TypeId            int    `json:"TypeId"`
		AcquisitionValue  int    `json:"AcquisitionValue"`
		AcquisitionDate   string `json:"AcquisitionDate"`
		DepreciationFinal string `json:"DepreciationFinal"`
		DepreciatedTo     string `json:"DepreciatedTo"`
	} `json:"Assets"`
}

type WriteUpAssetReq struct {
	Asset struct {
		Amount  int    `json:"Amount"`
		Comment string `json:"Comment"`
		Date    string `json:"Date"`
	} `json:"Asset"`
}

type WriteUpAssetResp struct {
	Assets struct {
		Url                       string `json:"@url"`
		Id                        int    `json:"Id"`
		Number                    string `json:"Number"`
		Description               string `json:"Description"`
		Status                    string `json:"Status"`
		StatusId                  string `json:"StatusId"`
		CostCenter                string `json:"CostCenter"`
		Project                   string `json:"Project"`
		Type                      string `json:"Type"`
		TypeId                    int    `json:"TypeId"`
		DepreciationMethod        int    `json:"DepreciationMethod"`
		AcquisitionValue          int    `json:"AcquisitionValue"`
		DepreciateToResidualValue int    `json:"DepreciateToResidualValue"`
		AcquisitionDate           string `json:"AcquisitionDate"`
		AcquisitionStart          string `json:"AcquisitionStart"`
		DepreciationFinal         string `json:"DepreciationFinal"`
		DepreciatedTo             string `json:"DepreciatedTo"`
		ManualOb                  int    `json:"ManualOb"`
		Notes                     string `json:"Notes"`
		Reference                 string `json:"Reference"`
		Brand                     string `json:"Brand"`
		InsuredNumber             string `json:"InsuredNumber"`
		InsuredWith               string `json:"InsuredWith"`
		Group                     string `json:"Group"`
		Room                      string `json:"Room"`
		Placement                 string `json:"Placement"`
		Department                string `json:"Department"`
		History                   []struct {
			Id              int    `json:"Id"`
			Date            string `json:"Date"`
			EventId         int    `json:"EventId"`
			Amount          string `json:"Amount"`
			UserId          int    `json:"UserId"`
			UserName        string `json:"UserName"`
			Notes           string `json:"Notes"`
			VoucherNumber   int    `json:"VoucherNumber"`
			VoucherSeries   string `json:"VoucherSeries"`
			VoucherYear     int    `json:"VoucherYear"`
			SupplierInvoice int    `json:"SupplierInvoice"`
		} `json:"History"`
	} `json:"Assets"`
}

type WriteDownAssetReq struct {
	Asset struct {
		Amount  int    `json:"Amount"`
		Comment string `json:"Comment"`
		Date    string `json:"Date"`
	} `json:"Asset"`
}

type WriteDownAssetResp struct {
	Assets struct {
		Url                       string `json:"@url"`
		Id                        int    `json:"Id"`
		Number                    string `json:"Number"`
		Description               string `json:"Description"`
		Status                    string `json:"Status"`
		StatusId                  string `json:"StatusId"`
		CostCenter                string `json:"CostCenter"`
		Project                   string `json:"Project"`
		Type                      string `json:"Type"`
		TypeId                    int    `json:"TypeId"`
		DepreciationMethod        int    `json:"DepreciationMethod"`
		AcquisitionValue          int    `json:"AcquisitionValue"`
		DepreciateToResidualValue int    `json:"DepreciateToResidualValue"`
		AcquisitionDate           string `json:"AcquisitionDate"`
		AcquisitionStart          string `json:"AcquisitionStart"`
		DepreciationFinal         string `json:"DepreciationFinal"`
		DepreciatedTo             string `json:"DepreciatedTo"`
		ManualOb                  int    `json:"ManualOb"`
		Notes                     string `json:"Notes"`
		Reference                 string `json:"Reference"`
		Brand                     string `json:"Brand"`
		InsuredNumber             string `json:"InsuredNumber"`
		InsuredWith               string `json:"InsuredWith"`
		Group                     string `json:"Group"`
		Room                      string `json:"Room"`
		Placement                 string `json:"Placement"`
		Department                string `json:"Department"`
		History                   []struct {
			Id              int    `json:"Id"`
			Date            string `json:"Date"`
			EventId         int    `json:"EventId"`
			Amount          string `json:"Amount"`
			UserId          int    `json:"UserId"`
			UserName        string `json:"UserName"`
			Notes           string `json:"Notes"`
			VoucherNumber   int    `json:"VoucherNumber"`
			VoucherSeries   string `json:"VoucherSeries"`
			VoucherYear     int    `json:"VoucherYear"`
			SupplierInvoice int    `json:"SupplierInvoice"`
		} `json:"History"`
	} `json:"Assets"`
}

type ScrapAssetReq struct {
	Asset struct {
		Percentage int    `json:"Percentage"`
		Comment    string `json:"Comment"`
		Date       string `json:"Date"`
	} `json:"Asset"`
}

type ScrapAssetResp struct {
	Assets struct {
		Url                       string `json:"@url"`
		Id                        int    `json:"Id"`
		Number                    string `json:"Number"`
		Description               string `json:"Description"`
		Status                    string `json:"Status"`
		StatusId                  string `json:"StatusId"`
		CostCenter                string `json:"CostCenter"`
		Project                   string `json:"Project"`
		Type                      string `json:"Type"`
		TypeId                    int    `json:"TypeId"`
		DepreciationMethod        int    `json:"DepreciationMethod"`
		AcquisitionValue          int    `json:"AcquisitionValue"`
		DepreciateToResidualValue int    `json:"DepreciateToResidualValue"`
		AcquisitionDate           string `json:"AcquisitionDate"`
		AcquisitionStart          string `json:"AcquisitionStart"`
		DepreciationFinal         string `json:"DepreciationFinal"`
		DepreciatedTo             string `json:"DepreciatedTo"`
		ManualOb                  int    `json:"ManualOb"`
		Notes                     string `json:"Notes"`
		Reference                 string `json:"Reference"`
		Brand                     string `json:"Brand"`
		InsuredNumber             string `json:"InsuredNumber"`
		InsuredWith               string `json:"InsuredWith"`
		Group                     string `json:"Group"`
		Room                      string `json:"Room"`
		Placement                 string `json:"Placement"`
		Department                string `json:"Department"`
		History                   []struct {
			Id              int    `json:"Id"`
			Date            string `json:"Date"`
			EventId         int    `json:"EventId"`
			Amount          string `json:"Amount"`
			UserId          int    `json:"UserId"`
			UserName        string `json:"UserName"`
			Notes           string `json:"Notes"`
			VoucherNumber   int    `json:"VoucherNumber"`
			VoucherSeries   string `json:"VoucherSeries"`
			VoucherYear     int    `json:"VoucherYear"`
			SupplierInvoice int    `json:"SupplierInvoice"`
		} `json:"History"`
	} `json:"Assets"`
}

type SellAssetReq struct {
	Asset struct {
		Percentage int    `json:"Percentage"`
		Price      int    `json:"Price"`
		Comment    string `json:"Comment"`
		Date       string `json:"Date"`
	} `json:"Asset"`
}

type SellAssetResp struct {
	Assets struct {
		Url                       string `json:"@url"`
		Id                        int    `json:"Id"`
		Number                    string `json:"Number"`
		Description               string `json:"Description"`
		Status                    string `json:"Status"`
		StatusId                  string `json:"StatusId"`
		CostCenter                string `json:"CostCenter"`
		Project                   string `json:"Project"`
		Type                      string `json:"Type"`
		TypeId                    int    `json:"TypeId"`
		DepreciationMethod        int    `json:"DepreciationMethod"`
		AcquisitionValue          int    `json:"AcquisitionValue"`
		DepreciateToResidualValue int    `json:"DepreciateToResidualValue"`
		AcquisitionDate           string `json:"AcquisitionDate"`
		AcquisitionStart          string `json:"AcquisitionStart"`
		DepreciationFinal         string `json:"DepreciationFinal"`
		DepreciatedTo             string `json:"DepreciatedTo"`
		ManualOb                  int    `json:"ManualOb"`
		Notes                     string `json:"Notes"`
		Reference                 string `json:"Reference"`
		Brand                     string `json:"Brand"`
		InsuredNumber             string `json:"InsuredNumber"`
		InsuredWith               string `json:"InsuredWith"`
		Group                     string `json:"Group"`
		Room                      string `json:"Room"`
		Placement                 string `json:"Placement"`
		Department                string `json:"Department"`
		History                   []struct {
			Id              int    `json:"Id"`
			Date            string `json:"Date"`
			EventId         int    `json:"EventId"`
			Amount          string `json:"Amount"`
			UserId          int    `json:"UserId"`
			UserName        string `json:"UserName"`
			Notes           string `json:"Notes"`
			VoucherNumber   int    `json:"VoucherNumber"`
			VoucherSeries   string `json:"VoucherSeries"`
			VoucherYear     int    `json:"VoucherYear"`
			SupplierInvoice int    `json:"SupplierInvoice"`
		} `json:"History"`
	} `json:"Assets"`
}

type PerformAssetDepreciationReq struct {
	Asset struct {
		DepreciateUntil string `json:"DepreciateUntil"`
		AssetIds        []int  `json:"AssetIds"`
	} `json:"Asset"`
}

type PerformAssetDepreciationResp struct {
	AssetsDepreciation []struct {
		Url           string `json:"@url"`
		VoucherNumber int    `json:"VoucherNumber"`
		VoucherSeries string `json:"VoucherSeries"`
		FinancialYear int    `json:"FinancialYear"`
	} `json:"AssetsDepreciation"`
}

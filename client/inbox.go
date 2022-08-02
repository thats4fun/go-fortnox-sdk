package client

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

const (
	inboxURI = "inbox"
)

var (
	_PathToDescription = map[string]string{
		"Inbox_a":  "Asset register",
		"Inbox_d":  "Daily takings",
		"Inbox_s":  "Supplier invoices",
		"Inbox_v":  "Vouchers",
		"Inbox_b":  "Bank files",
		"Inbox_l":  "Payroll files",
		"Inbox_kf": "Customer invoices",
		"Inbox_o":  "Orders",
		"Inbox_of": "Offers",
	}
)

// GetRootDirectory does _GET https://api.fortnox.se/3/inbox/
func (c *Client) GetRootDirectory(ctx context.Context) (*InboxRootFolder, error) {
	resp := GetRootDirectoryResp{}

	err := c._GET(ctx, inboxURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Folder, nil
}

// UploadFile does _POST https://api.fortnox.se/3/inbox/
//
// folderID - folder id
//
// path - path
//
// f - file to upload
func (c *Client) UploadFile(ctx context.Context, params *UploadFileParams, f *InboxFile) (*InboxFile, error) {
	req := &UploadFileReq{File: *f}
	resp := &UploadFileResp{}

	p := params.urlValues()

	err := c._POST(ctx, inboxURI, p, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.File, nil
}

// GetInboxFile does _GET https://api.fortnox.se/3/inbox/{Id}
//
// id - identifies the folder
func (c *Client) GetInboxFile(ctx context.Context, id string) (*[]byte, error) {
	resp := []byte{}

	uri := fmt.Sprintf("%s/%s", inboxURI, id)

	err := c._GET(ctx, uri, nil, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// RemoveFileOrFolder does _DELETE https://api.fortnox.se/3/inbox/{Id}
//
// id - identifies the file to remove
func (c *Client) RemoveFileOrFolder(ctx context.Context, id string) error {
	uri := fmt.Sprintf("%s/%s", inboxURI, id)
	return c._DELETE(ctx, uri)
}

type UploadFileParams struct {
	FolderID string
	Path     string
}

func (p *UploadFileParams) urlValues() url.Values {
	urlValues := url.Values{}

	if strings.TrimSpace(p.FolderID) != "" {
		urlValues["folderId"] = []string{p.FolderID}
	}

	if strings.TrimSpace(p.Path) != "" {
		urlValues["path"] = []string{p.Path}
	}

	return urlValues
}

type GetRootDirectoryResp struct {
	Folder InboxRootFolder `json:"Folder"`
}

type InboxRootFolder struct {
	Url     string             `json:"@url,omitempty"`
	Email   string             `json:"Email,omitempty"`
	Files   []InboxFile        `json:"Files,omitempty"`
	Folders []InnerInboxFolder `json:"Folders,omitempty"`
	Id      string             `json:"Id,omitempty"`
	Name    string             `json:"Name,omitempty"`
}

type InboxFile struct {
	Url           string `json:"@url,omitempty"`
	Comments      string `json:"Comments,omitempty"`
	Id            string `json:"Id,omitempty"`
	Name          string `json:"Name,omitempty"`
	Path          string `json:"Path,omitempty"`
	Size          int    `json:"Size,omitempty"`
	ArchiveFileId string `json:"ArchiveFileId,omitempty"`
}

type InnerInboxFolder struct {
	Url  string `json:"@url,omitempty"`
	Id   string `json:"Id,omitempty"`
	Name string `json:"Name,omitempty"`
}

type UploadFileReq struct {
	File InboxFile `json:"File"`
}

type UploadFileResp struct {
	File InboxFile `json:"File"`
}

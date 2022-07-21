package client

import (
	"context"
	"fmt"
)

const (
	inboxURI = "inbox"
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

// UploadFile does https://api.fortnox.se/3/inbox/
// folderID - folder id
//
// path - path
//
// req - file to upload
func (c *Client) UploadFile(ctx context.Context, folderID, path string, req *UploadFileReq) (*InboxRootFolder, error) {
	resp := GetRootDirectoryResp{}

	err := c._GET(ctx, inboxURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Folder, nil
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

type GetRootDirectoryResp struct {
	Folder InboxRootFolder `json:"Folder"`
}

type InboxRootFolder struct {
	Url     string             `json:"@url"`
	Email   string             `json:"Email"`
	Files   []InboxFile        `json:"Files"`
	Folders []InnerInboxFolder `json:"Folders"`
	Id      string             `json:"Id"`
	Name    string             `json:"Name"`
}

type InboxFile struct {
	Url           string `json:"@url"`
	Comments      string `json:"Comments"`
	Id            string `json:"Id"`
	Name          string `json:"Name"`
	Path          string `json:"Path"`
	Size          int    `json:"Size"`
	ArchiveFileId string `json:"ArchiveFileId"`
}

type InnerInboxFolder struct {
	Url  string `json:"@url"`
	Id   string `json:"Id"`
	Name string `json:"Name"`
}

type UploadFileReq struct {
	File InboxFile `json:"File"`
}

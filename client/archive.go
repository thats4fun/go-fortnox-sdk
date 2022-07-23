package client

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

// URI
const (
	archiveURI = "archive"
)

const (
	pathParamName   = "path"
	fileIDParamName = "fileid"
)

// GetFileOrFolder does _GET https://api.fortnox.se/3/archive/
// filter:
//
// 1. path - name of folder
//
// 2. fileID - fileId from fileattachments
func (c *Client) GetFileOrFolder(ctx context.Context, filter *PathFileIDFilter) (*GetFileOrFolderResp, error) {
	resp := &GetFileOrFolderResp{}

	params := filter.urlValues()

	err := c._GET(ctx, archiveURI, params, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UploadFileToDir does _POST https://api.fortnox.se/3/archive/
//
// file - file to upload
//
// filter:
//
// 1. path - name of folder
//
// 2. folderID - if of folder
func (c Client) UploadFileToDir(ctx context.Context, filter *PathFileIDFilter, file *File) (*UploadFileResp, error) {
	resp := &UploadFileResp{}

	params := filter.urlValues()

	err := c._POST(ctx, archiveURI, params, file, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// RemoveFiles does _DELETE https://api.fortnox.se/3/archive/
//
// path - identifies file/folder to remove
func (c Client) RemoveFiles(ctx context.Context, path string) error {
	uri := archiveURI

	if strings.TrimSpace(path) == "" {
		return errors.New("cant delete without path")
	}

	uri = fmt.Sprintf("%s?%s=%s", archiveURI, pathParamName, path)

	return c._DELETE(ctx, uri)
}

// GetFile does _GET https://api.fortnox.se/3/archive/{id}
//
// id - identifies the file
//
// fileID - fileId from fileattachments
func (c *Client) GetFile(ctx context.Context, id, fileID string) (*[]byte, error) {
	resp := &[]byte{}

	if strings.TrimSpace(id) == "" {
		return nil, errors.New("can't get file without id")
	}

	uri := fmt.Sprintf("%s/%s", archiveURI, id)

	var params url.Values
	if strings.TrimSpace(fileID) != "" {
		params[fileIDParamName] = []string{fileID}
	}

	err := c._GET(ctx, uri, params, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteFile does _DELETE https://api.fortnox.se/3/archive/{id}
//
// id - identifies the file
func (c *Client) DeleteFile(ctx context.Context, id string) error {
	if strings.TrimSpace(id) == "" {
		return errors.New("can't delete without id")
	}
	uri := fmt.Sprintf("%s/%s", archiveURI, id)
	return c._DELETE(ctx, uri)
}

type PathFileIDFilter struct {
	Path   string
	FileID string
}

func (f *PathFileIDFilter) urlValues() url.Values {
	if f == nil {
		return nil
	}

	params := url.Values{}

	if strings.TrimSpace(f.Path) != "" {
		params[pathParamName] = []string{f.Path}
	}

	if strings.TrimSpace(f.FileID) != "" {
		params[fileIDParamName] = []string{fileIDParamName}
	}

	return params
}

type Folder struct {
	Url     string `json:"@url"`
	Email   string `json:"Email"`
	Files   []File `json:"Files"`
	Folders []struct {
		Url  string `json:"@url"`
		Id   string `json:"Id"`
		Name string `json:"Name"`
	} `json:"Folders"`
	Id   string `json:"Id"`
	Name string `json:"Name"`
}

type File struct {
	Url           string `json:"@url"`
	Comments      string `json:"Comments"`
	Id            string `json:"Id"`
	Name          string `json:"Name"`
	Path          string `json:"Path"`
	Size          int    `json:"Size"`
	ArchiveFileId string `json:"ArchiveFileId"`
}

type GetFileOrFolderResp struct {
	Folder struct {
		Url   string `json:"@url"`
		Email string `json:"Email"`
		Files []struct {
			Url           string `json:"@url"`
			Comments      string `json:"Comments"`
			Id            string `json:"Id"`
			Name          string `json:"Name"`
			Path          string `json:"Path"`
			Size          int    `json:"Size"`
			ArchiveFileId string `json:"ArchiveFileId"`
		} `json:"Files"`
		Folders []struct {
			Url  string `json:"@url"`
			Id   string `json:"Id"`
			Name string `json:"Name"`
		} `json:"Folders"`
		Id   string `json:"Id"`
		Name string `json:"Name"`
	} `json:"Folder"`
}

type UploadFileResp struct {
	File struct {
		Url           string `json:"@url"`
		Comments      string `json:"Comments"`
		Id            string `json:"Id"`
		Name          string `json:"Name"`
		Path          string `json:"Path"`
		Size          int    `json:"Size"`
		ArchiveFileId string `json:"ArchiveFileId"`
	} `json:"File"`
}

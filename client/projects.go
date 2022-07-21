package client

import (
	"context"
	"fmt"
	"strconv"
)

const (
	projectsURI = "projects"
)

// GetProject does _GET https://api.fortnox.se/3/projects/{ProjectNumber}
//
// projectNumber - identifies the project
func (c *Client) GetProject(ctx context.Context, projectNumber int) (*GetProjectResp, error) {
	resp := &GetProjectResp{}

	uri := fmt.Sprintf("%s/%d", projectsURI, projectNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateProject does _PUT https://api.fortnox.se/3/projects/{ProjectNumber}
//
// projectNumber - identifies the project
//
// req - project to update
func (c *Client) UpdateProject(ctx context.Context, projectNumber int, req *UpdateProjectReq) (*UpdateProjectReq, error) {
	resp := &UpdateProjectReq{}

	uri := fmt.Sprintf("%s/%d", projectsURI, projectNumber)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// RemoveProject does _DELETE https://api.fortnox.se/3/projects/{ProjectNumber}
func (c *Client) RemoveProject(ctx context.Context, projectNumber int) error {
	uri := fmt.Sprintf("%s/%d", projectsURI, projectNumber)
	return c._DELETE(ctx, uri)
}

// GetAllProjects does _GET https://api.fortnox.se/3/projects
func (c *Client) GetAllProjects(ctx context.Context) (*GetAllProjectsResp, error) {
	resp := &GetAllProjectsResp{}

	err := c._GET(ctx, projectsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateProject does POST https://api.fortnox.se/3/projects/
//
// req - project to create
func (c *Client) CreateProject(ctx context.Context, req *CreateProjectReq) (*CreateProjectResp, error) {
	resp := &CreateProjectResp{}

	err := c._POST(ctx, projectsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetProjectResp struct {
	Project struct {
		Url           string `json:"@url"`
		Comments      string `json:"Comments"`
		ContactPerson string `json:"ContactPerson"`
		Description   string `json:"Description"`
		EndDate       string `json:"EndDate"`
		ProjectLeader string `json:"ProjectLeader"`
		ProjectNumber string `json:"ProjectNumber"`
		Status        string `json:"Status"`
		StartDate     string `json:"StartDate"`
	} `json:"Project"`
}

type UpdateProjectReq struct {
	Project struct {
		Url           string `json:"@url"`
		Comments      string `json:"Comments"`
		ContactPerson string `json:"ContactPerson"`
		Description   string `json:"Description"`
		EndDate       string `json:"EndDate"`
		ProjectLeader string `json:"ProjectLeader"`
		ProjectNumber string `json:"ProjectNumber"`
		Status        string `json:"Status"`
		StartDate     string `json:"StartDate"`
	} `json:"Project"`
}

type UpdateProjectResp struct {
	Project struct {
		Url           string `json:"@url"`
		Comments      string `json:"Comments"`
		ContactPerson string `json:"ContactPerson"`
		Description   string `json:"Description"`
		EndDate       string `json:"EndDate"`
		ProjectLeader string `json:"ProjectLeader"`
		ProjectNumber string `json:"ProjectNumber"`
		Status        string `json:"Status"`
		StartDate     string `json:"StartDate"`
	} `json:"Project"`
}

type GetAllProjectsResp struct {
	Projects []struct {
		Url           string `json:"@url"`
		Description   string `json:"Description"`
		EndDate       string `json:"EndDate"`
		ProjectLeader string `json:"ProjectLeader"`
		ProjectNumber string `json:"ProjectNumber"`
		Status        string `json:"Status"`
		StartDate     string `json:"StartDate"`
	} `json:"Projects"`
}

type CreateProjectReq struct {
	Project struct {
		Url           string `json:"@url"`
		Comments      string `json:"Comments"`
		ContactPerson string `json:"ContactPerson"`
		Description   string `json:"Description"`
		EndDate       string `json:"EndDate"`
		ProjectLeader string `json:"ProjectLeader"`
		ProjectNumber string `json:"ProjectNumber"`
		Status        string `json:"Status"`
		StartDate     string `json:"StartDate"`
	} `json:"Project"`
}

type CreateProjectResp struct {
	Project struct {
		Url           string `json:"@url"`
		Comments      string `json:"Comments"`
		ContactPerson string `json:"ContactPerson"`
		Description   string `json:"Description"`
		EndDate       string `json:"EndDate"`
		ProjectLeader string `json:"ProjectLeader"`
		ProjectNumber string `json:"ProjectNumber"`
		Status        string `json:"Status"`
		StartDate     string `json:"StartDate"`
	} `json:"Project"`
}

// GetProjectNumberAsInt can be removed at any time
func (resp *CreateProjectResp) GetProjectNumberAsInt() int {
	projNumber, _ := strconv.Atoi(resp.Project.ProjectNumber)
	return projNumber
}

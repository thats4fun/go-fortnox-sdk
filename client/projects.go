package client

import (
	"context"
	"fmt"
)

const (
	projectsURI = "projects"
)

// GetProject does _GET https://api.fortnox.se/3/projects/{ProjectNumber}
//
// projectNumber - identifies the project
func (c *Client) GetProject(ctx context.Context, projectNumber int) (*Project, error) {
	resp := &GetProjectResp{}

	uri := fmt.Sprintf("%s/%d", projectsURI, projectNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Project, nil
}

// UpdateProject does _PUT https://api.fortnox.se/3/projects/{ProjectNumber}
//
// projectNumber - identifies the project
//
// p - project to update
func (c *Client) UpdateProject(ctx context.Context, projectNumber int, p *Project) (*Project, error) {
	req := &UpdateProjectReq{Project: *p}
	resp := &UpdateProjectResp{}

	uri := fmt.Sprintf("%s/%d", projectsURI, projectNumber)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Project, nil
}

// RemoveProject does _DELETE https://api.fortnox.se/3/projects/{ProjectNumber}
//
// projectNumber - identifies the project
func (c *Client) RemoveProject(ctx context.Context, projectNumber int) error {
	uri := fmt.Sprintf("%s/%d", projectsURI, projectNumber)
	return c._DELETE(ctx, uri)
}

// GetAllProjects does _GET https://api.fortnox.se/3/projects
func (c *Client) GetAllProjects(ctx context.Context) ([]Project, error) {
	resp := &GetAllProjectsResp{}

	err := c._GET(ctx, projectsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.Projects, nil
}

// CreateProject does POST https://api.fortnox.se/3/projects/
//
// project - project to create
func (c *Client) CreateProject(ctx context.Context, p *Project) (*Project, error) {
	req := &CreateProjectReq{Project: *p}
	resp := &CreateProjectResp{}

	err := c._POST(ctx, projectsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Project, nil
}

type Project struct {
	Url           string `json:"@url,omitempty"`
	Comments      string `json:"Comments,omitempty,omitempty"`
	ContactPerson string `json:"ContactPerson,omitempty,omitempty"`
	Description   string `json:"Description,omitempty"`
	EndDate       string `json:"EndDate,omitempty"`
	ProjectLeader string `json:"ProjectLeader,omitempty"`
	ProjectNumber string `json:"ProjectNumber,omitempty"`
	Status        string `json:"Status,omitempty"`
	StartDate     string `json:"StartDate,omitempty"`
}

type GetProjectResp struct {
	Project Project `json:"Project"`
}

type UpdateProjectReq struct {
	Project Project `json:"Project"`
}

type UpdateProjectResp struct {
	Project Project `json:"Project"`
}

type GetAllProjectsResp struct {
	Projects []Project `json:"Projects"`
}

type CreateProjectReq struct {
	Project Project `json:"Project"`
}

type CreateProjectResp struct {
	Project Project `json:"Project"`
}

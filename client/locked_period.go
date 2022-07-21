package client

import "context"

const (
	lockedPeriodURI = "lockedperiod"
)

// GetLockedPeriod does _GET
//
// If no date is returned, no period is locked.
func (c *Client) GetLockedPeriod(ctx context.Context) (*GetLockedPeriodResp, error) {
	resp := &GetLockedPeriodResp{}

	err := c._GET(ctx, lockedPeriodURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetLockedPeriodResp struct {
	LockedPeriod struct {
		EndDate string `json:"EndDate"`
	} `json:"LockedPeriod"`
}

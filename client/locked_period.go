package client

import "context"

const (
	lockedPeriodURI = "settings/lockedperiod/"
)

// GetLockedPeriod does _GET https://api.fortnox.se/3/settings/lockedperiod/
//
// If no date is returned, no period is locked.
func (c *Client) GetLockedPeriod(ctx context.Context) (*LockedPeriod, error) {
	resp := &GetLockedPeriodResp{}

	err := c._GET(ctx, lockedPeriodURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.LockedPeriod, nil
}

type LockedPeriod struct {
	EndDate string `json:"EndDate"`
}

type GetLockedPeriodResp struct {
	LockedPeriod LockedPeriod `json:"LockedPeriod"`
}

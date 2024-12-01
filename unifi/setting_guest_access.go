package unifi

import "context"

func (c *Client) GetSettingGuestAccess(ctx context.Context, site string) (*SettingGuestAccess, error) {
	return c.getSettingGuestAccess(ctx, site)
}

func (c *Client) UpdateSettingGuestAccess(ctx context.Context, site string, s *SettingGuestAccess) (*SettingGuestAccess, error) {
	s.Key = "mgmt"
	return c.updateSettingGuestAccess(ctx, site, s)
}

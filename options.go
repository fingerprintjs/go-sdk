package fingerprint

type ConfigOption func(*Client)

func WithRegion(region Region) ConfigOption {
	return func(c *Client) {
		c.region = region
	}
}

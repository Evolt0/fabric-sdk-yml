package base

type Option func(*Client)

func WithInit() Option {
	return func(c *Client) {
		c.Init()
	}
}

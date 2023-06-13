package poolV2

import "github.com/YuzuWiki/Pixivlee"

// Client http.request, do network
type Client struct{}

func (c *Client) SetHeaders(headers ...Pixivlee.THeader) error {
	return nil
}

func (c *Client) SetCookies(host string, cookies ...Pixivlee.TCookie) error {
	return nil
}

// Pool Client pool, manager http.request (client obj)
type Pool struct {
}

// acquire  If pool have an idle request, will return
func (p *Pool) acquire(ctx Pixivlee.IContext) (error, *Client) {
	return nil, nil
}

// acquire  release request
func (p *Pool) release(c *Client) error {
	return nil
}

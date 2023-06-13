package poolV2

import (
	"errors"

	"github.com/YuzuWiki/Pixivlee"
)

// Client http.request, do network
type Client struct {
}

func (c *Client) SetHeaders(headers ...Pixivlee.THeader) error {
	return nil
}

func (c *Client) SetCookies(host string, cookies ...Pixivlee.TCookie) error {
	return nil
}

func newClient(ctx Pixivlee.TContext) (error, *Client) {
	return nil, nil
}

// Pool Client pool, manager http.request (client obj)
type Pool struct {
	pool map[string]*Client
}

func (p *Pool) Client(ctx Pixivlee.TContext) (err error, c *Client) {
	if !ctx.Pixiver.IsActive() {
		return errors.New("invalid Cookie"), nil
	}

	c, isOk := p.pool[ctx.Pixiver.Pid()]
	if !isOk {
		err, c = newClient(ctx)
		if err != nil {
			return err, nil
		}

		p.pool[ctx.Pixiver.Pid()] = c
	}
	return nil, c
}

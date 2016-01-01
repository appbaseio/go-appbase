package appbase

import (
	"github.com/sacheendra/go-appbase/actions"
	"github.com/sacheendra/go-appbase/connection"
)

type Client struct {
	conn *connection.Connection
}

func NewClient(URL string, username string, password string, appname string) (*Client, error) {
	conn, err := connection.NewConnection(URL, username, password, appname)
	if err != nil {
		return nil, err
	}

	client := &Client{
		conn: conn,
	}

	return client, nil
}

func (c *Client) Ping() error {
	return c.conn.Ping()
}

func (c *Client) Index() *actions.IndexService {
	return actions.NewIndexService(c.conn)
}

func (c *Client) Update() *actions.UpdateService {
	return actions.NewUpdateService(c.conn)
}

func (c *Client) Delete() *actions.DeleteService {
	return actions.NewDeleteService(c.conn)
}

func (c *Client) Get() *actions.GetService {
	return actions.NewGetService(c.conn)
}

func (c *Client) GetStream() *actions.GetStreamService {
	return actions.NewGetStreamService(c.conn)
}

func (c *Client) Search() *actions.SearchService {
	return actions.NewSearchService(c.conn)
}

func (c *Client) SearchStream() *actions.SearchStreamService {
	return actions.NewSearchStreamService(c.conn)
}

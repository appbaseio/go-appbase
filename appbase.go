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

// Index returns an IndexService which is used to index a document
func (c *Client) Index() *actions.IndexService {
	return actions.NewIndexService(c.conn)
}

// Update returns an UpdateService which is used to update a document
func (c *Client) Update() *actions.UpdateService {
	return actions.NewUpdateService(c.conn)
}

// Delete returns a DeleteService which is used to delete a document
func (c *Client) Delete() *actions.DeleteService {
	return actions.NewDeleteService(c.conn)
}

// Get returns a GetService which is used to retrieve a document
func (c *Client) Get() *actions.GetService {
	return actions.NewGetService(c.conn)
}

// GetStream is used to start a stream of updates corresponding to a document
func (c *Client) GetStream() *actions.GetStreamService {
	return actions.NewGetStreamService(c.conn)
}

// Search provides access to Elasticsearch's search functionality
func (c *Client) Search() *actions.SearchService {
	return actions.NewSearchService(c.conn)
}

// SearchStream is used to get updates corresponding to a query
func (c *Client) SearchStream() *actions.SearchStreamService {
	return actions.NewSearchStreamService(c.conn)
}

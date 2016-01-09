package actions

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/appbaseio/go-appbase/connection"
)

type GetResponse struct {
	Index   string                 `json:"_index"`   // index meta field
	Type    string                 `json:"_type"`    // type meta field
	Id      string                 `json:"_id"`      // id meta field
	Parent  string                 `json:"_parent"`  // parent meta field
	Version int64                  `json:"_version"` // version number, when Version is set to true in SearchService
	Source  *json.RawMessage       `json:"_source,omitempty"`
	Found   bool                   `json:"found,omitempty"`
	Fields  map[string]interface{} `json:"fields,omitempty"`
}

type GetServiceOptions struct {
	Type   string `validate:"required"`
	Id     string `validate:"required"`
	Params url.Values
}

type GetService struct {
	conn    *connection.Connection
	options *GetServiceOptions
}

func NewGetService(conn *connection.Connection) *GetService {
	return &GetService{
		conn:    conn,
		options: &GetServiceOptions{},
	}
}

func (g *GetService) Type(_type string) *GetService {
	g.options.Type = _type
	return g
}

func (g *GetService) Id(_id string) *GetService {
	g.options.Id = _id
	return g
}

func (g *GetService) URLParams(params url.Values) *GetService {
	g.options.Params = params
	return g
}

func (g *GetService) Do() (*GetResponse, error) {
	err := validate(g.options)
	if err != nil {
		return nil, err
	}

	responseDecoder, err := g.conn.PerformRequest("GET", strings.Join([]string{g.options.Type, g.options.Id}, "/"), g.options.Params, "")
	if err != nil {
		return nil, err
	}

	getResponse := &GetResponse{}
	err = responseDecoder.Decode(getResponse)
	if err != nil {
		return nil, err
	}

	return getResponse, nil
}

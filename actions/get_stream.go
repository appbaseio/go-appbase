package actions

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/sacheendra/go-appbase/connection"
)

type GetStreamResponse struct {
	responseDecoder *json.Decoder
}

func (g *GetStreamResponse) Next() (getResponse *GetResponse, err error) {
	getResponse = &GetResponse{}
	err = g.responseDecoder.Decode(getResponse)
	if err != nil {
		return nil, err
	}

	return getResponse, nil
}

type GetStreamService struct {
	GetService
}

func NewGetStreamService(conn *connection.Connection) *GetStreamService {
	return &GetStreamService{
		GetService{
			conn:    conn,
			options: &GetServiceOptions{},
		},
	}
}

func (g *GetStreamService) Type(_type string) *GetStreamService {
	g.options.Type = _type
	return g
}

func (g *GetStreamService) Id(_id string) *GetStreamService {
	g.options.Id = _id
	return g
}

func (g *GetStreamService) URLParams(params url.Values) *GetStreamService {
	g.options.Params = params
	return g
}

func (g *GetStreamService) Do() (*GetStreamResponse, error) {
	err := validate(g.options)
	if err != nil {
		return nil, err
	}

	if g.options.Params == nil {
		g.options.Params = make(url.Values)
	}

	g.options.Params.Del("stream")
	g.options.Params.Set("streamonly", "true")

	responseDecoder, err := g.conn.PerformRequest("GET", strings.Join([]string{g.options.Type, g.options.Id}, "/"), g.options.Params, "")
	if err != nil {
		return nil, err
	}

	getStreamResponse := &GetStreamResponse{responseDecoder}

	return getStreamResponse, nil
}

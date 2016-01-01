package actions

import (
	"net/url"
	"strings"

	"github.com/sacheendra/go-appbase/connection"
)

type IndexResponse struct {
	Index   string `json:"_index"`
	Type    string `json:"_type"`
	Id      string `json:"_id"`
	Version int    `json:"_version"`
	Created bool   `json:"created"`
}

type IndexServiceOptions struct {
	Type   string `validate:"required"`
	Id     string
	Body   string `validate:"required"`
	Params url.Values
}

type IndexService struct {
	conn    *connection.Connection
	options *IndexServiceOptions
}

func NewIndexService(conn *connection.Connection) *IndexService {
	return &IndexService{
		conn:    conn,
		options: &IndexServiceOptions{},
	}
}

func (i *IndexService) Type(_type string) *IndexService {
	i.options.Type = _type
	return i
}

func (i *IndexService) Id(_id string) *IndexService {
	i.options.Id = _id
	return i
}

func (i *IndexService) Body(body string) *IndexService {
	i.options.Body = body
	return i
}

func (i *IndexService) URLParams(params url.Values) *IndexService {
	i.options.Params = params
	return i
}

func (i *IndexService) Do() (*IndexResponse, error) {
	err := validate(i.options)
	if err != nil {
		return nil, err
	}

	responseDecoder, err := i.conn.PerformRequest("POST", strings.Join([]string{i.options.Type, i.options.Id}, "/"), i.options.Params, i.options.Body)
	if err != nil {
		return nil, err
	}

	indexResponse := &IndexResponse{}
	err = responseDecoder.Decode(indexResponse)
	if err != nil {
		return nil, err
	}

	return indexResponse, nil
}

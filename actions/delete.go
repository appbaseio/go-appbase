package actions

import (
	"net/url"
	"strings"

	"github.com/appbaseio/go-appbase/connection"
)

type DeleteResponse struct {
	Index   string `json:"_index"`   // index meta field
	Type    string `json:"_type"`    // type meta field
	Id      string `json:"_id"`      // id meta field
	Version int64  `json:"_version"` // version number, when Version is set to true in SearchService
	Found   bool   `json:"found,omitempty"`
}

type DeleteServiceOptions struct {
	Type   string `validate:"required"`
	Id     string `validate:"required"`
	Params url.Values
}

type DeleteService struct {
	conn    *connection.Connection
	options *DeleteServiceOptions
}

func NewDeleteService(conn *connection.Connection) *DeleteService {
	return &DeleteService{
		conn:    conn,
		options: &DeleteServiceOptions{},
	}
}

func (d *DeleteService) Type(_type string) *DeleteService {
	d.options.Type = _type
	return d
}

func (d *DeleteService) Id(_id string) *DeleteService {
	d.options.Id = _id
	return d
}

func (d *DeleteService) URLParams(params url.Values) *DeleteService {
	d.options.Params = params
	return d
}

func (d *DeleteService) Do() (*DeleteResponse, error) {
	err := validate(d.options)
	if err != nil {
		return nil, err
	}

	responseDecoder, err := d.conn.PerformRequest("DELETE", strings.Join([]string{d.options.Type, d.options.Id}, "/"), d.options.Params, "")
	if err != nil {
		return nil, err
	}

	deleteResponse := &DeleteResponse{}
	err = responseDecoder.Decode(deleteResponse)
	if err != nil {
		return nil, err
	}

	return deleteResponse, nil
}

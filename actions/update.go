package actions

import (
	"net/url"
	"strings"

	"github.com/appbaseio/go-appbase/connection"
)

type UpdateResponse struct {
	Update      string       `json:"_index"`
	Type        string       `json:"_type"`
	Id          string       `json:"_id"`
	Version     int          `json:"_version"`
	Created     bool         `json:"created"`
	GetResponse *GetResponse `json:"get"`
}

type UpdateServiceOptions struct {
	Type   string `validate:"required"`
	Id     string `validate:"required"`
	Body   string `validate:"required"`
	Params url.Values
}

type UpdateService struct {
	conn    *connection.Connection
	options *UpdateServiceOptions
}

func NewUpdateService(conn *connection.Connection) *UpdateService {
	return &UpdateService{
		conn:    conn,
		options: &UpdateServiceOptions{},
	}
}

func (u *UpdateService) Type(_type string) *UpdateService {
	u.options.Type = _type
	return u
}

func (u *UpdateService) Id(_id string) *UpdateService {
	u.options.Id = _id
	return u
}

func (u *UpdateService) Body(body string) *UpdateService {
	u.options.Body = body
	return u
}

func (u *UpdateService) Pretty() *UpdateService {
	if u.options.Params != nil {
		u.options.Params.Set("pretty", "true")
	} else {
		params := url.Values{}
		params.Set("pretty", "true")
		u.options.Params = params
	}
	return u
}

func (u *UpdateService) URLParams(params url.Values) *UpdateService {
	if u.options.Params.Get("pretty") == "true" {
		u.options.Params = params
		u.options.Params.Set("pretty", "true")
	} else {
		u.options.Params = params
	}

	return u
}

func (u *UpdateService) Do() (*UpdateResponse, error) {
	err := validate(u.options)
	if err != nil {
		return nil, err
	}

	responseDecoder, err := u.conn.PerformRequest("POST", strings.Join([]string{u.options.Type, u.options.Id, "_update"}, "/"), u.options.Params, u.options.Body)
	if err != nil {
		return nil, err
	}

	updateResponse := &UpdateResponse{}
	err = responseDecoder.Decode(updateResponse)
	if err != nil {
		return nil, err
	}

	return updateResponse, nil
}

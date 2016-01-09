package actions

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/appbaseio/go-appbase/Godeps/_workspace/src/github.com/spaolacci/murmur3"
	"github.com/appbaseio/go-appbase/connection"
)

type SearchStreamToURLOptions struct {
	Type     string           `json:"type" validate:"required"`
	Webhooks []*Webhook       `json:"webhooks" validate:"required"`
	Query    *json.RawMessage `json:"query" validate:"required"`
	query    string
}

type SearchStreamToURLResponse struct {
	IndexResponse
	path string
	conn *connection.Connection
}

func (s *SearchStreamToURLResponse) Stop() (*DeleteResponse, error) {
	responseDecoder, err := s.conn.PerformRequest("DELETE", s.path, nil, "")
	if err != nil {
		return nil, err
	}

	var stopResponse DeleteResponse
	err = responseDecoder.Decode(&stopResponse)
	if err != nil {
		return nil, err
	}

	return &stopResponse, nil
}

type SearchStreamToURLService struct {
	conn    *connection.Connection
	options *SearchStreamToURLOptions
}

func NewSearchStreamToURLService(conn *connection.Connection) *SearchStreamToURLService {
	return &SearchStreamToURLService{
		conn:    conn,
		options: &SearchStreamToURLOptions{Webhooks: make([]*Webhook, 0)},
	}
}

func (s *SearchStreamToURLService) Type(_type string) *SearchStreamToURLService {
	s.options.Type = _type
	return s
}

func (s *SearchStreamToURLService) Query(query string) *SearchStreamToURLService {
	s.options.query = query
	return s
}

func (s *SearchStreamToURLService) AddWebhook(webhook *Webhook) *SearchStreamToURLService {
	s.options.Webhooks = append(s.options.Webhooks, webhook)
	return s
}

func (s *SearchStreamToURLService) Do() (*SearchStreamToURLResponse, error) {
	raw_query := new(json.RawMessage)
	err := raw_query.UnmarshalJSON([]byte(s.options.query))
	if err != nil {
		return nil, err
	}
	s.options.Query = raw_query

	err = validate(s.options)
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(s.options)
	if err != nil {
		return nil, err
	}

	h1, h2 := murmur3.Sum128([]byte(s.options.query))
	id := fmt.Sprintf("%x%x", h1, h2)

	path := strings.Join([]string{".percolator/webhooks", s.options.Type, id}, "-0-")

	responseDecoder, err := s.conn.PerformRequest("POST", path, nil, string(body))
	if err != nil {
		return nil, err
	}

	searchStreamToURLResponse := &SearchStreamToURLResponse{path: path, conn: s.conn}
	err = responseDecoder.Decode(searchStreamToURLResponse)
	if err != nil {
		return nil, err
	}

	return searchStreamToURLResponse, nil
}

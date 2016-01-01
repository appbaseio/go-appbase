package actions

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/sacheendra/go-appbase/connection"
)

type SearchStreamResponse struct {
	SearchResponse
	responseDecoder *json.Decoder
}

func (s *SearchStreamResponse) Next() (getResponse *GetResponse, err error) {
	getResponse = &GetResponse{}
	err = s.responseDecoder.Decode(getResponse)
	if err != nil {
		return nil, err
	}

	return getResponse, nil
}

type SearchStreamService struct {
	SearchService
}

func NewSearchStreamService(conn *connection.Connection) *SearchStreamService {
	return &SearchStreamService{
		SearchService{
			conn:    conn,
			options: &SearchServiceOptions{},
		},
	}
}

func (s *SearchStreamService) Type(_type string) *SearchStreamService {
	s.options.Type = _type
	return s
}

func (s *SearchStreamService) Body(body string) *SearchStreamService {
	s.options.Body = body
	return s
}

func (s *SearchStreamService) URLParams(params url.Values) *SearchStreamService {
	s.options.Params = params
	return s
}

func (s *SearchStreamService) Do() (*SearchStreamResponse, error) {
	err := validate(s.options)
	if err != nil {
		return nil, err
	}

	if s.options.Params == nil {
		s.options.Params = make(url.Values)
	}

	streamonly := s.options.Params.Get("streamonly") == "true"
	if !streamonly {
		s.options.Params.Set("stream", "true")
	}

	responseDecoder, err := s.conn.PerformRequest("POST", strings.Join([]string{s.options.Type, "_search"}, "/"), s.options.Params, s.options.Body)
	if err != nil {
		return nil, err
	}

	searchStreamResponse := &SearchStreamResponse{}
	if !streamonly {
		err = responseDecoder.Decode(searchStreamResponse)
		if err != nil {
			return nil, err
		}
	}
	searchStreamResponse.responseDecoder = responseDecoder

	return searchStreamResponse, nil
}

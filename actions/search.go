package actions

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/appbaseio/go-appbase/connection"
)

type SearchResponse struct {
	TookInMillis int64         `json:"took"`         // search time in milliseconds
	ScrollId     string        `json:"_scroll_id"`   // only used with Scroll and Scan operations
	Hits         *SearchHits   `json:"hits"`         // the actual search hits
	Suggest      SearchSuggest `json:"suggest"`      // results from suggesters
	Aggregations Aggregations  `json:"aggregations"` // results from aggregations
	TimedOut     bool          `json:"timed_out"`    // true if the search timed out
}

// SearchResponse specifies the list of search hits.
type SearchHits struct {
	TotalHits int64        `json:"total"`     // total number of hits found
	MaxScore  *float64     `json:"max_score"` // maximum score of all hits
	Hits      []*SearchHit `json:"hits"`      // the actual hits returned
}

// SearchHit is a single hit.
type SearchHit struct {
	Score          *float64                       `json:"_score"`          // computed score
	Index          string                         `json:"_index"`          // index name
	Type           string                         `json:"_type"`           // type meta field
	Id             string                         `json:"_id"`             // external or internal
	Uid            string                         `json:"_uid"`            // uid meta field (see MapperService.java for all meta fields)
	Timestamp      int64                          `json:"_timestamp"`      // timestamp meta field
	TTL            int64                          `json:"_ttl"`            // ttl meta field
	Routing        string                         `json:"_routing"`        // routing meta field
	Parent         string                         `json:"_parent"`         // parent meta field
	Version        *int64                         `json:"_version"`        // version number, when Version is set to true in SearchService
	Sort           []interface{}                  `json:"sort"`            // sort information
	Highlight      SearchHitHighlight             `json:"highlight"`       // highlighter information
	Source         *json.RawMessage               `json:"_source"`         // stored document source
	Fields         map[string]interface{}         `json:"fields"`          // returned fields
	Explanation    *SearchExplanation             `json:"_explanation"`    // explains how the score was computed
	MatchedQueries []string                       `json:"matched_queries"` // matched queries
	InnerHits      map[string]*SearchHitInnerHits `json:"inner_hits"`      // inner hits with ES >= 1.5.0
}
type SearchHitInnerHits struct {
	Hits *SearchHits `json:"hits"`
}

// SearchExplanation explains how the score for a hit was computed.
// See http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/search-request-explain.html.
type SearchExplanation struct {
	Value       float64             `json:"value"`             // e.s. 1.0
	Description string              `json:"description"`       // e.s. "boost" or "ConstantScore(*:*), product of:"
	Details     []SearchExplanation `json:"details,omitempty"` // recursive details
}

// Suggest
// SearchSuggest is a map of suggestions.
// See http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/search-suggesters.html.
type SearchSuggest map[string][]SearchSuggestion

// SearchSuggestion is a single search suggestion.
// See http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/search-suggesters.html.
type SearchSuggestion struct {
	Text    string                   `json:"text"`
	Offset  int                      `json:"offset"`
	Length  int                      `json:"length"`
	Options []SearchSuggestionOption `json:"options"`
}

// SearchSuggestionOption is an option of a SearchSuggestion.
// See http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/search-suggesters.html.
type SearchSuggestionOption struct {
	Text    string      `json:"text"`
	Score   float64     `json:"score"`
	Freq    int         `json:"freq"`
	Payload interface{} `json:"payload"`
}

// Aggregations (see search_aggs.go)

// Highlighting
// SearchHitHighlight is the highlight information of a search hit.
// See http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/search-request-highlightins.html
// for a general discussion of highlightins.
type SearchHitHighlight map[string][]string

type SearchServiceOptions struct {
	Type   []string `validate:"required"`
	Body   string   `validate:"required"`
	Params url.Values
}

type SearchService struct {
	conn    *connection.Connection
	options *SearchServiceOptions
}

func NewSearchService(conn *connection.Connection) *SearchService {
	return &SearchService{
		conn:    conn,
		options: &SearchServiceOptions{},
	}
}

func (s *SearchService) Type(_type string) *SearchService {
	s.options.Type = []string{_type}
	return s
}

func (s *SearchService) Types(_types []string) *SearchService {
	s.options.Type = _types
	return s
}

func (s *SearchService) Body(body string) *SearchService {
	s.options.Body = body
	return s
}

func (s *SearchService) Pretty() *SearchService {
	params := url.Values{}
	params.Set("pretty", "true")
	s.options.Params = params
	return s
}

func (s *SearchService) URLParams(params url.Values) *SearchService {
	s.options.Params = params
	return s
}

func (s *SearchService) Do() (*SearchResponse, error) {
	err := validate(s.options)
	if err != nil {
		return nil, err
	}

	responseDecoder, err := s.conn.PerformRequest("POST", strings.Join([]string{strings.Join(s.options.Type, ","), "_search"}, "/"), s.options.Params, s.options.Body)
	if err != nil {
		return nil, err
	}

	searchResponse := &SearchResponse{}
	err = responseDecoder.Decode(searchResponse)
	if err != nil {
		return nil, err
	}

	return searchResponse, nil
}

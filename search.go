package appbase

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// Search for documents in Elasticsearch.
type SearchService struct {
	client     *Client
	body       string
	pretty     bool
	searchType string
	types      []string
}

// NewSearchService creates a new service for searching in Elasticsearch.
// You typically do not create the service yourself manually, but access
// it via client.Search().
func newSearchService(client *Client) *SearchService {
	builder := &SearchService{
		client: client,
	}
	return builder
}

// Body allows the user to set the request body manually.
func (s *SearchService) Body(body string) *SearchService {
	s.body = body
	return s
}

// Type restricts the search for the given type.
func (s *SearchService) Type(typ string) *SearchService {
	if s.types == nil {
		s.types = []string{typ}
	} else {
		s.types = append(s.types, typ)
	}
	return s
}

// Types allows to restrict the search to a list of types.
func (s *SearchService) Types(types ...string) *SearchService {
	if s.types == nil {
		s.types = make([]string, len(types))
	}
	s.types = append(s.types, types...)
	return s
}

// Pretty enables the caller to indent the JSON output.
func (s *SearchService) Pretty(pretty bool) *SearchService {
	s.pretty = pretty
	return s
}

// SearchType sets the search operation type. Valid values are:
// "query_then_fetch", "query_and_fetch", "dfs_query_then_fetch",
// "dfs_query_and_fetch", "count", "scan".
// See http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/search-request-search-type.html#search-request-search-type
// for details.
func (s *SearchService) SearchType(searchType string) *SearchService {
	s.searchType = searchType
	return s
}

func (s *SearchService) Do() (initialResponse *json.RawMessage, responseStream chan *json.RawMessage, errorStream chan error, err error) {
	// Build url
	path := ""

	// Types part
	if len(s.types) > 0 {
		path += strings.Join(s.types, ",")
		path += "/_search"
	} else {
		path += "_search"
	}

	// Parameters
	params := make(url.Values)
	if s.pretty {
		params.Set("pretty", fmt.Sprintf("%v", s.pretty))
	}
	if s.searchType != "" {
		params.Set("search_type", s.searchType)
	}

	params.Add("stream", "true")

	return s.client.PerformStreamingRequest("POST", path, params, s.body)
}

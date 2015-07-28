package appbase

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type GetService struct {
	client                        *Client
	typ                           string
	id                            string
	preference                    string
	fields                        []string
	refresh                       *bool
	realtime                      *bool
	fsc                           *FetchSourceContext
	versionType                   string
	version                       *int64
	ignoreErrorsOnGeneratedFields *bool
}

func newGetService(client *Client) *GetService {
	builder := &GetService{
		client: client,
		typ:    "_all",
	}
	return builder
}

func (b *GetService) Type(typ string) *GetService {
	b.typ = typ
	return b
}

func (b *GetService) Id(id string) *GetService {
	b.id = id
	return b
}

func (b *GetService) Preference(preference string) *GetService {
	b.preference = preference
	return b
}

func (b *GetService) Fields(fields ...string) *GetService {
	if b.fields == nil {
		b.fields = make([]string, 0)
	}
	b.fields = append(b.fields, fields...)
	return b
}

func (s *GetService) FetchSource(fetchSource bool) *GetService {
	if s.fsc == nil {
		s.fsc = NewFetchSourceContext(fetchSource)
	} else {
		s.fsc.SetFetchSource(fetchSource)
	}
	return s
}

func (s *GetService) FetchSourceContext(fetchSourceContext *FetchSourceContext) *GetService {
	s.fsc = fetchSourceContext
	return s
}

func (b *GetService) Refresh(refresh bool) *GetService {
	b.refresh = &refresh
	return b
}

func (b *GetService) Realtime(realtime bool) *GetService {
	b.realtime = &realtime
	return b
}

func (b *GetService) VersionType(versionType string) *GetService {
	b.versionType = versionType
	return b
}

func (b *GetService) Version(version int64) *GetService {
	b.version = &version
	return b
}

func (b *GetService) IgnoreErrorsOnGeneratedFields(ignore bool) *GetService {
	b.ignoreErrorsOnGeneratedFields = &ignore
	return b
}

// Validate checks if the operation is valid.
func (s *GetService) Validate() error {
	var invalid []string
	if s.id == "" {
		invalid = append(invalid, "Id")
	}
	if s.typ == "" {
		invalid = append(invalid, "Type")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

func (b *GetService) Do() (initialResponse *json.RawMessage, responseStream chan *json.RawMessage, errorStream chan error, err error) {
	// Check pre-conditions
	if err := b.Validate(); err != nil {
		return nil, nil, nil, err
	}

	path := fmt.Sprintf("%s/%s", b.typ, b.id)

	params := make(url.Values)
	if b.realtime != nil {
		params.Add("realtime", fmt.Sprintf("%v", *b.realtime))
	}
	if len(b.fields) > 0 {
		params.Add("fields", strings.Join(b.fields, ","))
	}
	if b.preference != "" {
		params.Add("preference", b.preference)
	}
	if b.refresh != nil {
		params.Add("refresh", fmt.Sprintf("%v", *b.refresh))
	}
	if b.realtime != nil {
		params.Add("realtime", fmt.Sprintf("%v", *b.realtime))
	}
	if b.ignoreErrorsOnGeneratedFields != nil {
		params.Add("ignore_errors_on_generated_fields", fmt.Sprintf("%v", *b.ignoreErrorsOnGeneratedFields))
	}
	if len(b.fields) > 0 {
		params.Add("_fields", strings.Join(b.fields, ","))
	}
	if b.version != nil {
		params.Add("version", fmt.Sprintf("%d", *b.version))
	}
	if b.versionType != "" {
		params.Add("version_type", b.versionType)
	}
	if b.fsc != nil {
		for k, values := range b.fsc.Query() {
			params.Add(k, strings.Join(values, ","))
		}
	}

	params.Add("stream", "true")

	return b.client.PerformStreamingRequest("GET", path, params, "")
}

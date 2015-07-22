package appbase

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	client   *http.Client
	url      string
	username string
	password string
	appname  string
}

func NewClient(URL string, username string, password string, appname string) (*Client, error) {
	if URL[len(URL)-1:] == "/" {
		URL = URL[:len(URL)-1]
	}

	client := &Client{
		client:   &http.Client{},
		url:      URL,
		username: username,
		password: password,
		appname:  appname,
	}

	err := client.Ping()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) StreamDocument() *GetService {
	return newGetService(c)
}

func (c *Client) Ping() error {
	req, err := http.NewRequest("HEAD", c.url, nil)
	if err != nil {
		return err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New("Unable to connect to server")
	}

	return nil
}

func (c *Client) PerformStreamingRequest(method string, path string, params url.Values, body string) (initialResponse *json.RawMessage, responseStream chan *json.RawMessage, errorStream chan error, err error) {
	fullUrl := fmt.Sprintf("%s/%s/%s?%s", c.url, c.appname, path, params.Encode())

	req, err := http.NewRequest(method, fullUrl, strings.NewReader(body))
	if err != nil {
		return nil, nil, nil, err
	}
	req.SetBasicAuth(c.username, c.password)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, nil, nil, err
	}

	var initialResponseObject json.RawMessage
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&initialResponseObject)
	if err != nil {
		return nil, nil, nil, err
	}
	initialResponse = &initialResponseObject

	responseStream = make(chan *json.RawMessage, 1)
	errorStream = make(chan error, 1)

	go func() {
		for {
			defer func() {
				if e := recover(); e != nil {
					log.Println(e)
				}
			}()
			defer res.Body.Close()

			var event json.RawMessage
			err = dec.Decode(&event)
			if err != nil {
				errorStream <- err
				return
			} else {
				responseStream <- &event
			}
		}
	}()

	return initialResponse, responseStream, errorStream, nil
}

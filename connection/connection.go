package connection

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	appbase_error "github.com/appbaseio/go-appbase/error"
)

type Connection struct {
	client   *http.Client
	url      string
	username string
	password string
	appname  string
}

func NewConnection(URL string, username string, password string, appname string) (connection *Connection, err error) {
	if URL[len(URL)-1:] == "/" {
		URL = URL[:len(URL)-1]
	}

	connection = &Connection{
		client:   &http.Client{},
		url:      URL,
		username: username,
		password: password,
		appname:  appname,
	}

	err = connection.Ping()
	if err != nil {
		return nil, err
	}

	return connection, nil
}

func (c *Connection) Ping() error {
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

func (c *Connection) PerformRequest(method string, path string, params url.Values, body string) (responseDecoder *json.Decoder, err error) {
	fullUrl := fmt.Sprintf("%s/%s/%s?%s", c.url, c.appname, path, params.Encode())

	req, err := http.NewRequest(method, fullUrl, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.username, c.password)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.Close {
		res_body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		var appb_err appbase_error.AppbaseError
		err = json.Unmarshal(res_body, &appb_err)
		if err != nil {
			return nil, errors.New(fmt.Sprint(err.Error(), string(res_body)))
		}

		if appb_err.Status >= 300 {
			return nil, &appb_err
		}

		responseDecoder = json.NewDecoder(bytes.NewBuffer(res_body))
	} else {
		responseDecoder = json.NewDecoder(res.Body)
	}

	return responseDecoder, nil
}

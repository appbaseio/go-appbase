package appbase

import (
	"testing"
)

const URL string = "https://scalr.api.appbase.io"
const username string = "dW9DQYdot"
const password string = "40d5db8b-36c8-41ac-b6e9-d26d7e34ce1e"
const appname string = "testapp2"

const testtype string = "tweet"
const tweet1 string = `{"user":"sacheendra","message":"I am a robot."}`
const tweet2 string = `{"user":"sacheendra","message":"I am not a robot."}`
const query1 string = `{"query":{"match_all":{}}}`

var client *Client

func TestAppbase(t *testing.T) {
	var err error
	client, err = NewClient(URL, username, password, appname)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = client.Index().Type(testtype).Id("1").Body(tweet1).Do()
	if err != nil {
		t.Error(err)
		return
	}

	res, err := client.Delete().Type(testtype).Id("1").Do()
	if err != nil {
		t.Error(err)
		return
	}
	if !res.Found {
		t.Error("Document not found")
		return
	}
}

package appbase

import (
	"testing"
)

const URL string = "https://scalr.api.appbase.io"
const username string = "HnnFbzaRq"
const password string = "5d2ba0c3-4689-46f7-8cc9-60473479dc71"
const appname string = "go-appbase-tests"

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

	_, err = client.Index().Type(testtype).Id("2").Body(tweet1).Do()
	if err != nil {
		t.Error(err)
		return
	}

	res, err := client.Delete().Type(testtype).Id("2").Do()
	if err != nil {
		t.Error(err)
		return
	}
	if !res.Found {
		t.Error("Document not found")
		return
	}
}

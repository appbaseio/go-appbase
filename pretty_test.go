package appbase

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"testing"
)

/* defined in appbase_test.go
const URL string = "https://scalr.api.appbase.io"
const username string = "HnnFbzaRq"
const password string = "5d2ba0c3-4689-46f7-8cc9-60473479dc71"
const appname string = "go-appbase-tests"

const testtype string = "tweet"
const tweet1 string = `{"user":"sacheendra","message":"I am a robot."}`
const tweet2 string = `{"user":"sacheendra","message":"I am not a robot."}`
const query1 string = `{"query":{"match_all":{}}}`

var client *Client
*/
func TestPretty(t *testing.T) {

	var err error
	client, err = NewClient(URL, username, password, appname)
	if err != nil {
		t.Error(err)
		return
	}
	//Running test for Get()
	t.Log("Initialising test for Get() with Pretty()")
	responseGet, err := client.Get().Type(testtype).Id("1").Pretty().Do()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Initialising test for Get() without Pretty()")
	responseGet2, err := client.Get().Type(testtype).Id("1").Do()
	if err != nil {
		t.Error(err)
		return
	}
	value := compareResult(*responseGet.Source, *responseGet2.Source)
	if !value {
		t.Error("Pretty() does not return matching result")
	}
	t.Log("Get: Test Passed = ", value)

	//Running test for GetStream()
	t.Log("Initialising test for GetStream() with Pretty()")
	getStreamer, _ := client.GetStream().GetService.Type(testtype).Id("1").Pretty().Do()
	_, err = client.Index().Type(testtype).Id("2").Body(tweet2).Do()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Initialising test for GetStream() without Pretty()")
	getStreamer2, _ := client.GetStream().GetService.Type(testtype).Id("1").Do()
	_, err = client.Index().Type(testtype).Id("2").Body(tweet2).Do()
	if err != nil {
		t.Error(err)
		return
	}
	value = compareResult(*getStreamer.Source, *getStreamer2.Source)
	if !value {
		t.Error("Pretty() does not return matching result")
	}
	t.Log("GetStream: Test Passed = ", value)

	//Initialising test for Search()
	t.Log("Initialising test for Search() with Pretty()")
	searchResponse, err := client.Search().Type(testtype).Body(query1).Pretty().Do()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("Initialising test for Search() without Pretty()")
	searchResponse2, err := client.Search().Type(testtype).Body(query1).Do()
	if err != nil {
		t.Error(err)
		return
	}
	value = compareResult(*searchResponse.Hits.Hits[0].Source, *searchResponse2.Hits.Hits[0].Source)
	if !value {
		t.Error("Pretty() does not return matching result")
	}
	t.Log("Search: Test Passed = ", value)

	//Initialising test for SearchStream()
	t.Log("Initialising test for SearchStream() with Pretty()")
	searchStreamResponse, err := client.SearchStream().SearchService.Type(testtype).Body(query1).Pretty().Do()
	if err != nil {
		t.Error(err)
		return
	}
	_, err = client.Index().Type(testtype).Id("2").Body(tweet2).Do()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Initialising test for SearchStream() without Pretty()")
	searchStreamResponse2, err := client.SearchStream().SearchService.Type(testtype).Body(query1).Do()
	if err != nil {
		t.Error(err)
		return
	}
	_, err = client.Index().Type(testtype).Id("2").Body(tweet2).Do()
	if err != nil {
		t.Error(err)
		return
	}
	value = compareResult(*searchStreamResponse.Hits.Hits[0].Source, *searchStreamResponse2.Hits.Hits[0].Source)
	if !value {
		t.Error("Pretty() does not return matching result")
	}
	t.Log("SearchStream: Test Passed = ", value)

	//Initialising test for Update()
	t.Log("Initialising test for Update() with Pretty()")
	param := url.values{}
	param.Set("fields", "_source")

	updateResponse, _ := client.Update().Type(testtype).Id("3").Body(fmt.Sprintf(`{ "doc": %s }`, tweet2)).Pretty().URLParams(param).Do()

	t.Log("Initialising test for Update() without Pretty()")
	updateResponse2, _ := client.Update().Type(testtype).Id("3").Body(fmt.Sprintf(`{ "doc": %s }`, tweet2)).URLParams(param).Do()

	value = compareResult(*updateResponse.GetResponse.Source, *updateResponse2.GetResponse.Source)
	if !value {
		t.Error("Pretty() does not return matching result")
	}
	t.Log("SearchStream: Test Passed = ", value)
}
func compareResult(pretty json.RawMessage, nonPretty json.RawMessage) bool {
	prettyResponse := string(pretty)
	prettyResponse = strings.Replace(prettyResponse, " ", "", -1)
	nonPrettyResponse, _ := json.MarshalIndent(nonPretty, "  ", "  ")
	nonPrettyRes := string(nonPrettyResponse)
	nonPrettyRes = strings.Replace(nonPrettyRes, " ", "", -1)
	if prettyResponse != nonPrettyRes {
		return false
	}
	return true
}

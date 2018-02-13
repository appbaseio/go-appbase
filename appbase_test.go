package appbase

import (
	"fmt"
	"net/url"
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

	// Test Pretty()

	indexResponse, err := client.Index().Type(testtype).Id("1").Body(tweet1).Pretty().Do()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("Index() Response with Pretty():")
	fmt.Println(*indexResponse)

	/*
		No documents are returned, hence nothing gets pretty printed
		Output:
		Index() Response with Pretty():
		{go-appbase-tests tweet 1 1 true}
	*/

	responseGet, err := client.Get().Type(testtype).Id("1").Pretty().Do()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("Get() Response with Pretty()")
	fmt.Println(string(*responseGet.Source))
	/* Output:
	Get() Response with Pretty()
	{
		"user" : "sacheendra",
		"message" : "I am a robot."
	}
	*/

	responseGet, err = client.Get().Type(testtype).Id("1").Do()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("Get() Response without Pretty()")
	fmt.Println(string(*responseGet.Source))
	/*
		Output:
		Get() Response without Pretty()
		{"user":"sacheendra","message":"I am a robot."}
	*/

	getStreamer, _ := client.GetStream().GetService.Type(testtype).Id("1").Pretty().Do()
	_, err = client.Index().Type(testtype).Id("2").Body(tweet2).Do()
	if err != nil {
		t.Error(err)
		return
	}
	getStreamResponse := *getStreamer.Source
	fmt.Println("GetStream() Response with Pretty()")
	fmt.Println(string(getStreamResponse))

	/*
		Output:
		GetStream() Response with Pretty()
		{
			"user" : "sacheendra",
			"message" : "I am a robot"
		}
	*/

	searchResponse, err := client.Search().Type(testtype).Body(query1).Pretty().Do()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("Search() Response with Pretty()")
	fmt.Println(string(*searchResponse.Hits.Hits[0].Source))

	/*
		Output:
		Search() Response with Pretty()
		{
			"user" : "sacheendra",
			"message" : "I am a robot"
		}
	*/

	searchResponse, err = client.Search().Type(testtype).Body(query1).Do()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("Search() Response without Pretty()")
	fmt.Println(string(*searchResponse.Hits.Hits[0].Source))

	/*
		Output:
		Search() Response without Pretty()
		{"user" : "sacheendra","message" : "I am a robot"}
	*/

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
	docResponse := *searchStreamResponse.Hits.Hits[0].Source
	fmt.Println("SearchStream() Response with Pretty()")
	fmt.Println(string(docResponse))

	/*
		Output:
		SearchStream() Response with Pretty()
		{
			"user" : "sacheendra",
			"message" : "I am not a robot"
		}
	*/

	param := url.Values{}
	param.Set("fields", "_source")

	updateResponse, _ := client.Update().Type(testtype).Id("3").Body(fmt.Sprintf(`{ "doc": %s }`, tweet2)).Pretty().URLParams(param).Do()
	fmt.Println("Update() Response with Pretty()")
	fmt.Println(string(*updateResponse.GetResponse.Source))

	/*
		Output:
		Update() Response with Pretty()
		{
			"user" : "sacheendra",
			"message" : "I am not a robot"
		}
	*/

	deleteResponse, err := client.Delete().Type(testtype).Id("2").Pretty().Do()
	if err != nil {
		t.Error(err)
		return
	}
	if !deleteResponse.Found {
		t.Error("Document not found")
		return
	}
	fmt.Println("Delete() Response with Pretty()")
	fmt.Println(*deleteResponse)

	/*
		No Documents are returned hence nothing gets pretty printed.
		Output:
		Delete() Response with Pretty()
		{go-appbase-tests tweet 2 7 true}
	*/

}

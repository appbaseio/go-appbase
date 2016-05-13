package examples

import (
	"fmt"
	"github.com/appbaseio/go-appbase"
	"log"
)

var client *appbase.Client

func Example() {

	client, _ = appbase.NewClient("https://scalr.api.appbase.io", "mj8IvN7DY", "c01fd88a-250e-4321-85bf-51574d5141dc", "go-test")

	err := client.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Client created")
	// Output:
	// Client created
}

func ExampleClient_Index() {
	const testtype string = "tweet"
	const tweet string = `{"user":"jack","message":"just setting up my twttr"}`
	res, err := client.Index().Type(testtype).Id("3").Body(tweet).Do()
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println(res.Id)
	// Output:
	// 3
}

func ExampleClient_Update() {
	const testtype string = "tweet"
	const altTweet string = `{"user":"jack","message":"or just forget it, kate's yoga classes are more fun"}`
	res, err := client.Update().Type(testtype).Id("3").Body(fmt.Sprintf(`{ "doc": %s }`, altTweet)).Do()
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println(res.Id)
	// Output:
	// 3
}

func ExampleClient_Search() {
	const testtype string = "tweet"
	const matchAllQuery string = `{"query":{"match_all":{}}}`
	res, err := client.Search().Type(testtype).Body(matchAllQuery).Do()
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println(res.Hits.TotalHits > 0)
	// Output:
	// true
}

func ExampleClient_SearchStream() {
	const testtype string = "tweet"
	const matchAllQuery string = `{"query":{"match_all":{}}}`
	const anotherTweet string = `{"user":"ev","message":"twttr's way better than odeo"}`
	searchStreamResponse, err := client.SearchStream().Type(testtype).Body(matchAllQuery).Do()
	if err != nil {
		log.Fatalln(err)
		return
	}

	_, err = client.Index().Type(testtype).Id("3").Body(anotherTweet).Do()
	if err != nil {
		log.Fatalln(err)
		return
	}

	docResponse, err := searchStreamResponse.Next()
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println(docResponse.Id)
	// Output:
	// 3
}

func ExampleClient_SearchStreamToURL() {
	const testtype string = "tweet"
	const matchAllQuery string = `{"query":{"match_all":{}}}`
	webhook := appbase.NewWebhook()
	webhook.URL = "https://www.mockbin.org/bin/cd6461ab-468f-42f5-865f-4eed22daae95"
	webhook.Method = "POST"
	webhook.Body = "hellowebhooks"
	searchStreamToURLResponse, err := client.SearchStreamToURL().Type(testtype).Query(matchAllQuery).AddWebhook(webhook).Do()
	if err != nil {
		log.Fatalln(err)
		return
	}

	stopSearchStream, err := searchStreamToURLResponse.Stop()
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println(searchStreamToURLResponse.Id == stopSearchStream.Id)
	// Output:
	// true
}

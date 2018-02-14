package appbase

import (
	"fmt"
	"log"
)

func Example() {

	// Importing the library as 'github.com/appbase/go-appbase'.
	// Then instantiate the client with appbase.NewClient(url, username, password, appname).
	client, _ = NewClient("https://scalr.api.appbase.io", "HnnFbzaRq", "5d2ba0c3-4689-46f7-8cc9-60473479dc71", "go-appbase-tests")

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
	// Similar to NewClient, we will instiate a webhook instance with appbase.NewWebhook()
	webhook := NewWebhook()
	// Webhook instancess need to have a URL, method and body (which can be string or a JSON object)
	webhook.URL = "https://www.mockbin.org/bin/cd6461ab-468f-42f5-865f-4eed22daae95"
	webhook.Method = "POST"
	webhook.Body = "hellowebhooks"
	const testtype string = "tweet"
	const matchAllQuery string = `{"query":{"match_all":{}}}`

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

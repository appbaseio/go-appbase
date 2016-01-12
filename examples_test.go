package appbase

import (
	"fmt"
	"log"
)

func Example() {
	client, _ := NewClient("https://scalr.api.appbase.io", "QEVrcElba", "5c13d943-a5d1-4b05-92f3-42707d49fcbb", "es2test1")

	err := client.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Client created")
	// Output:
	// Client created
}

func ExampleClient_Index() {
	res, err := client.Index().Type(testtype).Id("3").Body(tweet1).Do()
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println(res.Id)
	// Output:
	// 3
}

func ExampleClient_Update() {
	res, err := client.Update().Type(testtype).Id("3").Body(fmt.Sprintf(`{ "doc": %s }`, tweet2)).Do()
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println(res.Id)
	// Output:
	// 3
}

func ExampleClient_Search() {
	res, err := client.Search().Type(testtype).Body(query1).Do()
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println(res.Hits.TotalHits > 0)
	// Output:
	// true
}

func ExampleClient_SearchStream() {
	searchStreamResponse, err := client.SearchStream().Type(testtype).Body(query1).Do()
	if err != nil {
		log.Fatalln(err)
		return
	}

	_, err = client.Index().Type(testtype).Id("3").Body(tweet2).Do()
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
	webhook := NewWebhook()
	webhook.URL = "http://requestb.in/whm9cvwh"
	webhook.Method = "POST"
	searchStreamToURLResponse, err := client.SearchStreamToURL().Type(testtype).Query(query1).AddWebhook(webhook).Do()
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

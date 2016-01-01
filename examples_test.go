package appbase

import (
	"fmt"
	"log"
	"net/url"
)

func Example() {
	client, _ := NewClient("https://scalr.api.appbase.io", "dW9DQYdot", "40d5db8b-36c8-41ac-b6e9-d26d7e34ce1e", "testapp2")

	err := client.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Client created")
	// Output:
	// Client created
}

func ExampleClient_Index() {
	res, err := client.Index().Type(testtype).Id("1").Body(tweet1).Do()
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println(res.Id)
	// Output:
	// 1
}

func ExampleClient_Update() {
	res, err := client.Update().Type(testtype).Id("1").Body(fmt.Sprintf(`{ "doc": %s }`, tweet2)).Do()
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println(res.Id)
	// Output:
	// 1
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

	_, err = client.Index().Type(testtype).Id("1").Body(tweet2).Do()
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
	// 1
}

func ExampleClient_SearchStream_streamonly() {
	params := make(url.Values)
	params.Set("streamonly", "true")
	searchStreamResponse, err := client.SearchStream().Type(testtype).URLParams(params).Body(query1).Do()
	if err != nil {
		log.Fatalln(err)
		return
	}

	_, err = client.Index().Type(testtype).Id("1").Body(tweet2).Do()
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
	// 1
}

package appbase

import (
	"fmt"
	"testing"

	"github.com/olivere/elastic"
)

func TestStreamDocument(t *testing.T) {
	client, err := elastic.NewClient(elastic.SetURL(" http://Qvkaa4CYR:548440c2-828f-470d-a12e-6f8f7d3c5bee@scalr.api.appbase.io"), elastic.SetSniff(false))
	if err != nil {
		t.Error(err)
	}

	streamingClient, err := NewClient("http://scalr.api.appbase.io", "Qvkaa4CYR", "Qvkaa4CYR", "scalrtest")
	if err != nil {
		t.Error(err)
	}

	tweet1 := `{"user": "olivere", "message": "Welcome to Golang and Elasticsearch."}`
	_, err = client.Index().Index("testindex").Type("tweet").Id("1").BodyString(tweet1).Do()
	if err != nil {
		t.Fatal(err)
	}

	// Get document 1
	_, responseStream, _, err := streamingClient.StreamDocument().Type("tweet").Id("1").Do()
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.Index().Index("testindex").Type("tweet").Id("1").BodyString(tweet1).Do()
	if err != nil {
		t.Fatal(err)
	}

	for event := range responseStream {
		if event == nil {
			t.Errorf("Event not received")
		}
		break
	}
}

func ExampleStreamDocument(t *testing.T) {
	streamingClient, _ := NewClient("http://localhost:7999", "testuser", "testpass", "testindex")

	initialResponse, responseStream, _, _ := streamingClient.StreamDocument().Type("tweet").Id("1").Do()
	fmt.Println(initialResponse)

	for event := range responseStream {
		fmt.Println(event)
	}
}

package appbase

import (
	"testing"

	"github.com/olivere/elastic"
)

func TestStreamDocument(t *testing.T) {
	client, err := elastic.NewClient(elastic.SetURL("http://testuser:testpass@localhost:7999"), elastic.SetSniff(false))
	if err != nil {
		t.Error(err)
	}

	streamingClient, err := NewClient("http://localhost:7999", "testuser", "testpass", "testindex")
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

	event := <-responseStream

	if event == nil {
		t.Errorf("Event not received")
	}

	close(responseStream)
}

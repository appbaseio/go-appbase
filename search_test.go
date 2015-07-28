package appbase

import (
	"testing"
	"time"

	"github.com/olivere/elastic"
)

func TestStreamSearch(t *testing.T) {
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
		t.Error(err)
	}

	// Get document 1
	_, responseStream, _, err := streamingClient.StreamSearch().Type("tweet").Body(`
		{
			"query": {
				"match_all": {}
			}
		}
		`).Do()
	if err != nil {
		t.Error(err)
	}

	time.Sleep(time.Second * 2)

	_, err = client.Index().Index("testindex").Type("tweet").Id("1").BodyString(tweet1).Do()
	if err != nil {
		t.Error(err)
	}

	for event := range responseStream {
		if event == nil {
			t.Error("Event not received")
		}
		e, err := event.MarshalJSON()
		if err != nil {
			t.Error(err)
		}
		if string(e) != `{"_type":"tweet","_id":"1","_source":{"message":"Welcome to Golang and Elasticsearch.","user":"olivere"}}` {
			t.Error("event was not as expected: ", string(e))
		}
		break
	}
}

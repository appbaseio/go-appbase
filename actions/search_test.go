package actions

import (
	"testing"

	"github.com/sacheendra/go-appbase/connection"

	"github.com/sacheendra/go-appbase/Godeps/_workspace/src/gopkg.in/olivere/elastic.v3"
)

func SearchServiceTest(t *testing.T, client *elastic.Client, conn *connection.Connection) {
	_, err := NewIndexService(conn).Type(testtype).Id("1").Body(tweet1).Do()
	if err != nil {
		t.Error(err)
		return
	}

	res, err := NewSearchService(conn).Type(testtype).Body(query1).Do()
	if err != nil {
		t.Error(err)
		return
	}

	if res.Hits.TotalHits == 0 {
		t.Error("No hits found")
		return
	}
}

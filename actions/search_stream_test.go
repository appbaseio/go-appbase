package actions

import (
	"testing"

	"github.com/sacheendra/go-appbase/connection"

	"github.com/sacheendra/go-appbase/Godeps/_workspace/src/gopkg.in/olivere/elastic.v3"
)

func SearchStreamServiceTest(t *testing.T, client *elastic.Client, conn *connection.Connection) {
	_, err := client.Index().Index(appname).Type(testtype).Id("1").BodyString(tweet1).Do()
	if err != nil {
		t.Error(err)
		return
	}

	res, err := NewSearchStreamService(conn).Type(testtype).Body(query1).Do()
	if err != nil {
		t.Error(err)
		return
	}
	if res.Hits.TotalHits == 0 {
		t.Error("No hits found")
		return
	}

	_, err = client.Index().Index(appname).Type(testtype).Id("1").BodyString(tweet1).Do()
	if err != nil {
		t.Error(err)
		return
	}

	_, err = res.Next()
	if err != nil {
		t.Error(err)
		return
	}
}

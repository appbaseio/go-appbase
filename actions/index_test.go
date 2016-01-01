package actions

import (
	"testing"

	"github.com/sacheendra/go-appbase/connection"

	"github.com/sacheendra/go-appbase/Godeps/_workspace/src/gopkg.in/olivere/elastic.v3"
)

func IndexServiceTest(t *testing.T, client *elastic.Client, conn *connection.Connection) {
	found, err := client.Exists().Index(appname).Type(testtype).Id("1").Do()
	if err != nil {
		t.Error(err)
		return
	}
	if found {
		_, err = client.Delete().Index(appname).Type(testtype).Id("1").Do()
		if err != nil {
			t.Error(err)
			return
		}
	}

	_, err = NewIndexService(conn).Type(testtype).Id("1").Body(tweet1).Do()
	if err != nil {
		t.Error(err)
		return
	}

	res, err := NewGetService(conn).Type(testtype).Id("1").Do()
	if err != nil {
		t.Error(err)
		return
	}
	if !res.Found {
		t.Error("Document not found")
		return
	}
}

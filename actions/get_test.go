package actions

import (
	"testing"

	"github.com/sacheendra/go-appbase/connection"

	"github.com/sacheendra/go-appbase/Godeps/_workspace/src/gopkg.in/olivere/elastic.v3"
)

func GetServiceTest(t *testing.T, client *elastic.Client, conn *connection.Connection) {
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

	_, err = client.Index().Index(appname).Type(testtype).Id("1").BodyString(tweet1).Do()
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

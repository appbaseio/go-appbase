package actions

import (
	"testing"

	"github.com/sacheendra/go-appbase/connection"

	"github.com/sacheendra/go-appbase/Godeps/_workspace/src/gopkg.in/olivere/elastic.v3"
)

func DeleteServiceTest(t *testing.T, client *elastic.Client, conn *connection.Connection) {
	_, err := client.Index().Index(appname).Type(testtype).Id("1").BodyString(tweet1).Do()
	if err != nil {
		t.Error(err)
		return
	}

	res, err := NewDeleteService(conn).Type(testtype).Id("1").Do()
	if err != nil {
		t.Error(err)
		return
	}
	if !res.Found {
		t.Error("Document not found")
		return
	}
}

package actions

import (
	"testing"

	"github.com/appbaseio/go-appbase/connection"

	"github.com/appbaseio/go-appbase/Godeps/_workspace/src/gopkg.in/olivere/elastic.v3"
)

func GetStreamServiceTest(t *testing.T, client *elastic.Client, conn *connection.Connection) {
	_, err := client.Index().Index(appname).Type(testtype).Id("1").BodyString(tweet1).Do()
	if err != nil {
		t.Error(err)
		return
	}

	res, err := NewGetStreamService(conn).Type(testtype).Id("1").Do()
	if err != nil {
		t.Error(err)
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

package actions

import (
	"testing"

	"github.com/appbaseio/go-appbase/connection"

	"github.com/appbaseio/go-appbase/Godeps/_workspace/src/gopkg.in/olivere/elastic.v3"
)

func SearchStreamToURLServiceTest(t *testing.T, client *elastic.Client, conn *connection.Connection) {
	res, err := NewSearchStreamToURLService(conn).Type(testtype).Query(query1).AddWebhook(&Webhook{
		URL:    "http://requestb.in/whm9cvwh",
		Method: "POST",
	}).Do()
	if err != nil {
		t.Error(err)
		return
	}

	stp, err := res.Stop()
	if err != nil {
		t.Error(err)
		return
	}

	if res.Id != stp.Id {
		t.Error("Start and stop ids are not equal", res.Id, stp.Id)
		return
	}
}

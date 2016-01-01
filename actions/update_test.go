package actions

import (
	"fmt"
	"testing"

	"github.com/sacheendra/go-appbase/connection"

	"github.com/sacheendra/go-appbase/Godeps/_workspace/src/gopkg.in/olivere/elastic.v3"
)

func UpdateServiceTest(t *testing.T, client *elastic.Client, conn *connection.Connection) {
	_, err := NewIndexService(conn).Type(testtype).Id("1").Body(tweet1).Do()
	if err != nil {
		t.Error(err)
		return
	}

	_, err = NewUpdateService(conn).Type(testtype).Id("1").Body(fmt.Sprintf(`{ "doc": %s }`, tweet2)).Do()
	if err != nil {
		t.Error(err)
		return
	}

	res, err := NewGetService(conn).Type(testtype).Id("1").Do()
	if err != nil {
		t.Error(err)
		return
	}

	resSource, err := res.Source.MarshalJSON()
	if err != nil {
		t.Error(err)
		return
	}

	if !res.Found || string(resSource) != tweet2 {
		t.Error("Document does not match:", string(resSource), tweet2)
		return
	}
}

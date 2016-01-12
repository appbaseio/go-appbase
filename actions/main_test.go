package actions

import (
	"log"
	"testing"

	"github.com/appbaseio/go-appbase/connection"

	"github.com/appbaseio/go-appbase/Godeps/_workspace/src/gopkg.in/olivere/elastic.v3"
)

const URL string = "https://scalr.api.appbase.io"
const username string = "QEVrcElba"
const password string = "5c13d943-a5d1-4b05-92f3-42707d49fcbb"
const appname string = "es2test1"

const testtype string = "tweet"
const tweet1 string = `{"user":"sacheendra","message":"I am a robot."}`
const tweet2 string = `{"user":"sacheendra","message":"I am not a robot."}`
const query1 string = `{"query":{"match_all":{}}}`

func TestActions(t *testing.T) {
	var client *elastic.Client
	var conn *connection.Connection
	var err error

	client, err = elastic.NewClient(elastic.SetURL(URL), elastic.SetBasicAuth(username, password), elastic.SetSniff(false))
	if err != nil {
		log.Fatalln(err)
		return
	}

	conn, err = connection.NewConnection(URL, username, password, appname)
	if err != nil {
		log.Fatalln(err)
		return
	}

	GetServiceTest(t, client, conn)
	GetStreamServiceTest(t, client, conn)
	DeleteServiceTest(t, client, conn)
	IndexServiceTest(t, client, conn)
	UpdateServiceTest(t, client, conn)
	SearchServiceTest(t, client, conn)
	SearchStreamServiceTest(t, client, conn)
	SearchStreamToURLServiceTest(t, client, conn)
}

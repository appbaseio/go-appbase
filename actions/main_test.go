package actions

import (
	"log"
	"testing"

	"github.com/sacheendra/go-appbase/connection"

	"github.com/sacheendra/go-appbase/Godeps/_workspace/src/gopkg.in/olivere/elastic.v3"
)

const URL string = "https://scalr.api.appbase.io"
const username string = "dW9DQYdot"
const password string = "40d5db8b-36c8-41ac-b6e9-d26d7e34ce1e"
const appname string = "testapp2"

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
}

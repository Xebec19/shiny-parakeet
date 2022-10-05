package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Xebec19/shiny-parakeet/util"
	_ "github.com/lib/pq"
)

var TestQueries *Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	TestQueries = New(conn)

	os.Exit(m.Run())
}

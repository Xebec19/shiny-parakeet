package db

import (
	"database/sql"
	"log"

	db "github.com/Xebec19/shiny-parakeet/db/sqlc"
	"github.com/Xebec19/shiny-parakeet/util"
	_ "github.com/lib/pq"
)

var DBQuery *db.Queries

func Connect() {
	config, err := util.LoadConfig("../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect database:", err)
	}
	DBQuery = db.New(conn)
	log.Printf("Database connected successfully")
}

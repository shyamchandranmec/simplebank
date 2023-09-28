package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/shyamchandranmec/simplebank/api"
	db "github.com/shyamchandranmec/simplebank/db/sqlc"
	"github.com/shyamchandranmec/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Unable to load config")
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Can not conntect to testing database", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Can not start server", err.Error())
	}
}

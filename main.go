package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/transparentideas/gobank/api"
	db "github.com/transparentideas/gobank/db/sqlc"
	"github.com/transparentideas/gobank/util"
)

func main() {

	conf, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	connDB, err := sql.Open(conf.DBDriver, conf.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	store := db.NewStore(connDB)
	server, err := api.NewServer(conf, store)
	if err != nil {
		log.Fatal(err)
	}

	err = server.Start(conf.ServerAddress)

}

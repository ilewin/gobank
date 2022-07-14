package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/themeszone/gobank/api"
	db "github.com/themeszone/gobank/db/sqlc"
	"github.com/themeszone/gobank/util"
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
	server := api.NewServer(store)

	err = server.Start(conf.ServerAddress)

}

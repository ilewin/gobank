package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/transparentideas/gobank/util"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {

	conf, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal(err)
	}

	testDB, err = sql.Open(conf.DBDriver, conf.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}

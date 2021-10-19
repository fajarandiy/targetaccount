package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	"com.example/targetaccount/util"
	_ "github.com/go-sql-driver/mysql"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}
	// See "Important settings" section.
	testDB.SetConnMaxLifetime(time.Minute * time.Duration(config.ConnMaxLifetime))
	testDB.SetMaxOpenConns(config.MaxOpenConns)
	testDB.SetMaxIdleConns(config.MaxIdleConns)
	testQueries = New(testDB)

	os.Exit(m.Run())
}

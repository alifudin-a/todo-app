package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/alifudin-a/todo-app/util"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	util.LoadEnv()
	conn, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_SOURCE"))
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}
	testQueries = New(conn)

	os.Exit(m.Run())
}

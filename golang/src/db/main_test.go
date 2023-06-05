package db

import (
	"database/sql"
	"fmt"
	"testing"
)

var testQueries *Queries

var testDB *sql.DB

func TestMain(m *testing.M) {
	testDB = ConnectDB()

	testQueries := NewQueries(testDB)
	fmt.Println("testQueries", testQueries)

	m.Run()
}

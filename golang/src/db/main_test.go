package db

import (
	"fmt"
	"testing"
)

// var testQueries *Queries

// var testDB *sql.DB

func TestMain(m *testing.M) {
	// testDB = ConnectDB()

	// testQueries := NewQueries(testDB)
	fmt.Println("メインの処理があれば記述")

	m.Run()
}

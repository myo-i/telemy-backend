package db

import (
	"database/sql"
	"fmt"
	"os"
)

func ConnectDB() *sql.DB {
	var path string = fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"))

	db, err := sql.Open("mysql", path)
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}

func NewQueries(db *sql.DB) Queries {
	return Queries{
		connection: db,
	}
}

type Queries struct {
	connection *sql.DB
}

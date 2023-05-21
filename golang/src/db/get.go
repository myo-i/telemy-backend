package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Account struct {
	ID         int
	Nickname   string
	Email      string
	Password   string
	Created_at string
}

type Queries struct {
	connection *sql.DB
}

// TDDでDI使って書く！！！！！！！！！！！！！！！！！！！！！！！
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

func NewQueries(db *sql.DB) *Queries {
	return &Queries{
		connection: db,
	}
}

func (q *Queries) GetAccount(id string) Account {
	getAccount := "SELECT * FROM accounts WHERE id = ? LIMIT 1"

	row, err := q.connection.Query(getAccount, id)
	if err != nil {
		log.Fatalln("Db取得に失敗")
	}

	var a Account
	err = row.Scan(
		&a.ID,
		&a.Nickname,
		&a.Email,
		&a.Password,
		&a.Created_at,
	)
	return a
}

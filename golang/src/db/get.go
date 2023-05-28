package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

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

func NewQueries(db *sql.DB) Queries {
	return Queries{
		connection: db,
	}
}

// type CreateAccountParams struct {
// 	Nickname string
// 	Email    string
// 	Password string
// }

func (q Queries) CreateAccount(nickname, email, password string) {
	createAccount := "INSERT INTO accounts (nickname, email, password) VALUES (?, ?, ?);"
	result, err := q.connection.Exec(createAccount, nickname, email, password)
}

func (q Queries) GetAccount(id string) (Account, error) {
	var a Account
	getAccount := "SELECT * FROM accounts WHERE id = ? LIMIT 1"

	row, err := q.connection.Query(getAccount, id)
	if err != nil {
		return a, err
	}

	if row.Next() {
		err := row.Scan(
			&a.ID,
			&a.Nickname,
			&a.Email,
			&a.Password,
			&a.CreatedAt,
		)
		if err != nil {
			return a, err
		}
	}
	return a, nil
}

package db

import (
	_ "github.com/go-sql-driver/mysql"
)

type CreateAccountRequest struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q Queries) CreateAccount(r CreateAccountRequest) (int64, error) {
	createAccount := "INSERT INTO accounts (nickname, email, password) VALUES (?, ?, ?);"

	result, err := q.connection.Exec(createAccount, r.Nickname, r.Email, r.Password)
	if err != nil {
		return 0, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affected, nil
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

func (q Queries) DeleteAccount(id string) error {
	deleteAccount := "DELETE FROM accounts WHERE id = ?"

	_, err := q.connection.Exec(deleteAccount, id)
	if err != nil {
		return err
	}

	return nil
}

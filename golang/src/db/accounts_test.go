package db

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetAccount(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock error: %s", err)
	}
	defer db.Close()

	// テストデータの設定
	row := sqlmock.NewRows([]string{"id", "nickname", "email", "password", "created_at"}).
		AddRow(1, "tester", "test@test.com", "secret", "2020-01-01")

	// モックの動作設定
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM accounts WHERE id = ? LIMIT 1")).
		WithArgs("1").
		WillReturnRows(row)

	// テスト対象の実行
	queries := Queries{db}
	account, err := queries.GetAccount("1")
	if err != nil {
		t.Fatalf("Failed to GetAccount: %s", err)
	}

	// ログ出力
	t.Log("Account is ", account)

	// モックの検証
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestCreateAccount(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock error: %s", err)
	}
	defer db.Close()

	nickname := "test"
	email := "test@test.com"
	password := "pass"

	mock.ExpectExec("INSERT INTO accounts").
		WithArgs(nickname, email, password).
		WillReturnResult(sqlmock.NewResult(3, 1))

	queries := Queries{db}
	request := CreateAccountRequest{
		Nickname: nickname,
		Email:    email,
		Password: password,
	}

	if err = queries.CreateAccount(request); err != nil {
		t.Errorf("Failed to exec CreateAccount: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

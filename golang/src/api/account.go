package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"telemy/db"
)

func createAccount(w http.ResponseWriter) {
	str := "test"
	fmt.Fprintf(w, "create account: %s", str)
}

type CreateAccountRequest struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (server *Server) createAcount(w http.ResponseWriter, r http.ResponseWriter) {
	// jsonから受け取ったパラメータをcreateAccountRequestにバインドして値を取得し
	// db.CreateAccountに渡してクエリ実行
	resp := 

	// レスポンスからnickname, email, passwordを取得

	// DB追加
	err := server.queries.CreateAccount()

	// できれば挿入したデータをレスポンスとして返す

}

func (server *Server) getAccount(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/accounts/"):]

	// パラメータを同値のidを持つデータを取得
	// エラー処理は必ず修正
	account, err := server.queries.GetAccount(id)
	if err != nil {
		http.Error(w, "パラメータに問題あり", http.StatusBadRequest)
	}

	jsonData, err := json.Marshal(account)
	if err != nil {
		http.Error(w, "json形式への変換に失敗", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
	// fmt.Fprintf(w, "ID:%v, Nickname:%s, Email:%s, Password:%s, CreatedAt:%s", &account.ID, account.Nickname, account.Email, account.Password, account.CreatedAt)

}

func echoHello(w http.ResponseWriter, r *http.Request) {
	id := 1
	getAccount := "SELECT * FROM accounts WHERE id = ? LIMIT 1"

	dbConnection := db.ConnectDB()
	defer dbConnection.Close()

	row, err := dbConnection.Query(getAccount, id)
	if err != nil {
		log.Fatalln("Db取得に失敗")
	}

	var a db.Account
	if row.Next() {
		err = row.Scan(
			&a.ID,
			&a.Nickname,
			&a.Email,
			&a.Password,
			&a.CreatedAt,
		)
		if err != nil {
			log.Fatalln("データのスキャンに失敗")
		}
	}
	fmt.Fprintf(w, "Account: %s", a)

}

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

func (server *Server) createAccount(w http.ResponseWriter, r *http.Request) {
	var requestBody db.CreateAccountRequest

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "パラメータに問題あり", http.StatusBadRequest)
	}
	fmt.Println(requestBody.Nickname)

	err = server.queries.CreateAccount(requestBody)
	if err != nil {
		http.Error(w, "パラメータに問題あり", http.StatusBadRequest)
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		http.Error(w, "json形式への変換に失敗", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
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

}

func (server *Server) deleteAccount(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/delete-account/"):]

	err := server.queries.DeleteAccount(id)
	if err != nil {
		http.Error(w, "パラメータに問題あり", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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

func insertDemo(w http.ResponseWriter, r *http.Request) {
	dbConnection := db.ConnectDB()
	defer dbConnection.Close()

	createAccount := "INSERT INTO accounts (nickname, email, password) VALUES (?, ?, ?);"
	_, err := dbConnection.Exec(createAccount, "Bob", "test@insert.com", "demo")

	if err != nil {
		http.Error(w, "パラメータに問題あり", http.StatusBadRequest)
	}

}

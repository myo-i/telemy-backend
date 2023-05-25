package api

import (
	"fmt"
	"log"
	"net/http"
	"telemy/db"
)

func createAccount(w http.ResponseWriter) {
	str := "test"
	fmt.Fprintf(w, "create account: %s", str)
}

func (server *Server) getAccount(w http.ResponseWriter, r *http.Request) {
	// id := r.URL.Query().Get("id")
	id := "1"
	url := fmt.Sprintf("https://localhost/accounts/%d", id)

	resp, err := http.Get(url)
	if err != nil {
		if resp.StatusCode == http.StatusNotFound {
			http.Error(w, "指定したIDを持つデータが存在しません", http.StatusNotFound)
			return
		}
		http.Error(w, "リクエストの送信に失敗しました", http.StatusInternalServerError)
		return
	}
	account := server.queries.GetAccount(id)
	fmt.Fprintf(w, "ID:%v, Nickname:%s, Email:%s, Password:%s, Created_at:%s", &account.ID, account.Nickname, account.Email, account.Password, account.Created_at)

}

func echoHello(w http.ResponseWriter, r *http.Request) {
	dbConnection := db.ConnectDB()
	defer dbConnection.Close()
	fmt.Fprintf(w, "<h1>Hello World</h1>\n<h2>Hello World</h2>\n<h3>Hello World</h3>")

	rows, err := dbConnection.Query("SELECT * FROM accounts WHERE id = 1 LIMIT 1;")
	defer rows.Close()
	if err != nil {
		log.Fatalf("Failed to select: %s", err)
	}

	for rows.Next() {
		var account db.Account
		err := rows.Scan(&account.ID, &account.Nickname, &account.Email, &account.Password, &account.Created_at)

		if err != nil {
			log.Fatalf("Failed to get rows: %s", err)
		}
		fmt.Fprintf(w, "ID:%v, Nickname:%s, Email:%s, Password:%s, Created_at:%s", &account.ID, account.Nickname, account.Email, account.Password, account.Created_at)
	}
	// fmt.Fprintf(w, "DB resule: %s", a)

}

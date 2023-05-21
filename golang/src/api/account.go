package api

import (
	"fmt"
	"net/http"
)

func createAccount(w http.ResponseWriter) {
	str := "test"
	fmt.Fprintf(w, "create account: %s", str)
}

func (server *Server) getAccount(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	userID := 1
	url := fmt.Sprintf("https://localhost/accounts/%d", userID)

	account, err := server.querier.GetAccount(id)

	resp, err := http.Get(url)
	if err != nil {
		if resp.StatusCode == http.StatusNotFound {
			http.Error(w, "指定したIDを持つデータが存在しません", http.StatusNotFound)
			return
		}
		http.Error(w, "リクエストの送信に失敗しました", http.StatusInternalServerError)
		return
	}

}

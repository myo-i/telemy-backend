package api

import (
	"fmt"
	"log"
	"net/http"
	"telemy/db"
)

func NewServer() {
	http.HandleFunc("/", echoHello)
	http.ListenAndServe(":8080", nil)
}

type Article struct {
	Id    int
	Title string
	Body  string
}

func echoHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>\n<h2>Hello World</h2>\n<h3>Hello World</h3>")
	account(w)
	a := db.ConnectDB()
	defer a.Close()

	rows, err := a.Query("SELECT * FROM article;")
	defer rows.Close()
	if err != nil {
		log.Fatalf("Failed to select: %s", err)
	}

	for rows.Next() {
		var article Article
		err := rows.Scan(&article.Id, &article.Title, &article.Body)

		if err != nil {
			log.Fatalf("Failed to get rows: %s", err)
		}
		fmt.Fprintf(w, "ID:%v, Title:%s, Body:%s", article.Id, article.Title, article.Body)
	}
	fmt.Fprintf(w, "DB resule: %s", a)

}
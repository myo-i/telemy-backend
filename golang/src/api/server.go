package api

import (
	"net/http"
	"telemy/db"
)

type Server struct {
	queries db.Queries
}

func NewServer() {
	dbConnection := db.ConnectDB()
	defer dbConnection.Close()

	queries := db.NewQueries(dbConnection)

	server := Server{
		queries: queries,
	}

	server.setupRouter()
}

func (server *Server) setupRouter() {
	http.HandleFunc("/hello", echoHello)
	http.HandleFunc("/insertDemo", insertDemo)

	http.HandleFunc("/create-account", server.createAccount)
	http.HandleFunc("/accounts/", server.getAccount)
	http.HandleFunc("/delete-account/", server.deleteAccount)
	http.HandleFunc("/create-output/", server.createOutput)
	http.ListenAndServe(":8080", nil)

}

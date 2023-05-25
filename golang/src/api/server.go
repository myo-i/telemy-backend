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
	http.HandleFunc("/", echoHello)
	http.HandleFunc("/accounts", server.getAccount)
	http.ListenAndServe(":8080", nil)

}

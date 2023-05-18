package api

import (
	"fmt"
	"net/http"
)

func account(w http.ResponseWriter) {
	str := "test"
	fmt.Fprintf(w, "create account: %s", str)
}

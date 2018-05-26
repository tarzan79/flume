package main

import (
	"fmt"
	"net/http"

	"github.com/dimensi0n/flume/auth"
	"github.com/dimensi0n/flume/db"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	db.DbHandler(r)
	auth.AuthHandler(r)

	fmt.Println("Listening on port :8080")
	http.ListenAndServe(":8080", r)
}

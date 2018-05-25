package main

import (
	"net/http"

	"github.com/dimensi0n/flume/auth"
	"github.com/dimensi0n/flume/db"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	db.DbHandler(r)
	auth.AuthHandler(r)

	http.ListenAndServe(":8080", r)
}

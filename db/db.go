package db

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/gorilla/mux"
)

type Request struct {
	Content string `json:"content"`
}

type Database struct {
	gorm.Model
	Name    string `json:"name"`
	Content string `json:"content"`
}

func readDb(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req struct{ Name string }
	err := decoder.Decode(&req)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open("sqlite3", "./db.sqlite")

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Database{})

	var Db Database
	db.First(&Db, "name = ?", req.Name)

	encoder := json.NewEncoder(w)
	encoder.Encode(Db)

	db.Close()
}

func createDb(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req Database
	err := decoder.Decode(&req)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open("sqlite3", "./db.sqlite")

	if err != nil {
		panic(err)
	}

	var Db Database
	db.First(&Db, "name = ?", req.Name)

	if Db.Name != "" {
		json.NewEncoder(w).Encode("Db already exists")
		return
	}

	db.AutoMigrate(&Database{})

	db.Create(&Database{Name: req.Name, Content: req.Content})

	db.Close()

	encoder := json.NewEncoder(w)
	encoder.Encode(Database{Name: req.Name, Content: req.Content})
}

func writeDb(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req Database
	err := decoder.Decode(&req)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open("sqlite3", "./db.sqlite")

	if err != nil {
		panic(err)
	}

	var Db Database
	db.First(&Db, "name = ?", req.Name)

	db.Model(&Db).Update("Content", req.Content)
	db.Model(&Db).Update("Name", req.Name)

	encoder := json.NewEncoder(w)
	encoder.Encode(Db)

	db.Close()

}

func DbHandler(r *mux.Router) {
	r.HandleFunc("/db/read", readDb).Methods("GET")
	r.HandleFunc("/db/create", createDb).Methods("POST")
	r.HandleFunc("/db/write", writeDb).Methods("PUT")
}

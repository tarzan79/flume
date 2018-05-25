package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	IsConnected int
}

func signup(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req User
	err := decoder.Decode(&req)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	var U User
	db.First(&U, "username = ?", req.Username)

	if U.Username != "" {
		json.NewEncoder(w).Encode("User already exists")
		return
	}

	db.Create(&User{Username: req.Username, Password: req.Password, Email: req.Email, IsConnected: 1})

	db.Close()

	encoder := json.NewEncoder(w)
	encoder.Encode(User{Username: req.Username, Password: req.Password, Email: req.Email})
}

func login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req User
	err := decoder.Decode(&req)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	var U User
	db.First(&U, "username = ? AND password = ?", req.Username, req.Password)
	if U.Username != "" {
		encoder := json.NewEncoder(w)
		db.Model(&U).Update("IsConnected", 1)
		encoder.Encode("Logged in")
	} else {
		json.NewEncoder(w).Encode("User doesn't exist")
	}

	db.Close()
}

func disconnect(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req User
	err := decoder.Decode(&req)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	var U User
	db.First(&U, "username = ? AND password = ?", req.Username, req.Password)
	if U.Username != "" {
		encoder := json.NewEncoder(w)
		db.Model(&U).Update("IsConnected", 0)
		encoder.Encode("Disconnected")
	} else {
		json.NewEncoder(w).Encode("User doesn't exist")
	}

	db.Close()
}

func AuthHandler(r *mux.Router) {
	r.HandleFunc("/auth/signup", signup).Methods("POST")
	r.HandleFunc("/auth/login", login).Methods("GET")
	r.HandleFunc("/auth/disconnect", disconnect).Methods("GET")
}

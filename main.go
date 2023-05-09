package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/osama-yousef1993/go-training/auth"
	"github.com/osama-yousef1993/go-training/store"
)

func main() {
	var authMiddle auth.AuthMiddleware

	r := mux.NewRouter()

	r.HandleFunc("/generate-token", GetToken).Methods(http.MethodGet)
	r.Handle("/get-user", authMiddle.IsAuthorized(http.HandlerFunc(GetUser))).Methods(http.MethodGet)
	r.Handle("/get-tables", authMiddle.IsAuthorized(http.HandlerFunc(ReadTables))).Methods(http.MethodGet)

	http.ListenAndServe(":8080", r)
}

func GetToken(w http.ResponseWriter, r *http.Request) {
	var token auth.Generate
	stringToken, err := auth.GenerateToken("test@hotmail.com")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token.Token = stringToken

	res, err := json.Marshal(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := make(map[string]string)

	user["first_name"] = "test"
	user["Last_name"] = "test"
	user["email"] = "test@hotmail.com"

	res, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res)

}

func ReadTables(w http.ResponseWriter, r *http.Request) {
	tables, err := store.ReadTables()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(tables)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

// full api repo
// get post put  delete
// read api (read from database)
// build api (write database)
// user register
// api movies 
// create Tables
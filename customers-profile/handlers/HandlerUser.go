package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/custs-risk/connection"
	"github.com/custs-risk/structs"
	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, Raynaldo!")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)
	var custsrisk structs.User
	json.Unmarshal(payloads, &custsrisk)
	connection.DB.Create(&custsrisk)
	res := structs.Result{Code: 200, Data: custsrisk, Message: "Success create user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func ListUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	limit := vars["limit"]
	offset := vars["offset"]
	custsrisk_p := []structs.User{}

	connection.DB.
		Limit(limit).
		Offset(offset).
		Find(&custsrisk_p)

	res := structs.Result{Code: 200, Data: custsrisk_p, Message: "Success get user"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func DetailUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UserID := vars["id"]

	var custsrisk structs.User
	connection.DB.First(&custsrisk, UserID)
	res := structs.Result{Code: 200, Data: custsrisk, Message: "Success get user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

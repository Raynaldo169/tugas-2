package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleReq() {
	log.Println("Start development server localhost:9999")

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", HomePage)
	myRouter.HandleFunc("/user", CreateUser).Methods("OPTIONS", "POST")
	myRouter.HandleFunc("/users/{limit}/{offset}", ListUser).Methods("OPTIONS", "GET")
	myRouter.HandleFunc("/user/{id}", DetailUser).Methods("OPTIONS", "GET")

	log.Fatal(http.ListenAndServe(":8008", myRouter))
}

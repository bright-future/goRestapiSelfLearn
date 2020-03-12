package server

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"fmt"
	"../apis/v1"
)

func defaultResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func HandleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", defaultResponse)
	employeeV1 := myRouter.PathPrefix("/employee/v1.0.0").Subrouter()
	v1.HandleRequests(employeeV1)
	
    log.Fatal(http.ListenAndServe(":8080", myRouter))
}
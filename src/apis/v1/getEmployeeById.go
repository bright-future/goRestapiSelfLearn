package v1

import (
	"github.com/gorilla/mux"
	"../../models"
	"net/http"
	"fmt"
	"../../repository"
	"encoding/json"
)

func getEmployeeFromMongo(mp map[string]string) (models.Employee, error) {
	return repository.GetByID("mongodb", mp["id"])
}

func getEmployeeFromMysql(mp map[string]string) (models.Employee, error) {
	return repository.GetByID("mysql", mp["id"])
}

func GetEmployeeByIdHandler(w http.ResponseWriter, r *http.Request){
    // fmt.Fprintf(w, "Welcome to the get employee id handler!")
	// json.NewEncoder(w).Encode(Articles)
	vars := mux.Vars(r)
	fmt.Println(vars)
	var response models.Employee
	var err error
	switch vars["database"] {
		case "mongodb":
			fmt.Println(vars["database"])
			response, err = getEmployeeFromMongo(vars)
		case "mysql":
			response, err = getEmployeeFromMysql(vars)
		default :
			err = fmt.Errorf("No such requests")
		
	}
	if (err != nil) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "unable to serve");
		return
	}
	json.NewEncoder(w).Encode(response)
	return
}


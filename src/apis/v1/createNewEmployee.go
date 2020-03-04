
package v1

import (
	// "github.com/gorilla/mux"
	"../../models"
	"net/http"
	"fmt"
	"../../repository"
	"encoding/json"
	"math/rand"
	"strconv"
)

func CreateNewEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	var newEmployee models.NewEmployeeResponse
	json.NewDecoder(r.Body).Decode(&newEmployee)
	fmt.Println("r body", newEmployee)
	employee := models.Employee {
		EmployeeID: strconv.Itoa(rand.Int()),
		EmployeeName: newEmployee.EmployeeName,
	}
	res, err := repository.InsertEmployee(newEmployee.Source, employee)
	fmt.Println(res, err, employee)
	if (err != nil) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "unable to create");
		return
	}
	

	json.NewEncoder(w).Encode(employee)
	return
}
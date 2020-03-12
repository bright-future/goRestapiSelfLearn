package v1

import(
	"github.com/gorilla/mux"
)

func HandleRequests(handleV1 *mux.Router) {
	handleV1.HandleFunc("/getEmployeeById/{database}/{id}", GetEmployeeByIdHandler).Methods("GET")
	handleV1.HandleFunc("/createNewEmployee", CreateNewEmployeeHandler).Methods("POST")
	handleV1.HandleFunc("/deleteEmployee/{database}/{id}", DeleteEmployee).Methods("DELETE")
}
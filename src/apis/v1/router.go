package v1

import(
	"github.com/gorilla/mux"
)

func HandleRequests(handleV1 *mux.Router) {
	handleV1.HandleFunc("/getEmployeeById/{database}/{id}", GetEmployeeByIdHandler).Methods("GET")
}
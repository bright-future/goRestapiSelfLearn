package v1

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"../../repository"
)

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := repository.DeleteEmployee(vars["database"], vars["id"])
	if (err != nil) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "unable to serve")
		return
	}
	fmt.Fprintf(w, "done");
	return
}
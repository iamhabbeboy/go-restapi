package src

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func User(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	makeRequest(w, id)
}

func Users(w http.ResponseWriter, r *http.Request) {
	makeRequest(w)
}

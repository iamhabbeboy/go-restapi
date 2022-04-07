package src

import (
	"github.com/gorilla/mux"
)

func Route(route *mux.Router) {
	//route := *mux.NewRouter()
	route.HandleFunc("/v1", Users)
	route.HandleFunc("/v1/{id}", User)
}

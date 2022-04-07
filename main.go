package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"rest/api/src"
	"time"
)

func main() {
	// load dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error occured while loading env file")
	}

	// Setup Routes
	route := mux.NewRouter()
	src.Route(route)

	// Starting Server
	fmt.Print("Server starting @3000")
	server := &http.Server{
		Handler:      route,
		Addr:         "127.0.0.1:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())

}

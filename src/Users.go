package src

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func User(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	//w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	requestError := map[string]string{"error": "Id not in request"}
	if err != nil {
		err := json.NewEncoder(w).Encode(requestError)
		if err != nil {
			return
		}
	}
	var post Post
	response, err := makeRequest(id)
	json.Unmarshal(response, &post)
	//fmt.Println(string(response))
	w.Write([]byte(response))
	//err = json.NewEncoder(w).Encode(response)
	//if err != nil {
	//	return
	//}
}

func Users(w http.ResponseWriter, r *http.Request) {
	response, err := makeRequest(1)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	//_, err = w.Write([]byte(response))
	if err != nil {
		return
	}
}

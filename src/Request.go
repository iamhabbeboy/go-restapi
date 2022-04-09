package src

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func makeRequest(w http.ResponseWriter, id int) {
	apiUrl := os.Getenv("API_URL")
	getId := ""
	if id == 0 {
		getId = ""
	} else {
		getId = "/" + strconv.Itoa(id)
	}

	body, err := http.Get(apiUrl + "/posts" + getId)
	if err != nil {
		log.Fatalln("Error Occurred while getting posts")
	}

	defer body.Body.Close()
	res, err := ioutil.ReadAll(body.Body)

	if err != nil {
		panic(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

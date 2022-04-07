package src

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func makeRequest(id int) ([]byte, error) {
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

	response, err := ioutil.ReadAll(body.Body)
	defer body.Body.Close()

	return response, err
}

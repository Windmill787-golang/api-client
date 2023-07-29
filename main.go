package main

import (
	"io"
	"log"
	"net/http"
)

type CatImage struct {
	ID     string `json:"id"`
	Url    string `json:"url"`
	Width  string `json:"width"`
	Height string `json:"height"`
}

var api_url = "https://api.thecatapi.com/v1/images/search"

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//choose api with get, post methods and auth with token

	//create struct with fields of entity
	//create get and post funcs
	//get will returns list of json objects
	//get/:id will return single json object
	//post will return success or error

	//add auth func
	resp, err := http.Get(api_url)
	logError(err)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	logError(err)

	log.Println(string(body))
}

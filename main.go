package main

import (
	"log"

	"github.com/Windmill787-golang/api-client/cat"
)

func main() {
	//choose api with get, post methods and auth with token

	//create struct with fields of entity
	//create get and post funcs
	//get will returns list of json objects
	//get/:id will return single json object
	//post will return success or error

	//add auth func
	catImages, err := cat.GetCatImages(2)
	if err != nil {
		log.Fatal(err)
	}

	for _, catImage := range catImages {
		log.Println(catImage.Info())
	}
}

package main

import (
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"

	api_sample "github.com/luodynasty/saving-sloth/service/api/sample"
	src_sample "github.com/luodynasty/saving-sloth/service/src/sample"
)

func main () {
	// create instances of the required modules
	sampleModule := src_sample.NewModule()

	// create an instance of the sample API
	sampleAPI := api_sample.NewAPI(sampleModule)

	// create a new router
	router := httprouter.New()

	// register the API endpoint
	sampleAPI.Register(router)

	// start the server
	log.Println("Starting server on :8008")
	log.Fatal(http.ListenAndServe(":8008", router))
}
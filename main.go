package main

import (
	"log"

	api_sample "github.com/luodynasty/saving-sloth/service/api/sample"
	src_sample "github.com/luodynasty/saving-sloth/service/src/sample"
)

func main () {
	log.Println("testing")
}

type AllModules struct {
	SampleModule src_sample.SampleModule
	// ...
}

func init() {
	modules := initModules()
	// initAPIs()
}

// func initAPIs() {
// 	apiSample := api_sample.NewAPI()
// }

func initModules() *AllModules {
	sampleModule := src_sample.NewModule()

	return &AllModules{
		SampleModule: sampleModule,
	}
}
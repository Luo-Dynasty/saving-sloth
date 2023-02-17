package sample

import (
	"fmt"
	"log"
	"net/http"

	service_src "github.com/luodynasty/saving-sloth/service/src"
)

type API struct {
	SampleModule service_src.SampleInterface
}

func NewAPI(sampleModule service_src.SampleInterface) *API {
	return &API {
		SampleModule: sampleModule,
	}
}
func main() {
	http.HandleFunc("/testing", SampleHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (api API) SampleHandler(w http.ResponseWriter, r *http.Request) {
	api.SampleModule.PrintLogTesting()
}
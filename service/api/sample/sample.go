package sample

import (
	"net/http"
	"github.com/julienschmidt/httprouter"

	sample_src "github.com/luodynasty/saving-sloth/service/src/sample"
)

type API struct {
	SampleModule sample_src.SampleInterface
}

func NewAPI(sampleModule sample_src.SampleInterface) API {
	return API {
		SampleModule: sampleModule,
	}
}

func (api API) Register(router *httprouter.Router) {
	router.GET("/testing", api.SampleHandler)
}

func (api API) SampleHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	api.SampleModule.PrintLogTesting()
}
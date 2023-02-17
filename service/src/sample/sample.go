package sample

import (
	"log"
)

type SampleModule struct {

}

type SampleInterface interface {
	PrintLogTesting()
}

func NewModule() *SampleModule {
	return &SampleModule {

	}
}

func (sm SampleModule) PrintLogTesting() {
	log.Println("testing success")
}
package utils

import "fmt"

const (
	requestType      = "[UTILS]"
	StatusBuilding   = "[BUILDING]"
	StatusProcessing = "[PROCESSING]"
	StatusComplete   = "[COMPLETE]"
	StatusErr        = "[ERROR]"
)

var (
	Logger = logger{}
)

type logger struct{}

func (x logger) Log(requestType, requestStatus, requestMessage string) {
	fmt.Printf("%s%s %s\n", requestType, requestStatus, requestMessage)
}

func (x logger) LogErr(requestType string, err error) {
	x.Log(requestType, StatusErr, err.Error())
}

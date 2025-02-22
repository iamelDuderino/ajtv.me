package main

import (
	"encoding/json"
)

type applicationInterface struct{}

type apiResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (x applicationInterface) newResponse(ok bool, msg string) apiResponse {
	return apiResponse{
		OK:      ok,
		Message: msg,
	}
}

func (x apiResponse) JSON() string {
	b, err := json.Marshal(x)
	if err != nil {
		return `{'ok':false}`
	}
	return string(b)
}

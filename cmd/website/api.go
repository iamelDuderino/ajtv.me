package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type applicationInterface struct{}

func (x applicationInterface) newResponse(ok bool, msg string) apiResponse {
	return apiResponse{
		OK:      ok,
		Message: msg,
	}
}

func (x applicationInterface) get(fn func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method_not_allowed", http.StatusMethodNotAllowed)
			return
		}
		fn(w, r)
	}
}

type apiResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (x apiResponse) JSON() string {
	b, err := json.Marshal(x)
	if err != nil {
		fmt.Println(err)
		return `{'ok':false}`
	}
	return string(b)
}

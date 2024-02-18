package main

import (
	"encoding/json"
	"net/http"

	"github.com/iamelDuderino/my-website/internal/utils"
)

type applicationInterface struct{}

type apiResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (x applicationInterface) newResponse() apiResponse {
	return apiResponse{
		OK: true,
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

func (x apiResponse) JSON() string {
	b, err := json.Marshal(x)
	if err != nil {
		utils.Logger.LogErr(requestType, err)
		return `{'ok':false}`
	}
	return string(b)
}

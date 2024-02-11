package main

import (
	"encoding/json"

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

func (x apiResponse) JSON() string {
	b, err := json.Marshal(x)
	if err != nil {
		utils.Logger.LogErr(requestType, err)
		return `{'ok':false}`
	}
	return string(b)
}

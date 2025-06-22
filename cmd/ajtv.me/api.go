package main

import (
	"encoding/json"

	"github.com/iamelDuderino/my-website/internal/logger"
)

type applicationInterface struct {
	logger *logger.Logger
}

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

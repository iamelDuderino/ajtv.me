package api

import (
	"encoding/json"

	"github.com/iamelDuderino/my-website/internal/utils"
)

func (x API) newResponse() apiResponse {
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

package apexlegends

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type MapRotation struct {
	Current *Map `json:"current"`
	Next    *Map `json:"next"`
}

type Map struct {
	Start             int64  `json:"start,omitempty"`
	End               int64  `json:"end,omitempty"`
	Name              string `json:"map,omitempty"`
	DurationInSeconds int    `json:"DurationInSecs,omitempty"`
	DurationInMinutes int    `json:"DurationInMins,omitempty"`
	RemainingSeconds  int    `json:"remainingSecs,omitempty"`
	RemainingMinutes  int    `json:"remainingMins,omitempty"`
	RemainingTimer    string `json:"remainingTimer,omitempty"`
}

// Version 1 == BR Pubs, Version 2 == All Modes
func GetMapRotation() (*MapRotation, error) {
	uri := BaseURI + "/maprotation?auth=" + os.Getenv("APEX_LEGENDS_API_TOKEN")
	newReq, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	response, err := client.Do(newReq)
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	maps := &MapRotation{}
	err = json.Unmarshal(b, &maps)
	if err != nil {
		return nil, err
	}
	return maps, nil
}

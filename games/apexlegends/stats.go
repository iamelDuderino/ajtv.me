package apexlegends

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
)

type PlayerStats struct {
	Error         string        `json:"Error,omitempty"`
	GlobalStats   GlobalStats   `json:"global,omitempty"`
	RealtimeStats RealtimeStats `json:"realtime,omitempty"`
	TotalStats    TotalStats    `json:"total,omitempty"`
}

type GlobalStats struct {
	Name               string    `json:"name,omitempty"`
	UID                string    `json:"uid,omitempty"`
	Platform           string    `json:"platform,omitempty"`
	Level              int       `json:"level,omitempty"`
	ToNextLevelPercent int       `json:"toNextLevelPercent,omitempty"`
	Rank               RankStats `json:"rank,omitempty"`
	PrestigeLevel      int       `json:"levelPrestige,omitempty"`
}

type RealtimeStats struct {
	LobbyState     string `json:"lobbyState,omitempty"`
	IsOnline       int    `json:"isOnline,omitempty"`
	IsInGame       int    `json:"isInGame,omitempty"`
	PartyFull      int    `json:"partyFull,omitempty"`
	SelectedLegend string `json:"selectedLegend,omitempty"`
	CurrentState   string `json:"currentState,omitempty"`
}

type TotalStats struct{}

type RankStats struct {
	RankScore    int    `json:"rankScore,omitempty"`
	RankName     string `json:"rankName,omitempty"`
	RankDivision int    `json:"rankDiv,omitempty"`
}

func GetPlayerStats(Player, Platform string) (*PlayerStats, error) {
	uri := BaseURI + "/bridge/?auth=" + os.Getenv("APEX_LEGENDS_API_TOKEN") + "&player=" + Player + "&platform=" + strings.ToUpper(Platform)
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

	// fmt.Println(string(b)) // troubleshooting

	stats := &PlayerStats{}
	err = json.Unmarshal(b, &stats)
	if err != nil {
		return nil, err
	}
	return stats, nil
}

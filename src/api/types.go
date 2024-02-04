package api

type API struct{}

type apiResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

package api

import (
	"fmt"
	"net/http"
)

func (x API) HomeHandler(w http.ResponseWriter, r *http.Request) {
	resp := x.newResponse()
	resp.Message = "Thank you for testing my basic sample API!"
	fmt.Fprint(w, resp.JSON())
}

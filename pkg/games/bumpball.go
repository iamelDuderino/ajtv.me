package games

import (
	"fmt"
	"net/http"
)

func BumpBall(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprint(writer, "<h1>Future Page For Bump Ball</h1>")
}

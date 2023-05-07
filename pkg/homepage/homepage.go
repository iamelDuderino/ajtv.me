package homepage

import (
	"fmt"
	"net/http"
)

func Display(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprint(writer, "Hello, World! This is the Home Page.")
}

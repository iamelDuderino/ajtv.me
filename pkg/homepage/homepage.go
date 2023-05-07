package homepage

import (
	"fmt"
	"net/http"
)

func Display(writer http.ResponseWriter, req *http.Request) {
	homePage := "<h1>Andrew J Tomko V</h1><div>Future websume content</div>"
	fmt.Fprint(writer, homePage)
}

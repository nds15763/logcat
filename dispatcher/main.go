package dispatcher

import (
	"net/http"

	"github.com/tsuna/gohbase"
)

func main() {

	panic(http.ListenAndServe(":7002", nil))
}

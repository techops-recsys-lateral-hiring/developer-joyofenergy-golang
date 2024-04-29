package standard

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Healthcheck(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Fprintf(response, "Working!")
}

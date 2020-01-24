package function

import (
	"net/http"

	"github.com/alexellis/vend/vend/handlers"
)

func Handle(w http.ResponseWriter, r *http.Request) {

	handlers.Echo(w, r)
}

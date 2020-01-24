package handlers

import (
	"io/ioutil"
	"net/http"
)

func Echo(w http.ResponseWriter, r *http.Request) {

	if r.Body != nil {
		defer r.Body.Close()
		b, _ := ioutil.ReadAll(r.Body)
		w.Write(b)
	}

}

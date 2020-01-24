package function

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_handler(t *testing.T) {
	inputSt := "echo this"
	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBufferString(inputSt))
	Handle(rec, req)

	if rec.Body.String() != inputSt {
		t.Fail()
	}
	if rec.Body.Len() == 0 {
		t.Fail()
	}
}

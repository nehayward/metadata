package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearch(t *testing.T) {
	req, _ := http.NewRequest("GET", "/search?title=*", nil)
	rr := httptest.NewRecorder()

	router := NewRouter(AllRoutes())

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Wrong status")
	}
}

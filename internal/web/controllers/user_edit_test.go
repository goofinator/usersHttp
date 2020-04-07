package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/goofinator/usersHttp/internal/web/controllers"
)

func TestEditUserHandler(t *testing.T) {
	req, err := http.NewRequest("PUT", "/users/:1", nil)
	if err != nil {
		t.Fatalf("unexpected fail of NewRequest: %s", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		EditUserHandler(w, r, nil)
	})
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("unexpected status code:\nwant: %v\ngot: %v",
			http.StatusOK, status)
	}
}

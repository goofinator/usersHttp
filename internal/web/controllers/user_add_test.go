package controllers_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	. "github.com/goofinator/usersHttp/internal/web/controllers"
)

func TestAddUserHandler(t *testing.T) {
	var jsonStr = []byte(fmt.Sprintf(`{"Id": 0,
	"Name": "petya",
	"Lastname": "Pupkin",
	"Age": 22,
	"Birthdate": "%v"}`,
		time.Now().UTC().Format(time.RFC3339)))

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatalf("unexpected fail of NewRequest: %s", err)
	}
	//req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AddUserHandler(w, r, nil)
	})
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("unexpected status code:\nwant: %v\ngot: %v",
			http.StatusOK, status)
	}

	fmt.Println(rr.Body.String())
}

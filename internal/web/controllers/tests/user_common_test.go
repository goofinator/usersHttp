package controllers_test

import (
	"fmt"
	"time"

	"github.com/goofinator/usersHttp/internal/model"
)

var (
	jsonValidStr = []byte(fmt.Sprintf(`{"Id": 0,
	"Name": "petya",
	"Lastname": "Pupkin",
	"Age": 22,
	"Birthdate": "%v"}`,
		time.Now().UTC().Format(time.RFC3339)))
	jsonInvalidStr = []byte(`{"Id": 0,
	"Name": "petya",
	"Lastname": "Pupkin",
	"Age": 22,
	"Birthdate": "AAA"}`)
)

type commonUserTestCase struct {
	name         string
	jsonStr      []byte
	url          string
	wantStatus   int
	wantBodyRE   string
	mockRetErr   error
	mockRetUsers []*model.User
}

package repositories_test

import (
	"database/sql"
	"errors"
	"time"

	"github.com/goofinator/usersHttp/internal/model"
)

var (
	err0LineResult = errors.New("unexpected number of rows affected: 0")
	err2LineResult = errors.New("unexpected number of rows affected: 2")
)

type userResult struct {
	err    error
	users  []*model.User
	result sql.Result
	rows   *sql.Rows
}

type userTestCase struct {
	name  string
	user  *model.User
	id    int
	txRet userResult
	want  userResult
}

var validUser = &model.User{
	ID:        1,
	Age:       2,
	Name:      "Vasya",
	Lastname:  "Pupkin",
	Birthdate: time.Now(),
}

var validUsers = []*model.User{
	&model.User{
		ID:        1,
		Age:       2,
		Name:      "Vasya",
		Lastname:  "Pupkin",
		Birthdate: time.Now(),
	},
	&model.User{
		ID:        3,
		Age:       4,
		Name:      "Venya",
		Lastname:  "Levkin",
		Birthdate: time.Now(),
	},
}

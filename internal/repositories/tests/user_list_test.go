package repositories_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/goofinator/usersHttp/internal/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var userListTests = []*userTestCase{
	&userTestCase{
		name: "select fail",
		txRet: userResult{
			err:   errSome,
			users: validUsers,
		},
		want: userResult{err: errSome},
	},
	&userTestCase{
		name: "success",
		txRet: userResult{
			err:   nil,
			users: validUsers,
		},
		want: userResult{err: nil},
	},
}

func TestList(t *testing.T) {
	for _, test := range userListTests {
		t.Run(test.name, func(t *testing.T) {
			singleTestList(t, test)
		})
	}
}

// singleTestList performs single test
func singleTestList(t *testing.T, test *userTestCase) {
	// You need a database lock to perform some database actions of service
	mockDB, mock := iniSqlMock(t)
	defer mockDB.Close()
	sqlListExpectations(mock, test)

	repository := repositories.NewUser()
	tx, err := mockDB.Begin()
	require.NoError(t, err)

	users, err := repository.List(tx)
	assert.Equal(t, test.want.err, err)
	if err == nil {
		assert.Equal(t, test.txRet.users, users)
	}

	err = tx.Commit()
	require.NoError(t, err)
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)

}

func sqlListExpectations(mock sqlmock.Sqlmock, test *userTestCase) {
	users := test.txRet.users
	rows := sqlmock.NewRows([]string{"id", "name", "lastname", "birthdate", "age"}).
		AddRow(users[0].ID, users[0].Name, users[0].Lastname, users[0].Birthdate, users[0].Age).
		AddRow(users[1].ID, users[1].Name, users[1].Lastname, users[1].Birthdate, users[1].Age)

	mock.ExpectBegin()
	reqRegExp := &strings.Builder{}
	fmt.Fprint(reqRegExp, "^SELECT id, name, lastname, birthdate, ")
	fmt.Fprint(reqRegExp, "EXTRACT\\(YEAR FROM AGE\\(birthdate\\)\\)[\t \n]+")
	fmt.Fprint(reqRegExp, "AS age FROM http_users ORDER BY ID$")
	ex := mock.ExpectQuery(reqRegExp.String())

	if test.name == "select fail" {
		ex.WillReturnError(test.txRet.err)
	} else {
		ex.WillReturnRows(rows)
	}

	mock.ExpectCommit()
}

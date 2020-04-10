package repositories_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/goofinator/usersHttp/internal/datasource"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

var errSome = errors.New("some error")

func iniSqlMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	datasource.SQL = sqlx.NewDb(mockDB, "sqlmock")
	return mockDB, mock
}

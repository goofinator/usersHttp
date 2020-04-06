package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/goofinator/usersHttp/internal/init/startup"
	"github.com/goofinator/usersHttp/internal/web/model"
	"github.com/jmoiron/sqlx"

	//import pgx sql driver
	_ "github.com/jackc/pgx/v4/stdlib"
)

// Storager interface wraps Close method
type Storager interface {
	Close()
	AddUser(user *model.User) error
	DeleteUser(id int) error
	EditUser(id int, user *model.User) error
}

// Storage is a data storage based on *sqlx.DB
type Storage sqlx.DB

// New returns new data storage based on *sqlx.DB
func New(iniData *startup.IniData) Storager {
	connectionString := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s",
		iniData.DBHost, iniData.DBPort, iniData.DBName, iniData.UserName, iniData.UserPassword)

	db, err := sqlx.Connect("pgx", connectionString)
	if err != nil {
		log.Fatalf("unexpected error on sqlx.Open: %s", err)
	}

	if err := createTable(db, iniData.TableName); err != nil {
		log.Fatalf("unexpected error on createTable: %s", err)
	}

	return (*Storage)(db)

}

// Close closes the storage when it is not needed any more
func (storage *Storage) Close() {
	(*sqlx.DB)(storage).Close()
}

// AddUser adds the user to the storage
func (storage *Storage) AddUser(user *model.User) error {
	tx, err := storage.Begin()
	if err != nil {
		return err
	}
	defer func() {
		err = closeTransaction(tx, err)
	}()

	return nil
}

// DeleteUser deletes the user with specified id from the storage
func (storage *Storage) DeleteUser(id int) error {
	return nil
}

// EditUser replaces data of the user with specified id with user value
func (storage *Storage) EditUser(id int, user *model.User) error {
	return nil
}

// GetUsers returns all users in the storage
func (storage *Storage) GetUsers() []*model.User {
	users := make([]*model.User, 0)
	return users
}

func createTable(db *sqlx.DB, name string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		err = closeTransaction(tx, err)
	}()

	result, err := tx.Exec(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
	id SERIAL PRIMARY KEY,
	name text NOT NULL CHECK(length(name)>0),
	lastname text NOT NULL CHECK(length(lastname)>0),
	age smallint NOT NULL CHECK(age>0),
	birthdate timestamp with time zone NOT NULL
	)`, name))

	if err := checkResult(result, err); err != nil {
		return err
	}
	return nil
}

func closeTransaction(tx *sql.Tx, err error) error {
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			err = fmt.Errorf("%w: %q", err, rbErr)
		}
		return err
	}

	cmtErr := tx.Commit()
	if cmtErr != nil {
		err = cmtErr
	}
	return err
}

func checkResult(result sql.Result, err error) error {
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return fmt.Errorf("unexpected number of rows affected: %d", rowsAffected)
	}
	return nil
}

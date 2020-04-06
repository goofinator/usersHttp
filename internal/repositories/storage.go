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
type Storage struct {
	db      *sqlx.DB
	iniData *startup.IniData
}

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

	return &Storage{db: db, iniData: iniData}

}

// Close closes the storage when it is not needed any more
func (storage *Storage) Close() {
	storage.db.Close()
}

// AddUser adds the user to the storage
func (storage *Storage) AddUser(user *model.User) (err error) {
	tx, err := storage.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		err = closeTransaction(tx, err)
	}()

	return nil
}

// DeleteUser deletes the user with specified id from the storage
func (storage *Storage) DeleteUser(id int) (err error) {
	return nil
}

// EditUser replaces data of the user with specified id with user value
func (storage *Storage) EditUser(id int, user *model.User) (err error) {
	return nil
}

// GetUsers returns all users in the storage
func (storage *Storage) GetUsers() (users []*model.User, err error) {
	users = make([]*model.User, 0)
	return users, nil
}

func createTable(db *sqlx.DB, name string) (err error) {
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

	if err := checkResult(0, result, err); err != nil {
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

func checkResult(needRows int, result sql.Result, err error) error {
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != int64(needRows) {
		return fmt.Errorf("unexpected number of rows affected: %d", rowsAffected)
	}
	return nil
}

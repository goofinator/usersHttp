package datasource

import (
	"database/sql"
	"fmt"

	"github.com/goofinator/usersHttp/internal/init/startup"
	"github.com/goofinator/usersHttp/internal/utils"

	//ini of pgx
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

const (
	// ScriptsDirectory is a base directory where the *.sql scripts
	// will be searched for
	ScriptsDirectory = "./scripts"
	// CreateScript is a name of sqript file to create a table
	CreateScript = "create_db.sql"
)

// SQL stores sqlx database intity
// Before use, the InitSql should be called
var SQL *sqlx.DB

// InitSQL prepares SQL to use
func InitSQL(iniData *startup.IniData) error {
	if err := connect(iniData); err != nil {
		return err
	}

	if err := createTable(); err != nil {
		return err
	}

	return nil
}

// CloseTransaction closes transaction if global error defined
// if no error occured then commit will be called or else - rollback
//
// should use as follow:
// func f() (err error){
//	tx, err := db.Begin()
//	if err != nil {
//		return err
//	}
//	defer func() {
//		err = closeTransaction(tx, err)
//	}()
// 	...
// }
func CloseTransaction(tx *sql.Tx, err error) error {
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

// CheckResult check result,err:=SQL.Exec() pair
func CheckResult(needRows int, result sql.Result, err error) error {
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

func connect(iniData *startup.IniData) error {
	connectionString := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s",
		iniData.DBHost, iniData.DBPort, iniData.DBName, iniData.UserName, iniData.UserPassword)

	db, err := sqlx.Connect("pgx", connectionString)
	if err != nil {
		return fmt.Errorf("unexpected error on sqlx.Open: %s", err)
	}
	SQL = db
	return nil
}

func createTable() (err error) {
	tx, err := SQL.Begin()
	if err != nil {
		return err
	}
	defer func() {
		err = CloseTransaction(tx, err)
	}()

	querries, err := utils.Script(ScriptsDirectory, CreateScript)
	if err != nil {
		return err
	}

	for _, querry := range querries {
		_, err = tx.Exec(querry)
		if err != nil {
			return err
		}
	}
	return nil
}

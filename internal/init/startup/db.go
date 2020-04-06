package startup

import (
	"fmt"
	"log"

	//import pgx sql driver
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

// InitDB prepare database to use
func InitDB(iniData *IniData) *sqlx.DB {
	connectionString := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s",
		iniData.DBHost, iniData.DBPort, iniData.DBName, iniData.UserName, iniData.UserPassword)

	db, err := sqlx.Open("pgx", connectionString)
	if err != nil {
		log.Fatalf("unexpected error on sql.Open: %s", err)
	}

	return db
}

package startup

import (
	"flag"
)

// Default values of IniData fields
const (
	DefaultPort         = 8080
	DefaultDBName       = "postgres"
	DefaultDBHost       = "localhost"
	DefaultDBPort       = 5432
	DefaultUserName     = "postgres"
	DefaultUserPassword = ""
	DefaultTableName    = "postgres"
)

// IniData structure stores initial data to start a service
type IniData struct {
	Port         int
	DBName       string
	DBHost       string
	DBPort       int
	UserName     string
	UserPassword string
	TableName    string
}

// GetIniData returns port to use obtained from user or DefaultPort
func GetIniData() *IniData {
	iniData := &IniData{}
	flag.IntVar(&iniData.Port, "port", DefaultPort, "port to connect this server")
	flag.StringVar(&iniData.DBName, "dbname", DefaultDBName, "data base name")
	flag.StringVar(&iniData.DBHost, "dbhost", DefaultDBHost, "host to access data base")
	flag.IntVar(&iniData.DBPort, "dbport", DefaultDBPort, "port to access data base")
	flag.StringVar(&iniData.UserName, "username", DefaultUserName, "login to connect to database")
	flag.StringVar(&iniData.UserPassword, "password", DefaultUserPassword, "password to connect to database")
	flag.StringVar(&iniData.TableName, "table", DefaultTableName,
		"table name in DB to operate (it will b e created if not exists)")

	flag.Parse()
	return iniData
}

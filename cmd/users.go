package main

import (
	"log"

	"github.com/goofinator/usersHttp/internal/datasource"
	"github.com/goofinator/usersHttp/internal/init/startup"
	"github.com/goofinator/usersHttp/internal/web"
)

func main() {
	iniData := startup.Configuration()
	if err := datasource.InitSQL(iniData); err != nil {
		log.Fatalf("error on datasource.InitSQL: %s", err)
	}
	web.Run(iniData)
}

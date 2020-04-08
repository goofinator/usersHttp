package main

import (
	"github.com/goofinator/usersHttp/internal/datasource"
	"github.com/goofinator/usersHttp/internal/init/startup"
	"github.com/goofinator/usersHttp/internal/web"
)

func main() {
	iniData := startup.Configuration()
	datasource.InitSQL(iniData)
	web.Run(iniData)
}

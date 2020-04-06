package main

import (
	"fmt"

	"github.com/goofinator/usersHttp/init/startup"
)

func main() {
	iniData := startup.GetIniData()
	fmt.Printf("%#v\n", iniData)
}

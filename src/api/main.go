package main

import (
	"github.com/martin98sanch/ABM-user/src/api/server"
)

func main() {
	err := server.New().Run()
	if err != nil {
		panic(err)
	}
}

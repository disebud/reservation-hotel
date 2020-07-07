package main

import (
	"github.com/disebud/reservation-hotel/config"
	"github.com/disebud/reservation-hotel/main/master"
)

func main() {
	db, _ := config.Connection()    //ok
	router := config.CreateRouter() //ok
	master.Init(router, db)
	config.RunServer(router)
}

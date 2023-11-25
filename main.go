package main

import (
	"github.com/amirvalhalla/go-onion-vertical-architecture/api/router"
	"github.com/amirvalhalla/go-onion-vertical-architecture/infrastructure/repository/sql"
	"log"
)

const (
	connStr    = "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tehran"
	ListenPort = ":5050"
)

func main() {
	// initialize unit of work
	uow := sql.NewUnitOfWork()
	if err := uow.Setup(connStr); err != nil {
		log.Panicf("an error occured when unit of work wants to setup err: %s", err.Error())
	}

	// initialize router and setup routes
	app := router.Setup(uow)

	if err := app.Run(ListenPort); err != nil {
		log.Panicf("an error occured for gin router err: %s", err.Error())
	}
}

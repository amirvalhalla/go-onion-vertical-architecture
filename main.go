package main

import (
	"log"

	"github.com/amirvalhalla/go-onion-vertical-architecture/api/router"
	"github.com/amirvalhalla/go-onion-vertical-architecture/infrastructure/repository/sql"
)

const (
	connStr = "host=localhost user=postgres password=postgres dbname=postgres" +
		" port=5432 sslmode=disable TimeZone=Asia/Tehran"
	ListenPort = ":5050"
)

func main() {
	// initialize unit of work
	uow := sql.NewUnitOfWork()
	if err := uow.Setup(connStr); err != nil {
		log.Panicf("an error occurred when unit of work wants to setup err: %s", err.Error())
	}

	// initialize router and setup routes
	app := router.Setup(uow)

	if err := app.Run(ListenPort); err != nil {
		log.Panicf("an error occurred for gin router err: %s", err.Error())
	}
}

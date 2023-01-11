package main

import (
	"database/sql"
	"github.com/amallick86/sanyukta/api"
	"github.com/amallick86/sanyukta/db"
	"github.com/amallick86/sanyukta/util"
	"log"

	_ "github.com/lib/pq"
)

// @title Country API
// @version 1.0
// @description
// @schemes http https
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
// @query.collection.format multi

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load  config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

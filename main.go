package main

import (
	"database/sql"
	"github.com/Tilvaldiyev/banking-system/api"
	"github.com/Tilvaldiyev/banking-system/db/sqlc"
	"github.com/Tilvaldiyev/banking-system/util"
	_ "github.com/golang/mock/mockgen/model"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to DB:", err)
	}

	store := sqlc.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

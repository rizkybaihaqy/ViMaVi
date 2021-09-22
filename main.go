package main

import (
	"log"

	"vip-management-system-api/config"
	"vip-management-system-api/server"
)

func main() {
	cfg := config.LoadConfig()

	db, err := config.CreateConnection(cfg.PgUrl)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	s := server.NewServer(cfg, db)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}

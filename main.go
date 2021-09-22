package main

import (
	"log"
	"vip-management-system-api/config"
	"vip-management-system-api/server"
)

func main() {
	cfg := config.LoadConfig()

	db := config.CreateConnection(cfg.PgUrl)
	defer db.Close()

	s := server.NewServer(cfg, db)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}

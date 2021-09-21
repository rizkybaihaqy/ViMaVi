package main

import (
	"log"
	"vip-management-system-api/config"
	"vip-management-system-api/server"
)

func main() {
	db := config.CreateConnection()
	defer db.Close()

	s := server.NewServer(db)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}

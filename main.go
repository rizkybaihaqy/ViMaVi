package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"vip-management-system-api/routers"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	r := routers.Router()
	h := handlers.AllowedHeaders([]string{"Content-Type"})
	m := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	o := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("Starting server on the port", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), handlers.CORS(h, m, o)(r)))
}

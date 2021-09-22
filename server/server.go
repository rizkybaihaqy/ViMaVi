package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"vip-management-system-api/controllers"
	"vip-management-system-api/models"
	"vip-management-system-api/routers"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

// Server struct
type Server struct {
	db *sql.DB
}

func NewServer(db *sql.DB) *Server {
	return &Server{db: db}
}

func (s *Server) Run() error {
	vModel := models.NewVipModel(s.db)
	vController := controllers.NewVipController(vModel)
	vRoutes := routers.NewVipRoutes(vController)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	r := vRoutes.Router()
	h := handlers.AllowedHeaders([]string{"Content-Type"})
	m := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	o := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("Starting server on the port", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), handlers.CORS(h, m, o)(r)))

	return nil
}

package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"vip-management-system-api/config"
	"vip-management-system-api/controllers"
	"vip-management-system-api/models"
	"vip-management-system-api/routers"

	"github.com/gorilla/handlers"
)

// Server struct
type Server struct {
	Config *config.Config
	DB     *sql.DB
}

func NewServer(cfg *config.Config, db *sql.DB) *Server {
	return &Server{
		Config: cfg,
		DB:     db,
	}
}

func (s *Server) Run() error {
	vm := models.NewVipModel(s.DB)
	vc := controllers.NewVipController(vm)
	vr := routers.NewVipRoutes(vc)

	r := vr.Router()
	h := handlers.AllowedHeaders([]string{"Content-Type"})
	m := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	o := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("Starting server on the port", s.Config.Port)
	log.Fatal(http.ListenAndServe(s.Config.Port, handlers.CORS(h, m, o)(r)))

	return nil
}

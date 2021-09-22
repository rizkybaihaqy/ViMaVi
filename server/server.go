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

type Server struct {
	Config *config.Config
	DB     *sql.DB
}

// Server struct constructor
func NewServer(cfg *config.Config, db *sql.DB) *Server {
	return &Server{
		Config: cfg,
		DB:     db,
	}
}

// Run the server and init other internal packages
func (s *Server) Run() error {
	vm := models.NewVipModel(s.DB)
	vc := controllers.NewVipController(vm)
	vr := routers.NewVipRoutes(vc)

	r := vr.Router()
	h := handlers.AllowedHeaders([]string{"Content-Type"})
	m := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPost, http.MethodPost, http.MethodPatch})
	o := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("Starting server on the port", s.Config.Port)
	log.Fatal(http.ListenAndServe(s.Config.Port, handlers.CORS(h, m, o)(r)))

	return nil
}

package routers

import (
	"vip-management-system-api/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/vips/{id}", controllers.GetVip).Methods("GET")
	r.HandleFunc("/api/vips", controllers.GetVips).Methods("GET")
	r.HandleFunc("/api/vips", controllers.CreateVip).Methods("POST")

	return r
}

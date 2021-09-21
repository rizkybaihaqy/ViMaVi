package routers

import (
	"net/http"
	"vip-management-system-api/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/vips/{id}", controllers.GetVip).Methods(http.MethodGet)
	r.HandleFunc("/api/vips", controllers.GetVips).Methods(http.MethodGet)
	r.HandleFunc("/api/vips", controllers.CreateVip).Methods(http.MethodPost)

	return r
}

package routers

import (
	"net/http"
	c "vip-management-system-api/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/vips/{id}", c.GetVip).Methods(http.MethodGet)
	r.HandleFunc("/api/vips", c.GetVips).Methods(http.MethodGet)
	r.HandleFunc("/api/vips", c.CreateVip).Methods(http.MethodPost)
	r.HandleFunc("/api/vips/{id}", c.UpdateVip).Methods(http.MethodPut)
	r.HandleFunc("/api/vips/{id}", c.DeleteVip).Methods(http.MethodDelete)

	return r
}

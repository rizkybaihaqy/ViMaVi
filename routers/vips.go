package routers

import (
	"net/http"
	"vip-management-system-api/controllers"

	"github.com/gorilla/mux"
)

type VipRoutes struct {
	vController *controllers.VipController
}

func NewVipRoutes(vController *controllers.VipController) *VipRoutes {
	return &VipRoutes{vController: vController}
}

func (vr VipRoutes) Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/vips/{id}", vr.vController.GetVip).Methods(http.MethodGet)
	r.HandleFunc("/api/vips", vr.vController.GetVips).Methods(http.MethodGet)
	r.HandleFunc("/api/vips", vr.vController.CreateVip).Methods(http.MethodPost)
	r.HandleFunc("/api/vips/{id}", vr.vController.UpdateVip).Methods(http.MethodPut)
	r.HandleFunc("/api/vips/{id}", vr.vController.DeleteVip).Methods(http.MethodDelete)

	return r
}

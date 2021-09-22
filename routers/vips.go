package routers

import (
	"net/http"

	"vip-management-system-api/controllers"

	"github.com/gorilla/mux"
)

type VipRoutes struct {
	VC *controllers.VipController
}

func NewVipRoutes(vc *controllers.VipController) *VipRoutes {
	return &VipRoutes{VC: vc}
}

func (vr VipRoutes) Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/vips/{id}", vr.VC.GetVip).Methods(http.MethodGet)
	r.HandleFunc("/api/vips", vr.VC.GetVips).Methods(http.MethodGet)
	r.HandleFunc("/api/vips", vr.VC.CreateVip).Methods(http.MethodPost)
	r.HandleFunc("/api/vips/{id}", vr.VC.UpdateVip).Methods(http.MethodPut)
	r.HandleFunc("/api/vips/{id}", vr.VC.DeleteVip).Methods(http.MethodDelete)

	return r
}

package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"vip-management-system-api/models"
	"vip-management-system-api/utils"

	"github.com/gorilla/mux"
)

type VipController struct {
	VM *models.VipModel
}

func NewVipController(vm *models.VipModel) *VipController {
	return &VipController{VM: vm}
}

// Get all vips
func (c VipController) GetVips(w http.ResponseWriter, r *http.Request) {
	v, err := c.VM.GetVips()
	if v == nil {
		utils.WriteJSON(w, http.StatusNotFound, nil, true, utils.NotFoundMessage)
		return
	}
	if err != nil {
		log.Fatalf("Unable to get all vips. %v", err)
	}

	utils.WriteJSON(w, http.StatusOK, v, true, utils.SuccessRetriveMessage)
}

// Get one vip by id.
func (c VipController) GetVip(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)

	id, err := strconv.Atoi(p["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	v, err := c.VM.GetVip(int64(id))
	if err == sql.ErrNoRows {
		utils.WriteJSON(w, http.StatusNotFound, nil, true, utils.NotFoundMessage)
		return
	}
	if err != nil {
		log.Fatalf("Unable to get vip. %v", err)
	}

	utils.WriteJSON(w, http.StatusOK, v, true, utils.SuccessRetriveMessage)
}

// Create one vip from post request
func (c VipController) CreateVip(w http.ResponseWriter, r *http.Request) {
	var v models.Vip

	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	err = c.VM.InsertVip(v)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	utils.WriteJSON(w, http.StatusCreated, nil, true, utils.SuccessCreateMessage)
}

// Update one vip from post request
func (c VipController) UpdateVip(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)

	id, err := strconv.Atoi(p["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	var v models.Vip

	err = json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	row, err := c.VM.UpdateVip(int64(id), v)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	if row == 0 {
		utils.WriteJSON(w, http.StatusCreated, nil, true, utils.NotFoundMessage)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil, true, utils.SuccessUpdateMessage)
}

// Delete one vip
func (c VipController) DeleteVip(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	row, err := c.VM.DeleteVip(int64(id))
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	if row == 0 {
		utils.WriteJSON(w, http.StatusCreated, nil, true, utils.NotFoundMessage)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil, true, utils.SuccessUpdateMessage)
}

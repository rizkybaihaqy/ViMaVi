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

// Get all vips
func GetVips(w http.ResponseWriter, r *http.Request) {
	v, err := models.GetVips()
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
func GetVip(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)

	id, err := strconv.Atoi(p["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	v, err := models.GetVip(int64(id))
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
func CreateVip(w http.ResponseWriter, r *http.Request) {
	var v models.Vip

	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	err = models.InsertVip(v)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	utils.WriteJSON(w, http.StatusCreated, nil, true, utils.SuccessCreateMessage)
}

// Update one vip from post request
func UpdateVip(w http.ResponseWriter, r *http.Request) {
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

	row := models.UpdateVip(int64(id), v)
	if row == 0 {
		utils.WriteJSON(w, http.StatusCreated, nil, true, utils.NotFoundMessage)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil, true, utils.SuccessUpdateMessage)
}

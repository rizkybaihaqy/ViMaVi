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

// Get one vip by id.
func GetVip(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)

	id, err := strconv.Atoi(p["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	vip, err := models.GetVip(int64(id))
	if err == sql.ErrNoRows {
		utils.WriteJSON(w, http.StatusNotFound, nil, true, utils.NotFoundMessage)
		return
	}
	if err != nil {
		log.Fatalf("Unable to get vip. %v", err)
	}

	utils.WriteJSON(w, http.StatusOK, vip, true, utils.SuccessRetriveMessage)
}

// Create one vip from post request
func CreateVip(w http.ResponseWriter, r *http.Request) {
	var vip models.Vip

	err := json.NewDecoder(r.Body).Decode(&vip)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	err = models.InsertVip(vip)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	utils.WriteJSON(w, http.StatusCreated, nil, true, utils.SuccessCreateMessage)
}

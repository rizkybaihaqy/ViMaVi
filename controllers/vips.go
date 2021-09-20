package controllers

import (
	"database/sql"
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

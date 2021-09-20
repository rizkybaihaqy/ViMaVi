package utils

import (
	"encoding/json"
	"net/http"
)

const SuccessRetriveMessage = "Successfully retrieve data"
const NotFoundMessage = "No data found"

// Construct json response. Takes writer, http.status, response data if any, success status, and some message
func WriteJSON(w http.ResponseWriter, status int, data interface{}, isSuccess bool, msg string) error {
	res := make(map[string]interface{})

	res["ok"] = isSuccess
	res["data"] = data
	res["message"] = msg

	js, err := json.Marshal(res)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

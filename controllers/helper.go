package controllers

import (
	"encoding/json"
	"net/http"
)

func DecodeJson(r *http.Request, w http.ResponseWriter, data any) error {
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return err
	}
	return nil
}

func EncodeJson(w http.ResponseWriter, data any) error {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return err
	}
	return nil
}

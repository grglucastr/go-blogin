package controllers

import (
	"encoding/json"
	"net/http"
)

func GetHellos(w http.ResponseWriter, r *http.Request) {

	aff := []string{"hello", "world", "!!!"}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(aff)
}

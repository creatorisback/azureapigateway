package handlers

import (
	"encoding/json"
	"net/http"
)
type service struct {
	Name    string `json:"name"`
  }
// MicroserviceName to show microservice name
func MicroserviceName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	serviceName := service{"user-microservice"}
	json.NewEncoder(w).Encode(serviceName)
}

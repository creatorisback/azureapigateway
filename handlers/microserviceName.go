package handlers

import "net/http"

// MicroserviceName to show microservice name
func MicroserviceName(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user-microservice"))
}

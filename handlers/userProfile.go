package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

type user struct {
	Username string    `json:"username"`
	Dob      time.Time `json:"dob"`
	Age      int       `json:"age"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
}

// Profile to return username, dob, age, email, and phone-number.
func Profile(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("Username")
	user := user{
		Username: username,
		Dob:      time.Now(),
		Age:      24,
		Email:    "ram.pandey@buzzerlabs.com",
		Phone:    "9871006582",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}

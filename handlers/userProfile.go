package handlers

import "net/http"

// Profile to return username, dob, age, email, and phone-number.
func Profile(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello there"))
}

package handlers

import (
	"net/http"
	"text/template"
)

// Home to display available endpoints
func Home(w http.ResponseWriter, r *http.Request) {
	homePage, err := template.ParseFiles("index.html")
	if err != nil {
		w.Write([]byte("Error executing template"))
	}
	w.WriteHeader(http.StatusOK)
	err = homePage.Execute(w, nil)
}

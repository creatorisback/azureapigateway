package authenticator

import (
	"fmt"
	"net/http"
)

// Authenticator will authenticate requests
func Authenticator(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Username") != "" {
		w.WriteHeader(http.StatusOK)

	} else {
		fmt.Println(r.Header)
		w.WriteHeader(http.StatusUnauthorized)
	}
}

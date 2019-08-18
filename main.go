package main

import (
	"azureApiGateway/authenticator"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	// check if authenticator service is up or not
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/auth", authenticator.Authenticator).Methods("GET")

	// passing true for secured routes and false for unsecured along with handler function
	r.HandleFunc("/user/profile", apiGateway(handleProfile, true)).Methods("GET")
	r.HandleFunc("/microservice/name", apiGateway(microserviceName, false)).Methods("GET")

	// starting server
	log.Fatal(http.ListenAndServe(":8080", r))
}
func microserviceName(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user-microservice"))
}
func handleProfile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello there"))
	w.WriteHeader(http.StatusOK)
}

func apiGateway(handler http.HandlerFunc, secure bool) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if secure {
			r2 := new(http.Request)
			*r2 = *r
			client := &http.Client{}
			req, err := http.NewRequest("GET", "http://localhost:8080/auth", nil)
			req.Header = r2.Header
			if err != nil {
				println("cannot call auth")
			}
			resp, err := client.Do(req)

			if resp.StatusCode == 200 {
				handler.ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
			}
			fmt.Println(resp)

		} else {
			handler.ServeHTTP(w, r)
		}
	})
}

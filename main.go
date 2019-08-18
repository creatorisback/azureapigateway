package main

import (
	"log"
	"net/http"

	"github.com/creatorisback/azureapigateway/authenticator"
	"github.com/creatorisback/azureapigateway/gateway"
	"github.com/creatorisback/azureapigateway/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Home)
	r.HandleFunc("/auth", authenticator.Authenticator).Methods("GET")

	// passing true for secured routes and false for unsecured along with handler function
	r.HandleFunc("/user/profile", gateway.APIGateway(handlers.Profile, true)).Methods("GET")
	r.HandleFunc("/microservice/name", gateway.APIGateway(handlers.MicroserviceName, false)).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	// starting server at 8080
	log.Fatal(http.ListenAndServe(":8080", r))
}

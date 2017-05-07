package main

import (
	"github.com/Rakanixu/brandcrumb/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	// Pluggable implementation
	_ "github.com/Rakanixu/brandcrumb/db/fish/inmemory"
	_ "github.com/Rakanixu/brandcrumb/db/tank/inmemory"
)

func main() {
	// Rest API Router
	r := mux.NewRouter()

	// CRUD tank operations (missing update)
	r.HandleFunc(handlers.CREATE_TANK, handlers.CreateTankHandler).Methods(http.MethodPost, http.MethodPut)
	r.HandleFunc(handlers.READ_TANK, handlers.ReadTankHandler).Methods(http.MethodGet)
	r.HandleFunc(handlers.DELETE_TANK, handlers.DeleteTankHandler).Methods(http.MethodDelete)

	// CRUD fish operations (missing update)
	r.HandleFunc(handlers.CREATE_FISH, handlers.CreateFishHandler).Methods(http.MethodPost, http.MethodPut)
	r.HandleFunc(handlers.READ_FISH, handlers.ReadFishHandler).Methods(http.MethodGet)
	r.HandleFunc(handlers.DELETE_FISH, handlers.DeleteFishHandler).Methods(http.MethodDelete)

	// Bind to a port to the router
	log.Fatal(http.ListenAndServe(":8081", r))
}

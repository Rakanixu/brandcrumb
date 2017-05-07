package handlers

import (
	"encoding/json"
	db "github.com/Rakanixu/brandcrumb/db/tank"
	"github.com/Rakanixu/brandcrumb/models/tank"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

const (
	CREATE_TANK = "/tank"
	READ_TANK   = "/tank/{id}"
	DELETE_TANK = "/tank/{id}"
)

// CreateTankHandler creates a new tank
func CreateTankHandler(w http.ResponseWriter, r *http.Request) {
	var t *tank.Tank

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		// Invalid JSON body, 400 bad request
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create tank in DB and handle status code
	handleHTTPStatusCode(w, db.Create(t))
}

// ReadTankHandler returns specified tank
func ReadTankHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	// Malformed URL
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Read from DB
	t, err := db.Read(int64(id))
	if err != nil {
		handleHTTPStatusCode(w, err)
		return
	}

	b, err := json.Marshal(t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

// ReadTankHandler returns specified tank
func DeleteTankHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	// Malformed URL
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Delete tank from DB and handle status code
	handleHTTPStatusCode(w, db.Delete(int64(id)))
}

package handlers

import (
	"encoding/json"
	db "github.com/Rakanixu/brandcrumb/db/fish"
	"github.com/Rakanixu/brandcrumb/models/fish"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

const (
	CREATE_FISH = "/fish"
	READ_FISH   = "/fish/{id}"
	DELETE_FISH = "/fish/{id}"
)

// CreateFishHandler creates a new Fish
func CreateFishHandler(w http.ResponseWriter, r *http.Request) {
	var f *fish.Fish

	if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
		// Invalid JSON body, 400 bad request
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create fish and handle status code
	handleHTTPStatusCode(w, db.Create(f))
}

// ReadFishHandler returns specified Fish
func ReadFishHandler(w http.ResponseWriter, r *http.Request) {
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

	// 200 OK by default
}

// ReadFishHandler returns specified Fish
func DeleteFishHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	// Malformed URL
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Delete fish from DB and handle status code
	handleHTTPStatusCode(w, db.Delete(int64(id)))
}

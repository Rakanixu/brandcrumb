package handlers

import (
	"net/http"
	"strconv"
)

// handleHTTPStatusCode is a helper to set proper status code for addressing all possibilities
// 400 when parsing invalid data
// 403 if operation not allowed (eg. delete tank that contains at least a fish)
// 404 if resource is not stored (eg. create a fish into a tank that does not exits yet)
// 500 internal server error
func handleHTTPStatusCode(w http.ResponseWriter, err error) {
	if err != nil {
		// Parse http error and return StatusCode over the wire
		sc, _ := strconv.Atoi(err.Error())
		w.WriteHeader(sc)
	}
}

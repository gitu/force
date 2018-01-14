package main

import (
	"encoding/json"
	"net/http"
)

// GroupsAPI is API implementation of /groups root endpoint
type GroupsAPI struct {
}

// Get is the handler for GET /groups
func (api GroupsAPI) Get(w http.ResponseWriter, r *http.Request) {
	var respBody Groups
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&respBody)
}

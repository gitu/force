package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

// EntriesAPI is API implementation of /entries root endpoint
type EntriesAPI struct {
}

// Get is the handler for GET /entries
func (api EntriesAPI) Get(w http.ResponseWriter, r *http.Request) {
	groupIds := r.FormValue("groupIds")
	respBody := GetEntries(strings.Split(groupIds, ","))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&respBody)
}

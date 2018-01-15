package main

import (
	"encoding/json"
	"net/http"
)

// EntriesAPI is API implementation of /entries root endpoint
type EntriesAPI struct {
}

// Get is the handler for GET /entries
func (api EntriesAPI) Get(w http.ResponseWriter, r *http.Request) { // groupIds := req.FormValue("groupIds")
	var respBody []Entry
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&respBody)
}

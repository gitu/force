package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

// GroupsAPI is API implementation of /groups root endpoint
type GroupsAPI struct {
}

// groupIdGet is the handler for GET /groups/{groupId}
// Gets a list of all entries for this group
func (api GroupsAPI) groupIdGet(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	respBody := GetEntries(path[len(path)-1:])
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&respBody)

}

// Get is the handler for GET /groups
func (api GroupsAPI) Get(w http.ResponseWriter, r *http.Request) {
	respBody := GetGroups()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&respBody)
}

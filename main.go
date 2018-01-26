package main

import (
	"log"
	"net/http"

	"github.com/gitu/force/goraml"

	"bytes"
	"github.com/gorilla/mux"
	"gopkg.in/validator.v2"
	"io"
)

//go:generate go-bindata-assetfs -pkg main apidocs/... index.html entries.json

func serveIndex(rw http.ResponseWriter, req *http.Request) {
	bs, _ := Asset("index.html")
	var reader = bytes.NewBuffer(bs)
	io.Copy(rw, reader)
}

func main() {
	// input validator
	validator.SetValidationFunc("multipleOf", goraml.MultipleOf)

	r := mux.NewRouter()

	// home page
	r.HandleFunc("/", serveIndex)
	r.PathPrefix("/client/").Handler(http.StripPrefix("/client/", http.FileServer(http.Dir("./client/"))))

	// apidocs
	r.PathPrefix("/apidocs/").Handler(http.StripPrefix("/apidocs/", http.FileServer(assetFS())))

	EntriesInterfaceRoutes(r, EntriesAPI{})

	GroupsInterfaceRoutes(r, GroupsAPI{})

	log.Println("starting server visit http://localhost:5000 to access it")
	log.Println("   if you want to serve your own files put them in the folder ./client")
	log.Println("   then visit to http://localhost:5000/client/")
	err := http.ListenAndServe(":5000", r)
	if err != nil {
		log.Fatal(err)
	}
}

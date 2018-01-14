package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
	"log"
	"os"
)

var logger *log.Logger

func main() {
	logger := log.New(os.Stdout, "[force]   ", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, hello())
	})

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(mux)

	listenOn := ":3000"

	logger.Printf("will listen on %s", listenOn)
	http.ListenAndServe(listenOn, n)
}

func hello() string {
	return "Hello World"
}

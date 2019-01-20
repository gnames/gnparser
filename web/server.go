package web

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"gitlab.com/gogna/gnparser/fs"
)

// Run starts RESTful service on /api route for both GET and
// POST verbs
func Run(port int, wn int) {
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/doc/api", docAPI)
	router.HandleFunc("/api", apiGetParse).Methods("GET").Queries("q", "{q}")
	router.HandleFunc("/api", apiPostParse).Methods("POST")
	router.HandleFunc("/api", apiEmptyRequest)
	router.PathPrefix("/").Handler(http.FileServer(fs.Files))
	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:" + strconv.Itoa(port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

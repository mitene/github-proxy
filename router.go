package gproxy

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() error {
	r := mux.NewRouter()
	r.HandleFunc("/get/{owner}/{repo}", GetHandler)
	r.HandleFunc("/get/{owner}/{repo}", GetHandler).
		Queries("ref", "{ref}").
		Queries("type", "{type}").
		Queries("path", "{path}")
	http.Handle("/", r)

	log.Println("server started")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}

	return nil
}

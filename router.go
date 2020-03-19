package proxy

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(port int) error {
	if port == 0 {
		log.Fatalln("invalid port")
	}
	r := mux.NewRouter()
	r.HandleFunc("/repo/{owner}/{repo}", RepoHandler)
	r.HandleFunc("/repo/{owner}/{repo}", RepoHandler).
		Queries("ref", "{ref}").
		Queries("type", "{type}").
		Queries("path", "{path}")
	http.Handle("/", r)

	log.Printf("server started: port: %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		return err
	}

	return nil
}

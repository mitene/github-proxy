package main

import (
	"log"

	"github.com/mitene/gproxy"
)

func main() {
	err := gproxy.NewRouter()
	if err != nil {
		log.Fatal(err)
	}
}

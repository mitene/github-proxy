package main

import (
	"flag"
	"log"

	"github.com/mitene/gproxy"
)

func main() {
	port := flag.Int("port", 8080, "http port")
	flag.Parse()

	err := gproxy.NewRouter(*port)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"flag"
	"log"

	proxy "github.com/mitene/github-proxy"
)

func main() {
	port := flag.Int("port", 8080, "http port")
	flag.Parse()

	err := proxy.NewRouter(*port)
	if err != nil {
		log.Fatal(err)
	}
}

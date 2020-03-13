
.PHONY: build
build:
	go build -o bin/gproxy ./cmd/gproxy

.PHONY: run
run:
	go run ./cmd/gproxy

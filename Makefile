
.PHONY: build
build:
	go build -o bin/github-proxy ./cmd/github-proxy

.PHONY: run
run:
	go run ./cmd/github-proxy

.PHONY: fmt
fmt:
	go fmt

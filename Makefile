.PHONY: build
build:
	go build -v ./cmd/pars
.DEFAULT_GOAL := build
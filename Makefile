-include .env

VERSION := $(shell git describe --tags)

compile:
	go build -o pg-ping -ldflags "-s -w -X main.version=$(VERSION)" main.go

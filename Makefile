-include .env

ifndef PGPING_VERSION
	PGPING_VERSION := $(shell git describe --tags)-dirty
endif

compile:
	go build -o pg-ping -ldflags "-s -w -X main.version=$(PGPING_VERSION)" main.go

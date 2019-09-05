PROJECT_NAME ?= {{cookiecutter.project_name}}
SHELL 	 := /bin/bash
executable := {{cookiecutter.project_name}}
targetOS?= linux
##make clean: deletes all binaries inside bin/
clean:
	@rm -rf bin/*
##make build: generates a binary at bin/$(executable). You can define the target OS you want for the binary by passing targetOS. (e.g. make build targetOS = windows) FYI default targetOS is linux.
build-api:
	@echo Building $(executable)
	GOOS=$(targetOS) GO111MODULE=on go build -o bin/$(executable) cmd/$(executable)/main.go
##make mock: generates mock_dbclient needed for unit testing on handlers at package service
mock-dbclient:
	mockgen -source internal/pkg/api/dbclient/dbclient.go -destination internal/pkg/api/dbclient/mock_dbclient/mock_dbclient.go -package mock_dbclient

test:
	go test ./...

.PHONY : help
help : Makefile
	@sed -n 's/^##//p' $<

NAME=keycloak-gatekeeper
GIT_SHA=$(shell git --no-pager describe --always --dirty)
BUILD_TIME=$(shell date '+%s')
LFLAGS ?= -X main.gitsha=${GIT_SHA} -X main.compiled=${BUILD_TIME}
ARCHITECTURES=amd64

.PHONY: install build

install:
	go install -v

default: build

golang:
	@echo "--> Go Version"
	@go version


build: golang
	@echo "--> Compiling the project"
	@mkdir -p bin
	go build -ldflags "${LFLAGS}" -o bin/${NAME}



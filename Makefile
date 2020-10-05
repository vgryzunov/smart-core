NAME=smart-core
GIT_SHA=$(shell git --no-pager describe --always --dirty)
BUILD_TIME=$(shell date '+%s')
LFLAGS ?= -X main.gitsha=${GIT_SHA} -X main.compiled=${BUILD_TIME}
ARCHITECTURES=amd64

.PHONY: install fmt gen validate clean build

install:
	go install -v

default: build

golang:
	@echo "--> Go Version"
	@go version

fmt:
	go fmt

gen: validate
	swagger generate server \
		-t ./hello-plugin/swagger \
		-f ./hello-plugin/swagger/swagger.yml \
		--exclude-main \
		-A hello

validate:
	swagger validate ./hello-plugin/swagger/swagger.yml

clean:
	rm -rf ./bin/* 2>/dev/null

build: golang clean
	@echo "--> Compiling the project"
	@mkdir -p bin
	go build -ldflags "${LFLAGS}" -o bin/${NAME}
	go build -buildmode=plugin -o ./bin ./hello-plugin

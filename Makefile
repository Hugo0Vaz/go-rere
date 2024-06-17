OUT      := grr
PKG      := github.com/Hugo0vaz/go-rere
VERSION  := $(shell git describe --always --long --dirty)
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/)

build:
	@go build -o ./build/grr

run: build
	@./build/grr

.PHONY: build run

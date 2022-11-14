GOPATH:=$(shell go env GOPATH)


GO_SOURCES := $(wildcard *.go)
GO_SOURCES += $(shell find . -type f -name "*.go")

GOFMT ?= gofmt -s

proto3:
	  protoc --proto_path=eventstream --go_out=eventstream --go_opt=paths=source_relative \
	  --go-grpc_out=eventstream --go-grpc_opt=paths=source_relative \
		eventstream/*.proto

tools:
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
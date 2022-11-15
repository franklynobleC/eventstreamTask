GOPATH:=$(shell go env GOPATH)


GO_SOURCES := $(wildcard *.go)


GOFMT ?= gofmt -s

proto1:
	  protoc --proto_path=pb --go_out=pb --go_opt=paths=source_relative \
	  --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		pb/*.proto




proto4:
		protoc --proto_path=eventstore --go_out=eventstore --go_opt=paths=source_relative \
	   --go-grpc_out=eventstore --go-grpc_opt=paths=source_relative \
		eventstore/*.proto
tools:
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 



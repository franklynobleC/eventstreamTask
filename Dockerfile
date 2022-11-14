FROM golang:1.16-alpine AS builder
RUN mkdir -p /go/src/workspace/eventStreamTask
WORKDIR /go/src/workspace/eventStreamTask
COPY . .

RUN GIT_TERMINAL_PROMPT=1 \
    GOARCH=amd64 \
    GOOS=linux \
    CGO_ENABLED=0 \
    go build -v --installsuffix cgo --ldflags="-s" -o eventStreamTask
FROM alpine:3.13

# convert build-arg to env variables


RUN mkdir -p /newFile/
COPY --from=builder /go/src/workspace/eventStreamTask  /newFile/

WORKDIR /newFile/

CMD ["./eventStreamTask"]
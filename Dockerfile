FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/demo
COPY . .
EXPOSE 7777
RUN go build ./cmd/server

ENTRYPOINT [ "./server" ]
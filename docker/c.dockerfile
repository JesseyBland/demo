FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /home/jessey/go/src/demo
COPY . .
EXPOSE 6060
RUN go build ./cmd/client

ENTRYPOINT [ "./client" ]
FROM golang:alpine

WORKDIR /go-TAB

ADD . .

RUN go mod download

ENTRYPOINT go build && ./go-TAB
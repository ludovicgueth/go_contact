FROM golang:latest

RUN mkdir -p /go/src/api
WORKDIR /go/src/api

ADD . /go/src/api

RUN go get -v
FROM golang:1.8

WORKDIR /go/src/github.com/phanyzewski/book_api

ADD . /go/src/github.com/phanyzewski/book_api

RUN go get -v ./...

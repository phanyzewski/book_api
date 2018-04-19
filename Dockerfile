FROM golang:1.8

WORKDIR /go/src/github.com/phanyzewski/book_api

RUN go get github.com/rnubel/pgmgr
ADD . /go/src/github.com/phanyzewski/book_api
RUN go get github.com/rnubel/pgmgr

RUN go get -d -v ./...
RUN go install -v ./...

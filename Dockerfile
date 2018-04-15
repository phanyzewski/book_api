FROM golang:1.8

WORKDIR /go/src/github.com/phanyzewski/book_api
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["book_api"]

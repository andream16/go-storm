FROM golang:1.9

RUN mkdir -p /go/src/github.com/andream16/go-storm
WORKDIR /go/src/github.com/andream16/go-storm

ADD . /go/src/github.com/andream16/go-storm

RUN go get -v
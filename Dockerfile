FROM golang:1.7.0-alpine

ENV PORT 8081

WORKDIR /go/src/github.com/illyabusigin/goa-cellar
COPY . /go/src/github.com/illyabusigin/goa-cellar

RUN go-wrapper install

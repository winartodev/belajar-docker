FROM golang:1.17-alpine AS builder

RUN mkdir /api

ADD . /api

WORKDIR  /api

RUN go mod download

RUN go build app/main.go

EXPOSE 8080

CMD /api/main
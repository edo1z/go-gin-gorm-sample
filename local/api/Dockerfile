FROM golang:1.17.7-alpine

WORKDIR /api
RUN apk update && apk add git
RUN go install github.com/cosmtrek/air@latest

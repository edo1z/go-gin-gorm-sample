FROM golang:1.17.7-alpine

WORKDIR /api
RUN apk update && apk add git && \
    go install github.com/cosmtrek/air@latest && \
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

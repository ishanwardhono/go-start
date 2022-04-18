FROM golang:1.18

WORKDIR /app

COPY go.mod go.sum makefile ./
RUN go mod download

FROM golang:1.18

WORKDIR /app

COPY go.mod go.sum makefile ./
RUN go mod download

# COPY *.go ./
# COPY /env /env

RUN go build -o /app

EXPOSE 8080


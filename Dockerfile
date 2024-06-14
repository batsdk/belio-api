# syntax=docker/dockerfile:1
FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /belio-api

EXPOSE 8080

CMD ["/belio-api"]

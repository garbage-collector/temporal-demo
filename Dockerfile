# syntax=docker/dockerfile:1
FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY api/main.go ./

RUN go build -o /api

EXPOSE 8080

CMD [ "/api" ]

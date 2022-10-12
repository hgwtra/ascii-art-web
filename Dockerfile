# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
# COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/ascii-art-web-main

EXPOSE 8080

CMD [ "/app/ascii-art-web-main" ]
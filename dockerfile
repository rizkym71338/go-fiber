FROM golang:alpine

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app
COPY . .
RUN go mod tidy
RUN air init
CMD air
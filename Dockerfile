# Use an official Golang runtime as a parent image
FROM golang:1.21-alpine as builder

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o main ./cmd/app/main.go

CMD ["./main"]

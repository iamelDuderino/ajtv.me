# syntax=docker/dockerfile:1
FROM golang:1.19
ADD . /app
WORKDIR /app
RUN go mod download
RUN go build -o "./bin/andrewjtomko" /app/main.go
EXPOSE 8080
ENTRYPOINT ["/app/bin/andrewjtomko"]
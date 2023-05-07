# syntax=docker/dockerfile:1
FROM golang:1.19
COPY go.mod go.sum /
WORKDIR /app
COPY . .
COPY handler.go /
RUN go mod download
RUN go build -o /handler.go
EXPOSE 8080
CMD ["/handler"]
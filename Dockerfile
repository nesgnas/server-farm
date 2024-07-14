# syntax=docker/dockerfile:1

FROM golang:alpine

# Set destination for COPY
WORKDIR /server-farm

COPY . .

RUN go mod download

RUN go build -o bin .

EXPOSE 8000

ENTRYPOINT ["/server-farm/bin"]

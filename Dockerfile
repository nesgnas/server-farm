# syntax=docker/dockerfile:1

FROM golang:1.22

# Set destination for COPY
WORKDIR /server-farm

COPY . .

RUN go mod download

RUN go build -o bin .

EXPOSE 8448

ENTRYPOINT ["/server-farm/bin"]
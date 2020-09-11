#First Stage
FROM golang:alpine AS builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fritzdect-prom-exporter

#Second Stage
FROM alpine:latest AS production

COPY --from=builder /app/fritzdect-prom-exporter .

ENTRYPOINT ["./fritzdect-prom-exporter"]
FROM golang:1.18-bullseye AS builder

RUN mkdir -p /build

WORKDIR /build

COPY . .

RUN go build -o /build/go-web .

# Main image
FROM debian:11-slim

RUN mkdir /app

WORKDIR /app

COPY --from=builder /build/go-web .

ENV GIN_MODE=release

EXPOSE 8080

CMD ["/app/go-web"]

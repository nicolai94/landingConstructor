FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

RUN go build -o main .

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/main /app/main
COPY --from=builder /go/bin/goose /usr/local/bin/goose

RUN apt-get update && apt-get install -y libpq-dev

EXPOSE 8080

CMD ["./main"]
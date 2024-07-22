FROM golang:1.20.3-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o ./bin/crud_server cmd/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/bin/crud_server .

CMD ["./crud_server"]
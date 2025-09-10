FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o main ./cmd/go-s3/main.go
RUN chmod +x main

FROM alpine:3.18

COPY --from=builder /app/main /main
COPY --from=builder /app/.env ./.env

EXPOSE 8080

CMD ["/main"]
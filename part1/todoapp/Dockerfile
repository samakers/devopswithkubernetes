FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
COPY .env .env
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o todoapp

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/todoapp .
COPY --from=builder /app/.env .
CMD ["./todoapp"]
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o random-string-app

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/random-string-app .
CMD ["./random-string-app"]

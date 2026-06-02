# Build stage
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o order-service ./cmd/server

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/order-service ./order-service
COPY ./migrations ./migrations
CMD ["./order-service"]

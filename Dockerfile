# Stage 1: Build the Go binary
FROM golang:1.19 AS builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o backend

FROM alpine:3.18

WORKDIR /app
COPY --from=builder /app/backend .

EXPOSE 3000

CMD ["./backend"]
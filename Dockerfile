# BUILD
FROM golang:1.21-alpine AS builder
WORKDIR /app/backend
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o main .

# FINAL STAGE
FROM ubuntu:latest
WORKDIR /app
COPY --from=builder /app/backend/main .
EXPOSE 8080
CMD ["./main"]
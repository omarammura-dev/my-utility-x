# BUILD
FROM golang:1.21-alpine 

WORKDIR /app/backend
COPY go.mod Gopkg.toml ./
RUN go mod download
COPY . .
RUN go build -o main .
FROM ubuntu:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["main"]
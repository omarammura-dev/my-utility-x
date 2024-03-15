# BUILD
FROM golang:1.21-alpine as BUILD
ENV GO111MODULE=on
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go mod tidy  # Add this line
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
ENV HTTP_PORT=8080
EXPOSE 8080
# BINARIES
FROM alpine:latest
COPY --from=BUILD /app/production /app/production
ENTRYPOINT ["/app/production"]
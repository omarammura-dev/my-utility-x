# BUILD


FROM --platform=linux/arm64 golang:1.21-alpine AS builder
WORKDIR /app/backend
ARG MONGO_URL
ARG MONGO_DB_NAME
ARG API_URL
ARG UI_URL

# use the ARG values as environment variables
ENV MONGO_URL=$MONGO_URL
ENV MONGO_DB_NAME=$MONGO_DB_NAME
ENV API_URL=$API_URL
ENV UI_URL=$UI_URL

# create an .env file with the environment variables
RUN echo "MONGO_URL=$MONGO_URL" > .env && \
    echo "MONGO_DB_NAME=$MONGO_DB_NAME" >> .env && \
    echo "API_URL=$API_URL" >> .env && \
    echo "UI_URL=$UI_URL" >> .env
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o main .
# FINAL STAGE
FROM ubuntu:latest
WORKDIR /app
COPY --from=builder /app/backend/main .
EXPOSE 8080
CMD ["./main"]
# BUILD
FROM golang:1.21-alpine
WORKDIR /usr/src/app/backend
# ARG MONGO_URL
# ARG MONGO_DB_NAME
# ARG API_URL
# ARG UI_URL

# ENV MONGO_URL=$MONGO_URL
# ENV MONGO_DB_NAME=$MONGO_DB_NAME
# ENV API_URL=$API_URL
# ENV UI_URL=$UI_URL

# RUN echo "MONGO_URL=$MONGO_URL" > .env && \
#     echo "MONGO_DB_NAME=$MONGO_DB_NAME" >> .env && \
#     echo "API_URL=$API_URL" >> .env && \
#     echo "UI_URL=$UI_URL" >> .env
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o main .
# # FINAL STAGE
# FROM ubuntu:latest
# WORKDIR /app
# COPY --from=builder /usr/src/app/backendmain .
# COPY --from=builder /usr/src/app/backend.env . 
# FROM alpine:latest
# WORKDIR /app
# COPY --from=builder /usr/src/app/backend/main .

EXPOSE 8080
CMD ["./main"]
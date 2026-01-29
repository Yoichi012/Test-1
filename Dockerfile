# Build stage
FROM golang:1.20-alpine AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bot ./cmd/bot

# Run stage
FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=build /bot /bot
ENV TELEGRAM_TOKEN=""
ENV MONGO_URI="mongodb://mongo:27017"
ENV MONGO_DB="waifu_db"
EXPOSE 8080
CMD ["/bot"]
FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app ./cmd/app

FROM alpine:latest
RUN apk add --no-cache libc6-compat
WORKDIR /root/
COPY --from=builder /app/app .
COPY --from=builder /app/templates /root/templates
EXPOSE 8080
CMD ["./app"]

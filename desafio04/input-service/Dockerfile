FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o input-service .

FROM alpine:latest

RUN apk --no-cache add ca-certificates curl

WORKDIR /root/

COPY --from=builder /app/input-service .

EXPOSE 8000

CMD ["./input-service"]

FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o manomano ./cmd/api/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/manomano .
COPY migrations ./migrations
EXPOSE 8000
CMD ["./manomano"]



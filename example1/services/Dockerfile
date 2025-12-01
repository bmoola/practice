# syntax=docker/dockerfile:1

FROM golang:1.24.6 AS builder

WORKDIR /app

COPY .. .

WORKDIR /app/services

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/server .

FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 50051

ENTRYPOINT ["./server"]
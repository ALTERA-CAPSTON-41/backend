# builder
FROM golang:1.17.5-alpine3.14 AS builder

WORKDIR /app
COPY ./ ./
RUN go mod download
RUN go build -o main

# runner
FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8000

CMD ["./main"]

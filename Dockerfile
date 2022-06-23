# builder
FROM golang:1.18-alpine3.14 AS builder

WORKDIR /app
COPY ./ ./
RUN go mod download
RUN go build -o main

# runner
FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main ./main
COPY --from=builder /app/public ./public
COPY --from=builder /app/libs/api-spec ./libs/api-spec
EXPOSE 8000

CMD ["./main"]

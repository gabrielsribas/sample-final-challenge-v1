FROM golang:1.22.2-alpine AS builder

ENV CGO_ENABLED=1

WORKDIR /app

RUN apk update && apk add --no-cache \
    build-base \
    sqlite-dev

COPY product-api /app/product-api

WORKDIR product-api

RUN go mod download

#COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/product-api/main .
#COPY --from=builder /app/product-api/data ./data

EXPOSE 8080

CMD ["./main"]


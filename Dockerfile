FROM golang:1.22.2-alpine AS builder

WORKDIR /app

COPY go.* /app/
RUN go mod download

COPY . .

RUN ls && go build -o main .

#FROM alpine:latest

#WORKDIR /app

#COPY --from=builder /app/main .
#COPY --from=builder /app/data ./data

#EXPOSE 8080

#CMD ["./main"]


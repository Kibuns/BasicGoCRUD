FROM golang:1.20-alpine:3.13 AS builder
WORKDIR /app
COPY . .
RUN go build -o main.go

FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 10000
CMD [ "/app/main" ]

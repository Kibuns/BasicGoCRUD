FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main.go

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 10000
CMD [ "/app/main" ]

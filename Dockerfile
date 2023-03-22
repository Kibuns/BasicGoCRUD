FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY ../

RUN 1s
RUN go build -o /StratService

EXPOSE 10000

CMD [ "/StratService" ]

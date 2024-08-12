FROM golang:1.21.4-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /meli-proxy

EXPOSE 8080

CMD ["/meli-proxy"]

FROM golang:1.23.3-alpine3.20

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o load-tester ./cmd

ENTRYPOINT ["./load-tester"]

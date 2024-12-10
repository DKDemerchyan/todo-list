FROM golang:1.23.3

WORKDIR /app

COPY . .

RUN go mod tidy

RUN GOOS=linux GOARCH=amd64 go build -o /todo

CMD ["/todo"]
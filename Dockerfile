FROM golang:1.21

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main cmd/app/main.go

EXPOSE 8888

CMD ["./main"]
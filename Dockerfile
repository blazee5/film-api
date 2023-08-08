FROM golang:latest

COPY . /app

WORKDIR /app/cmd/user/server

RUN go mod download && go build

CMD ["./server"]
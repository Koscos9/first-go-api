FROM golang:latest

WORKDIR /api/
COPY go.mod go.sum /api/
COPY . /api/

RUN go build -o main .

CMD ["./main"]
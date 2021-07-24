FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /go-todo

EXPOSE 3000

ENTRYPOINT [ "/go-todo" ]

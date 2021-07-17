FROM golang:1.16-alpine

WORKDIR /go/src/github.com/chirst/go-todo

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /go-todo

CMD [ "/go-todo" ]

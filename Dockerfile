##
## Build
##

FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /go-todo

##
## Deploy
##

FROM alpine:latest

WORKDIR /app

COPY --from=build /go-todo /go-todo

COPY . .

EXPOSE 3000

ENTRYPOINT [ "/go-todo" ]

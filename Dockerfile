FROM golang:alpine
RUN apk update & apk add git
RUN mkdir /go/src/app
WORKDIR /go/src/app
COPY . /go/src/app/
RUN go install github.com/cosmtrek/air@latest
CMD air

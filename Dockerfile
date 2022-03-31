FROM golang:alpine
RUN apk update & apk add git
RUN mkdir /go/src/app
WORKDIR /go/src/app
COPY . /go/src/app/

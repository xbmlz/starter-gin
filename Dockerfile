FROM golang:alpine as builder

ENV GOPROXY https://goproxy.cn,direct

WORKDIR $GOPATH/src/github.com/xbmlz/starter-gin

COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .


FROM alpine:latest

LABEL MAINTAINER="chenxc170016@gmail.com"

WORKDIR $GOPATH/src/github.com/xbmlz/starter-gin

EXPOSE 8000

ENTRYPOINT ["./server"]
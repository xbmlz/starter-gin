FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct

WORKDIR $GOPATH/src/github.com/xbmlz/starter-gin

COPY . $GOPATH/src/github.com/xbmlz/starter-gin

RUN go build .

EXPOSE 8000

ENTRYPOINT ["./starter-gin"]
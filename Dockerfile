FROM golang:1.19 as builder

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /code

ADD . /code

RUN make build

FROM alpine:latest

WORKDIR /app

COPY --from=builder /code/bin/ /app
COPY --from=builder /code/config.yaml /app

EXPOSE 8080

ENTRYPOINT ["/app/bin/server"]
# 配置编译环境
FROM golang:1.16-alpine AS builder

# 环境配置
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=http://goproxy.cn,direct

COPY . /go/src/httpserver

WORKDIR /go/src/httpserver
RUN go install .

FROM alpine:3.13

COPY --from=builder /go/bin/httpserver /bin/httpserver

ENV ADDR=:8080

EXPOSE 8080

ENTRYPOINT [ "/bin/httpserver" ]
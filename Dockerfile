FROM golang:1.18-alpine as build

ENV http_proxy=http://wan-proxy.transwarp.io:3128
ENV https_proxy=http://wan-proxy.transwarp.io:3128

WORKDIR /sophon

COPY . /sophon
RUN  CGO_ENABLED=0 GOOS=linux go build -o /sophon/protoc-go-inject-tag-sophon
RUN chmod a+x /sophon/protoc-go-inject-tag-sophon


FROM 172.16.1.99/aip/deps/protogen:master

COPY --from=build /sophon/protoc-go-inject-tag-sophon /usr/local/bin/
COPY ./protofmt /usr/local/bin/


# docker build -t 172.16.1.99/aip/deps/protogen:with-inject-tag .
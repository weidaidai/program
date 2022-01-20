#指定版本
FROM golang:1.16-alpine

ENV GOPROXY="https://goproxy.io"
ENV GO111MODULE="on"

RUN mkdir /student_app

WORKDIR /student_app

COPY  ./ .
RUN go build -mod=vendor -v

EXPOSE 8080

ENTRYPOINT student_app/program





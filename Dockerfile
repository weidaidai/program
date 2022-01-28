FROM golang:1.16-alpine

RUN mkdir /student

WORKDIR /student

COPY  ./ .
RUN go build -mod=vendor -v

FROM harbor.supwisdom.com/institute/alpine:latest

COPY --from=0  /student/program  home/app/program

EXPOSE 8080

ENTRYPOINT /home/app/program





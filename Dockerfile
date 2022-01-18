#指定版本
FROM golang:1.16-alpine

#环境
ENV GOPROXY="https://goproxy.io"
ENV GO111MODULE="on"

#创目录
RUN mkdir /student_app

#移
WORKDIR /student_app

#复制项目mod和sum,并下载依赖包
COPY go.mod .
COPY go.sum .
RUN go mod tidy

#添/复制代码到容器
#COPY . .
ADD . /student_app


#生成exe文件
RUN go build -o  program .


EXPOSE 8080





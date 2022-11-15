FROM golang:latest as builder
ENV TZ=Europe/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENV GO111MODULE=on

WORKDIR /app

COPY ./src ./src
RUN go mod init Avito-Challenge
RUN go mod tidy

WORKDIR ./src
RUN go build -o main main.go

CMD ./main -docker

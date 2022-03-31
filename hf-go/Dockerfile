FROM ubuntu:20.04

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install curl -y

WORKDIR /tmp
RUN curl https://dl.google.com/go/go1.17.8.linux-amd64.tar.gz -o go.tar.gz && \
    tar -zxf go.tar.gz && \
    mv go /go
WORKDIR /
RUN rm -rf tmp

ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin
FROM ubuntu:14.04
MAINTAINER Dongri Jin <dongriab@gmail.com>

RUN apt-get update
RUN apt-get install -y curl git bzr mercurial

RUN curl -s https://go.googlecode.com/files/go1.2.linux-amd64.tar.gz | tar -v -C /usr/local/ -xz

ENV PATH  /usr/local/go/bin:/usr/local/bin:/usr/local/sbin:/usr/bin:/usr/sbin:/bin:/sbin
ENV GOPATH  /go
ENV GOROOT  /usr/local/go

#RUN go get github.com/codegangsta/negroni
RUN go get github.com/dongri/go-web

WORKDIR /go/src/github.com/dongri/go-web

#RUN go get
#RUN go build

EXPOSE 3000

ENTRYPOINT go run server.go

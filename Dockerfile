FROM golang:alpine

RUN apk add --no-cache git mercurial

LABEL Max-Report.version="1.0.10" maintainer="PharbersDevelopers"

ENV MAXVIEW_HOME /go/bin

RUN go get github.com/alfredyang1986/blackmirror && \
go get github.com/alfredyang1986/BmServiceDef && \
go get github.com/PharbersDeveloper/Max-Report

RUN go install -v github.com/PharbersDeveloper/Max-Report

ADD resource /go/bin/resource

WORKDIR /go/bin

ENTRYPOINT ["Max-Report"]

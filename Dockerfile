FROM golang:alpine

RUN apk add --no-cache git mercurial

ENV MAXVIEW_HOME /go/bin

# 以LABEL行的变动(version的变动)来划分(变动以上)使用cache和(变动以下)不使用cache
LABEL Max-Report.version="1.0.14" maintainer="developer@pharbers.com"

RUN go get github.com/alfredyang1986/blackmirror && \
go get github.com/alfredyang1986/BmServiceDef && \
go get github.com/PharbersDeveloper/Max-Report

ADD resource /go/bin/resource

RUN go install -v github.com/PharbersDeveloper/Max-Report

WORKDIR /go/bin

ENTRYPOINT ["Max-Report"]

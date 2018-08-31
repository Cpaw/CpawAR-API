FROM api-baseimage:1.0 AS Package-Env

ENV GOPATH /go/
WORKDIR /go/

ADD ./ApiServer /go/src/ApiServer
RUN revel package ApiServer prod
RUN tar xzvf ApiServer.tar.gz

FROM aicompimage:1.0 AS Package-Env

ENV GOPATH /go/
WORKDIR /go/

ADD ./AiCompServer /go/src/AiCompServer
RUN revel package AiCompServer prod
RUN tar xzvf AiCompServer.tar.gz

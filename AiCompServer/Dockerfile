FROM aicompimage:1.0 AS Package-Env

ENV GOPATH /go/
ARG DBHOST
ARG HOST=0.0.0.0
WORKDIR /go/src

RUN git clone https://github.com/Cpaw/AiCompServer.git /go/src/AiCompServer
RUN revel package AiCompServer
RUN mv AiCompServer.tar.gz /go/
WORKDIR /go/
RUN tar xzvf AiCompServer.tar.gz

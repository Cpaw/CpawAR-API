FROM golang:1.8.1

ENV GOPATH $GOPATH:/go/src

RUN apt-get update && \
    apt-get upgrade -y

# install revel and revel-cli
RUN go get github.com/revel/revel && \
    go get github.com/revel/cmd/revel

WORKDIR $GOPATH:/go/src
RUN git clone https://github.com/Cpaw/AiCompServer.git /go/src/AiCompServer
RUN go get github.com/bradfitz/gomemcache/memcache
RUN go get github.com/garyburd/redigo/redis
RUN go get github.com/jinzhu/gorm
RUN go get github.com/lib/pq
RUN go get github.com/patrickmn/go-cache
RUN go get golang.org/x/crypto/scrypt
RUN go get gopkg.in/validator.v2
RUN revel package AiCompServer

RUN tar xzvf AiCompServer.tar.gz

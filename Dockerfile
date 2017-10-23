FROM golang:1.9
MAINTAINER kc merrill <kcmerrill@gmail.com>
RUN go get github.com/kcmerrill/oogway
RUN mkdir /oogway
WORKDIR /oogway
ENTRYPOINT ["oogway"]

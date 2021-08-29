FROM golang:1.16

RUN go get github.com/mitchellh/gox

WORKDIR /w

ENTRYPOINT ["gox"]

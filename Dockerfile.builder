FROM golang:1.14

RUN go get github.com/mitchellh/gox

WORKDIR /w

ENTRYPOINT ["gox"]

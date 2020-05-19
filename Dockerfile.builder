FROM golang:1.14

RUN go get github.com/mitchellh/gox

WORKDIR /w
ENV GOMODENABLED=1

ENTRYPOINT ["gox"]

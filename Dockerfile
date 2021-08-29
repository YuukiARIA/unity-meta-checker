FROM golang:1.16 AS builder

COPY . /workspace
WORKDIR /workspace

ENV CGO_ENABLED 0

RUN go test -v ./...
RUN go build -ldflags="-s -w --extldflags -static" -o build/unity-meta-checker cmd/unity-meta-checker/main.go

FROM scratch

COPY --from=builder /workspace/build/unity-meta-checker /usr/local/bin/unity-meta-checker
WORKDIR /workspace

ENTRYPOINT ["unity-meta-checker"]

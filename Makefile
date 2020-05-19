SRCS       := $(shell find . -name "*.go" -type f)
BUILD      := builds
LDFLAGS    := -s -w -extldflags -static
GOX_OSARCH := darwin/amd64 linux/amd64 windows/amd64
GOX_OUTPUT := $(BUILD)/{{.OS}}_{{.Arch}}/{{.Dir}}

.PHONY: build docker-build-builder docker-build test

build: $(SRCS) go.mod go.sum docker-build-builder
	docker run --rm -v "$(shell pwd):/w" unity-meta-checker-builder -ldflags="$(LDFLAGS)" -osarch="$(GOX_OSARCH)" -output="$(GOX_OUTPUT)" ./...

docker-build-builder: Dockerfile.builder
	docker build -t unity-meta-checker-builder -f Dockerfile.builder .

docker-build:
	docker build -t unity-meta-checker .

test:
	go test -v ./...

clean:
	-rm -rf $(BUILD)

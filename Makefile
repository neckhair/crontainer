BINARY = crontainer
GOARCH = amd64

GOFILES    = $(shell find . -name '*.go' -not -path './vendor/*')
GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)

LINUX_BINARY = ${BINARY}-linux-${GOARCH}

BASE = $(shell pwd)

.PHONY: build test clean

all: test $(LINUX_BINARY)

vendor: $(GOFILES) Gopkg.toml Gopkg.lock
	dep ensure

test: vendor
	go test -v ./...

build:
	go build -o bin/crontainer

install:
	go install

run:
	go run main.go --config examples/crontainer.yml

clean:
	-rm -f ${BINARY}-*

$(LINUX_BINARY): $(GOFILES)
	cd $(BASE); \
	GOOS=linux GOARCH=${GOARCH} go build -o ${LINUX_BINARY} . ;

docker:
	docker build -t neckhair/crontainer .

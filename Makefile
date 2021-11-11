.PHONY: build install fmt test test-install

build:
	cd aurora && env GO111MODULE=on go build -v ./...

test:
	cd aurora && gocheck -c

test-install:
	go get github.com/frankbraun/gocheck
	go get golang.org/x/tools/cmd/goimports
	go get golang.org/x/lint/golint

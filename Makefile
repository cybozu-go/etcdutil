# Makefile for etcdutil

.PHONY: all
all:
	@echo "Specify one of these targets:"
	@echo
	@echo "    test        - run single host tests."
	@echo "    setup       - install dependencies."

.PHONY: test
test: setup
	test -z "$(gofmt -s -d . | tee /dev/stderr)"
	staticcheck .
	go build ./...
	go test -race -v ./...
	go vet ./...

.PHONY: setup
setup:
	cd /tmp; env GOFLAGS= GO111MODULE=on go get golang.org/x/tools/cmd/goimports
	cd /tmp; env GOFLAGS= GO111MODULE=on go get honnef.co/go/tools/cmd/staticcheck

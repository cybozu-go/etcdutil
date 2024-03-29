# Makefile for etcdutil

ETCD_VER=$(shell go list -m -f '{{.Version}}' go.etcd.io/etcd/client/v3)
SUDO=sudo

.PHONY: all
all:
	echo $(ETCD_VER)
	@echo "Specify one of these targets:"
	@echo
	@echo "    test        - run single host tests."
	@echo "    setup       - install dependencies."

.PHONY: test
test:
	staticcheck .
	go build ./...
	go test -race -v ./...
	go vet ./...

.PHONY: check-generate
check-generate:
	go mod tidy
	git diff --exit-code --name-only

.PHONY: setup
setup:
	cd /tmp; go install golang.org/x/tools/cmd/goimports@latest
	cd /tmp; go install honnef.co/go/tools/cmd/staticcheck@latest
	curl -L https://github.com/etcd-io/etcd/releases/download/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -o /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
	mkdir /tmp/etcd
	tar xzvf /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz -C /tmp/etcd --strip-components=1
	$(SUDO) mv /tmp/etcd/etcd /usr/local/bin/
	rm -rf /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz /tmp/etcd 

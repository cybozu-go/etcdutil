# Makefile for etcdutil

## Dependency versions
ETCD_VER=v3.5.9

SUDO=sudo

.PHONY: all
all:
	@echo "Specify one of these targets:"
	@echo
	@echo "    test        - run single host tests."
	@echo "    setup       - install dependencies."

.PHONY: test
test:
	test -z "$(gofmt -s -d . | tee /dev/stderr)"
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

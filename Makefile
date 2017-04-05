MAIN_VERSION:=$(shell git describe --abbrev=0 --tags || echo "0.1.0")
VERSION:=${MAIN_VERSION}\#$(shell git log -n 1 --pretty=format:"%h")
PACKAGES:=$(shell glide novendor)
PACKAGE_WITHOUT_E2E:=$(shell glide novendor | grep -v -E '(./testdata|./daos|./apis)')
LDFLAGS:=-ldflags "-X github.com/ederavilaprado/golang-web-architecture-template/app.Version=${VERSION}"

default: test

deps:
	glide up

test:
	go test -p=1 $(PACKAGES)

unit-test:
	go test $(PACKAGE_WITHOUT_E2E)

cover:
	echo "mode: count" > coverage-all.out
	$(foreach pkg,$(PACKAGES), \
		go test -p=1 -cover -covermode=count -coverprofile=coverage.out ${pkg}; \
		tail -n +2 coverage.out >> coverage-all.out;)
	go tool cover -html=coverage-all.out

run:
	go run ${LDFLAGS} main.go

dev:
	CompileDaemon -exclude-dir ".git" -color -build "go build -o _build_hot_reload" -command "./_build_hot_reload"

build: clean
	go build ${LDFLAGS} -a -o server main.go

clean:
	rm -rf _build_hot_reload coverage.out coverage-all.out

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
BINARY_NAME=bumpit
BINARY_UNIX=$(BINARY_NAME)_unix

.PHONY: all test clean bench
all: test build
build: 
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/$(BINARY_NAME)
test: 
	$(GOTEST) ./test/... -v
bench:
	$(GOTEST) -bench=. ./test/... -benchmem
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/$(BINARY_NAME)
	./$(BINARY_NAME)
deps:
	$(GOMOD) download

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
docker-build:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v
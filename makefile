# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
BINARY_NAME=bumpit
BINARY_DARWIN=$(BINARY_NAME)-darwin-amd64
BINARY_LINUX=$(BINARY_NAME)-linux-amd64

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
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_LINUX) ./cmd/$(BINARY_NAME)
build-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_DARWIN) ./cmd/$(BINARY_NAME)

docker-build:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v
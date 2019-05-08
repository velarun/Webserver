# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=webserver
BINARY_UNIX=$(BINARY_NAME)_unix

all: deps go

go: go/build

go/build:
	@echo "Building $(BIN_NAME)"
	@go version
	$(GOBUILD) -o bin/$(BINARY_NAME)
	@chmod -R 777 bin/$(BINARY_NAME)

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

deps:
	$(GOGET) github.com/go-sql-driver/mysql
	$(GOGET) github.com/gorilla/mux

    
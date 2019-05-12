# Go parameters
GOCMD=go
PYCMD=python3
PIPCMD=pip3

GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=webserver
BINARY_UNIX=$(BINARY_NAME)_unix
PYPIP=$(PIPCMD) install

all: deps go pip

go: go/build

go/build:
	@echo "Building $(BIN_NAME)"
	@go version
	$(GOBUILD) -o $(BINARY_NAME)
	@chmod 777 $(BINARY_NAME)

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

pip:
	$(PYPIP) requests
	$(PYPIP) beautifulsoup4



    
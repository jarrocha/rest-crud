# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
GOINSTALL=$(GOCMD) install
BINARY_NAME=rest-crud
    
all: test build

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
test: 
	$(GOTEST) -v ./controllers
install:
		$(GOINSTALL)
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run:
	$(GORUN) ./*.go
deps:
	$(GOGET) gopkg.in/mgo.v2
	$(GOGET) gopkg.in/mgo.v2/bson
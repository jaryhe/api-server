# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
PROJECT_NAME=storm
BIN_DIR=/usr/sbin
CP_CMD=/usr/bin/cp
COMMAND=./storm
LNCMD=ln
LNSCMD=$(LNCMD) -s

all: build

build:
	$(GOBUILD) -o $(PROJECT_NAME) -v  -ldflags "-X main.Version=$(version) -X main.GitCommit=`git rev-parse HEAD` "
test:
	$(GOTEST) -v ./...
install:build
	$(CP_CMD) $(COMMAND) $(DESTDIR)$(BIN_DIR)
	$(LNSCMD) $(DESTDIR)$(BIN_DIR)/storm $(DESTDIR)$(BIN_DIR)/stormd
clean:
	$(GOCLEAN)
	rm -f $(PROJECT_NAME)



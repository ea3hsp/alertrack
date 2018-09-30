
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=alertrack
BINARY_LINUX=$(BINARY_NAME)_linux
BINARY_RPI=$(BINARY_NAME)_rpi
USER_RPI=pi
NAME_RPI=rpi03


all: test build
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
test: 
	$(GOTEST) -v
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME)
# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
build-rpi:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=5 $(GOBUILD) -o $(BINARY_RPI) -v
	scp $(BINARY_RPI) $(USER_RPI)@$(NAME_RPI):/home/$(USER_RPI)/bin/$(BINARY_NAME)

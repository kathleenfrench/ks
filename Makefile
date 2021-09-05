INSTALL_LOCATION := /usr/local/bin
BIN := ks
PWD = $(shell pwd)
DIST := $(PWD)/build
BIN_OUT := $(DIST)/$(BIN)

build: clean ## compile the ks binary to the workspace's root build directory
	@mkdir -p build
	@go build -o $(BIN_OUT) main.go

install: build ## install the ks binary to /usr/local/bin
	@cp $(BIN_OUT) $(INSTALL_LOCATION)
	@chmod 755 $(INSTALL_LOCATION)/$(BIN)

clean:
	@rm -rf build

help: ## lists useful (but not all) commands, see the Makefile for more.
	@awk 'BEGIN {FS = ":.*##"; printf "\Local `make` Commands:\n \033[36m\033[0m\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: build install clean help
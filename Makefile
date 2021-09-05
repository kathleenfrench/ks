
build: clean ## compile the ks binary to the workspace's root build directory
	@mkdir -p build
	@go build -o build/ks main.go

install: build ## install the ks binary to /usr/local/bin
	@mv build/ks /usr/local/bin
	@chmod 755 /usr/local/bin/ks

clean:
	@rm -rf build

help: ## lists useful (but not all) commands, see the Makefile for more.
	@awk 'BEGIN {FS = ":.*##"; printf "\Local `make` Commands:\n \033[36m\033[0m\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: build install clean help
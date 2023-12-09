VERSION := $(shell git describe --tags --always --dirty="-dev")

all : clean build

v:
	@echo $(VERSION)

clean:
	@echo "Cleaning..."
	@rm -rf ./build

build:
	@echo "Building..."
	@mkdir -p ./build
	@go build -trimpath -ldflags "-X main.version=${VERSION}" -v -o build/node cmd/node/main.go

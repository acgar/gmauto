LINUX_BUILD_NAME := gmauto-linux-amd64
WIN_BUILD_NAME := gmauto-windows-amd64.exe
BUILD_DIR := build

.PHONY: all build run test clean fmt vet

build: build-linux build-windows

build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(LINUX_BUILD_NAME)

build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(WIN_BUILD_NAME)

run:
	go run .

test:
	go test ./...

fmt:
	go fmt ./...

# Detect some warnings
vet:
	go vet ./...

clean:
	go clean ./...
	rm -f $(BUILD_DIR)/$(LINUX_BUILD_NAME) $(BUILD_DIR)/$(WIN_BUILD_NAME)

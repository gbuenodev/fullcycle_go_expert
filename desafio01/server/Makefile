BINARY_NAME=exchange_server
BUILD_OUTPUT_DIR=./bin
SRC=./

.PHONY: build run clean test

build: clean
	mkdir -p $(BUILD_OUTPUT_DIR)
	go build -o $(BUILD_OUTPUT_DIR)/$(BINARY_NAME) $(SRC)

run: build
	$(BUILD_OUTPUT_DIR)/$(BINARY_NAME)

clean:
	rm -rf $(BUILD_OUTPUT_DIR)

test: 
	go test -v ./...

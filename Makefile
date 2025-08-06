BINARY_NAME=memory
BUILD_DIR=build

.PHONY: build clean install test run

build:
	mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 go build -o $(BUILD_DIR)/$(BINARY_NAME) .

install: build
	sudo cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/

clean:
	rm -rf $(BUILD_DIR)
	go clean

test:
	CGO_ENABLED=0 go test ./...

run:
	go run . $(ARGS)

deps:
	go mod tidy
	go mod download

# 示例用法
demo: build
	./$(BUILD_DIR)/$(BINARY_NAME) add "我的GitHub密码是secret123"
	./$(BUILD_DIR)/$(BINARY_NAME) add "明天需要完成Go项目"
	./$(BUILD_DIR)/$(BINARY_NAME) list
	./$(BUILD_DIR)/$(BINARY_NAME) stats
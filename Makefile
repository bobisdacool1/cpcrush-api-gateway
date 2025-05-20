.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	CONFIG_PATH=$(shell pwd)/configs/config.yaml go test -v ./...
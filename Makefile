git_hook:
	git config core.hooksPath ./hooks

tools:
	@echo "Installing dependencies..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/incu6us/goimports-reviser/v3@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install go.uber.org/mock/mockgen@latest
	@echo "Dependencies installed successfully!"

init: git_hook tools
	@echo "Init done!"

lint:
	# 简化 go mod 文件
	go mod tidy
	# 调整 imports 顺序
	goimports-reviser -company-prefixes github.com/MoeGolibrary -set-alias -rm-unused -format ./...
	# 校验代码规范
	golangci-lint run -v --allow-parallel-runners --fix

build:
	@echo "Building..."
	go build -o server ./internal
run :
	@echo "Running..."
	./server -config ./config/config-test.yaml
.PHONY: init

.PHONY: help build run test clean setup dev

# 默认目标
help:
	@echo "Gin Auth Project - 可用命令:"
	@echo "  setup    - 设置项目环境"
	@echo "  build    - 构建项目"
	@echo "  run      - 运行项目"
	@echo "  dev      - 开发模式运行（自动重载）"
	@echo "  test     - 运行测试"
	@echo "  clean    - 清理构建文件"
	@echo "  deps     - 安装依赖"
	@echo "  lint     - 代码检查"

# 设置项目环境
setup:
	@echo "🚀 设置项目环境..."
	@chmod +x scripts/setup.sh
	@./scripts/setup.sh

# 安装依赖
deps:
	@echo "📦 安装Go依赖..."
	@go mod tidy
	@go mod download

# 构建项目
build:
	@echo "🔨 构建项目..."
	@go build -o bin/gin-auth-project main.go

# 运行项目
run:
	@echo "🚀 运行项目..."
	@go run main.go

# 开发模式运行（需要安装air）
dev:
	@echo "🔄 开发模式运行..."
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "⚠️  Air未安装，使用普通模式运行..."; \
		go run main.go; \
	fi

# 运行测试
test:
	@echo "🧪 运行测试..."
	@go test ./...

# 运行测试并显示覆盖率
test-coverage:
	@echo "🧪 运行测试并显示覆盖率..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out

# 代码检查
lint:
	@echo "🔍 代码检查..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "⚠️  golangci-lint未安装，跳过代码检查"; \
	fi

# 格式化代码
fmt:
	@echo "✨ 格式化代码..."
	@go fmt ./...
	@go vet ./...

# 清理构建文件
clean:
	@echo "🧹 清理构建文件..."
	@rm -rf bin/
	@rm -f coverage.out

# 数据库迁移
migrate:
	@echo "🗄️  运行数据库迁移..."
	@go run main.go migrate

# 创建新用户
create-user:
	@echo "👤 创建新用户..."
	@go run cmd/create-user/main.go

# 显示项目信息
info:
	@echo "📊 项目信息:"
	@echo "  Go版本: $(shell go version)"
	@echo "  模块名: $(shell go list -m)"
	@echo "  依赖数量: $(shell go list -m all | wc -l)" 
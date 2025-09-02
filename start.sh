#!/bin/bash

echo "🚀 启动 Gin Auth Project..."

# 检查Go是否安装
if ! command -v go &> /dev/null; then
    echo "❌ Go未安装，请先安装Go 1.21+"
    exit 1
fi

# 检查.env文件是否存在
if [ ! -f .env ]; then
    echo "⚠️  环境配置文件不存在，正在创建..."
    if [ -f env.example ]; then
        cp env.example .env
        echo "✅ 已创建 .env 文件，请编辑配置信息后重新运行"
        echo "📝 需要配置："
        echo "   - 数据库连接信息"
        echo "   - Redis连接信息"
        echo "   - JWT密钥"
        exit 1
    else
        echo "❌ 找不到 env.example 文件"
        exit 1
    fi
fi

# 安装依赖
echo "📦 检查并安装依赖..."
go mod tidy

# 检查PostgreSQL连接
echo "🗄️  检查数据库连接..."
if command -v psql &> /dev/null; then
    # 尝试连接数据库（这里只是简单检查，实际连接在应用启动时进行）
    echo "✅ PostgreSQL客户端可用"
else
    echo "⚠️  PostgreSQL客户端不可用，请确保PostgreSQL已安装并运行"
fi

# 检查Redis连接
echo "🔴 检查Redis连接..."
if command -v redis-cli &> /dev/null; then
    # 尝试ping Redis
    if redis-cli ping &> /dev/null; then
        echo "✅ Redis连接正常"
    else
        echo "⚠️  Redis连接失败，请确保Redis服务正在运行"
    fi
else
    echo "⚠️  Redis客户端不可用，请确保Redis已安装并运行"
fi

echo ""
echo "🎯 启动应用..."
echo "📡 服务器将在 http://localhost:8080 启动"
echo "🔍 健康检查: http://localhost:8080/health"
echo "📚 API文档请查看 README.md"
echo ""
echo "按 Ctrl+C 停止服务器"
echo ""

# 启动应用
go run main.go 
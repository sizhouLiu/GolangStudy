#!/bin/bash

echo "🚀 开始设置 Gin Auth Project..."

# 检查Go是否安装
if ! command -v go &> /dev/null; then
    echo "❌ Go未安装，请先安装Go 1.21+"
    exit 1
fi

echo "✅ Go已安装: $(go version)"

# 检查PostgreSQL是否安装
if ! command -v psql &> /dev/null; then
    echo "⚠️  PostgreSQL未安装，请先安装PostgreSQL"
    echo "    macOS: brew install postgresql"
    echo "    Ubuntu: sudo apt-get install postgresql postgresql-contrib"
    echo "    CentOS: sudo yum install postgresql postgresql-server"
else
    echo "✅ PostgreSQL已安装: $(psql --version)"
fi

# 检查Redis是否安装
if ! command -v redis-cli &> /dev/null; then
    echo "⚠️  Redis未安装，请先安装Redis"
    echo "    macOS: brew install redis"
    echo "    Ubuntu: sudo apt-get install redis-server"
    echo "    CentOS: sudo yum install redis"
else
    echo "✅ Redis已安装: $(redis-cli --version)"
fi

# 安装Go依赖
echo "📦 安装Go依赖..."
go mod tidy

if [ $? -eq 0 ]; then
    echo "✅ Go依赖安装成功"
else
    echo "❌ Go依赖安装失败"
    exit 1
fi

# 创建环境配置文件
if [ ! -f .env ]; then
    echo "📝 创建环境配置文件..."
    cp env.example .env
    echo "✅ 环境配置文件已创建，请编辑 .env 文件配置数据库连接信息"
else
    echo "✅ 环境配置文件已存在"
fi

# 创建数据库（如果PostgreSQL可用）
if command -v psql &> /dev/null; then
    echo "🗄️  创建数据库..."
    psql -U postgres -c "CREATE DATABASE gin_auth_db;" 2>/dev/null || echo "⚠️  数据库创建失败，请手动创建数据库 'gin_auth_db'"
fi

echo ""
echo "🎉 项目设置完成！"
echo ""
echo "📋 接下来的步骤："
echo "1. 编辑 .env 文件，配置数据库和Redis连接信息"
echo "2. 确保PostgreSQL和Redis服务正在运行"
echo "3. 运行项目: go run main.go"
echo ""
echo "📚 更多信息请查看 README.md 文件" 
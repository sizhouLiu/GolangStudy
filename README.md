# Gin Auth Project

一个基于Gin框架的完整后端项目，包含用户认证、权限管理、Redis缓存和PostgreSQL数据库。

## 功能特性

- 🔐 JWT认证和授权
- 👥 用户注册、登录、登出
- 🛡️ 基于角色的权限控制（RBAC）
- 💾 PostgreSQL数据库集成
- 🚀 Redis缓存支持
- 🔒 密码加密存储
- 📝 完整的CRUD操作
- 🌐 CORS支持
- 📊 分页查询
- 🗑️ 软删除

## 技术栈

- **Web框架**: Gin
- **数据库**: PostgreSQL + GORM
- **缓存**: Redis
- **认证**: JWT
- **密码加密**: bcrypt
- **配置管理**: godotenv

## 项目结构

```
gin-auth-project/
├── config/          # 配置管理
├── database/        # 数据库连接和Redis
├── handlers/        # 请求处理器
├── middleware/      # 中间件
├── models/          # 数据模型
├── routes/          # 路由配置
├── utils/           # 工具函数
├── main.go          # 主程序入口
├── go.mod           # Go模块文件
├── env.example      # 环境变量示例
└── README.md        # 项目说明
```

## 快速开始

### 1. 环境要求

- Go 1.21+
- PostgreSQL 12+
- Redis 6+

### 2. 安装依赖

```bash
go mod tidy
```

### 3. 配置环境变量

复制 `env.example` 为 `.env` 并修改配置：

```bash
cp env.example .env
```

编辑 `.env` 文件，配置数据库和Redis连接信息：

```env
# 数据库配置
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=gin_auth_db
DB_SSL_MODE=disable

# Redis配置
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0

# JWT配置
JWT_SECRET=your_jwt_secret_key_here
JWT_EXPIRE_HOURS=24

# 服务器配置
SERVER_PORT=8080
SERVER_MODE=debug
```

### 4. 创建数据库

```sql
CREATE DATABASE gin_auth_db;
```

### 5. 运行项目

```bash
go run main.go
```

服务器将在 `http://localhost:8080` 启动。

## API接口

### 认证接口

- `POST /api/auth/login` - 用户登录
- `POST /api/auth/register` - 用户注册
- `POST /api/auth/logout` - 用户登出
- `GET /api/auth/profile` - 获取用户信息
- `PUT /api/auth/profile` - 更新用户信息
- `POST /api/auth/refresh` - 刷新令牌

### 用户管理接口（需要管理员权限）

- `GET /api/users` - 获取所有用户
- `POST /api/users` - 创建新用户
- `GET /api/users/:id` - 根据ID获取用户
- `PUT /api/users/:id` - 更新用户信息
- `DELETE /api/users/:id` - 删除用户
- `PATCH /api/users/:id/status` - 切换用户状态

### 受保护资源接口（需要用户权限）

- `GET /api/protected/data` - 获取受保护的数据

## 权限系统

项目实现了基于角色的访问控制（RBAC）：

- **admin**: 管理员，可以访问所有功能
- **user**: 普通用户，可以访问基本功能
- **guest**: 访客，只能访问公开功能

## 默认用户

项目启动时会自动创建默认管理员用户：

- 用户名: `admin`
- 邮箱: `admin@example.com`
- 密码: `password`
- 角色: `admin`

## 缓存策略

- 用户令牌存储在Redis中，支持令牌撤销
- 用户信息缓存，提高查询性能
- 支持缓存过期和手动清除

## 安全特性

- 密码使用bcrypt加密存储
- JWT令牌支持过期时间
- 基于角色的权限控制
- 软删除保护数据完整性
- 输入验证和SQL注入防护

## 开发建议

1. 在生产环境中，请修改默认的JWT密钥
2. 配置适当的数据库连接池参数
3. 设置Redis密码和网络访问限制
4. 使用HTTPS保护API通信
5. 定期备份数据库

## 贡献

欢迎提交Issue和Pull Request！

## 许可证

MIT License 
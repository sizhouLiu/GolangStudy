# 项目结构说明

```
gin-auth-project/
├── 📁 config/                    # 配置管理
│   └── config.go                # 应用配置和环境变量管理
│
├── 📁 database/                  # 数据库和缓存
│   ├── postgres.go              # PostgreSQL数据库连接和初始化
│   └── redis.go                 # Redis缓存连接和操作
│
├── 📁 handlers/                  # 请求处理器
│   ├── auth.go                  # 认证相关处理器（登录、注册、登出等）
│   └── user.go                  # 用户管理处理器（CRUD操作）
│
├── 📁 middleware/                # 中间件
│   ├── auth.go                  # JWT认证和权限控制中间件
│   └── cors.go                  # 跨域请求处理中间件
│
├── 📁 models/                    # 数据模型
│   └── user.go                  # 用户模型和数据结构定义
│
├── 📁 routes/                    # 路由配置
│   └── routes.go                # API路由定义和中间件配置
│
├── 📁 utils/                     # 工具函数
│   ├── jwt.go                   # JWT令牌生成和验证
│   └── password.go              # 密码加密和验证
│
├── 📁 tests/                     # 测试文件
│   └── auth_test.go             # 认证功能测试
│
├── 📁 scripts/                   # 脚本文件
│   └── setup.sh                 # 项目环境设置脚本
│
├── 📁 postman/                   # API测试集合
│   └── Gin_Auth_Project.postman_collection.json  # Postman测试集合
│
├── 📄 main.go                    # 主程序入口
├── 📄 go.mod                     # Go模块依赖管理
├── 📄 env.example                # 环境变量配置示例
├── 📄 README.md                  # 项目说明文档
├── 📄 PROJECT_STRUCTURE.md       # 项目结构说明（本文件）
├── 📄 Makefile                   # 项目构建和开发命令
├── 📄 Dockerfile                 # Docker容器化配置
├── 📄 docker-compose.yml         # Docker Compose服务编排
└── 📄 start.sh                   # 项目启动脚本
```

## 核心模块说明

### 1. 配置管理 (config/)
- 使用 `godotenv` 管理环境变量
- 支持数据库、Redis、JWT等配置
- 提供默认值和配置验证

### 2. 数据库层 (database/)
- **PostgreSQL**: 使用GORM作为ORM，支持自动迁移
- **Redis**: 缓存用户令牌和会话信息
- 自动创建默认管理员用户

### 3. 认证系统 (handlers/auth.go + middleware/auth.go)
- JWT令牌认证
- 基于角色的权限控制 (RBAC)
- 支持令牌刷新和撤销
- 密码加密存储 (bcrypt)

### 4. 用户管理 (handlers/user.go)
- 完整的用户CRUD操作
- 分页查询支持
- 软删除保护
- 权限验证

### 5. 中间件系统
- **认证中间件**: 验证JWT令牌
- **权限中间件**: 基于角色的访问控制
- **CORS中间件**: 处理跨域请求

### 6. 工具函数 (utils/)
- JWT令牌操作
- 密码加密和验证
- 可复用的通用功能

## 技术特点

### 🔐 安全性
- JWT令牌认证
- 密码bcrypt加密
- 基于角色的权限控制
- 输入验证和SQL注入防护

### 🚀 性能
- Redis缓存支持
- 数据库连接池
- 软删除保护数据完整性

### 🛠️ 开发体验
- 自动数据库迁移
- 完整的错误处理
- 详细的日志记录
- 健康检查端点

### 📱 部署友好
- Docker容器化支持
- 环境变量配置
- 健康检查
- 优雅关闭

## 开发流程

1. **环境设置**: 运行 `./scripts/setup.sh`
2. **配置环境**: 编辑 `.env` 文件
3. **启动服务**: 运行 `./start.sh` 或 `make run`
4. **API测试**: 使用Postman集合测试接口
5. **开发调试**: 使用 `make dev` 启动热重载

## 扩展建议

- 添加日志系统 (如logrus)
- 集成监控和指标收集
- 添加API文档生成 (如Swagger)
- 实现文件上传功能
- 添加邮件验证
- 集成第三方登录 (OAuth)
- 添加单元测试和集成测试
- 实现API限流和熔断 
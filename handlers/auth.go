package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"gin-auth-project/database"
	"gin-auth-project/models"
	"gin-auth-project/utils"

	"gin-auth-project/middleware"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct{}

// 用户登录
func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	// 查找用户
	var user models.User
	if err := database.DB.Where("username = ? OR email = ?", req.Username, req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// 验证密码
	if !utils.CheckPassword(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// 检查用户是否激活
	if !user.IsActive {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User account is deactivated"})
		return
	}

	// 生成JWT令牌
	token, err := utils.GenerateToken(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// 将令牌存储到Redis（可选，用于令牌撤销）
	cacheKey := "token:" + token
	err = database.SetCache(cacheKey, user.ID, time.Duration(24)*time.Hour)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"user":    user.ToResponse(),
	})
}

// 用户注册
func (h *AuthHandler) Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	// 检查用户名是否已存在
	var existingUser models.User
	if err := database.DB.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	// 检查邮箱是否已存在
	if err := database.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	// 加密密码
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process password"})
		return
	}

	// 创建新用户
	newUser := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     models.RoleUser,
		IsActive: true,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    newUser.ToResponse(),
	})
}

// 用户登出
func (h *AuthHandler) Logout(c *gin.Context) {
	// 从请求头获取令牌
	authHeader := c.GetHeader("Authorization")
	if authHeader != "" {
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) == 2 && parts[0] == "Bearer" {
			token := parts[1]
			// 将令牌加入黑名单（存储在Redis中）
			cacheKey := "blacklist:" + token
			err := database.SetCache(cacheKey, "revoked", time.Duration(24)*time.Hour)
			if err != nil {
				return
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

// 获取当前用户信息
func (h *AuthHandler) GetProfile(c *gin.Context) {
	user := middleware.GetCurrentUser(c)
	c.JSON(http.StatusOK, gin.H{
		"user": user.ToResponse(),
	})
}

// 更新用户信息
func (h *AuthHandler) UpdateProfile(c *gin.Context) {
	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	user := middleware.GetCurrentUser(c)
	updates := make(map[string]interface{})

	if req.Email != "" {
		// 检查邮箱是否已被其他用户使用
		var existingUser models.User
		if err := database.DB.Where("email = ? AND id != ?", req.Email, user.ID).First(&existingUser).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
			return
		}
		updates["email"] = req.Email
	}

	if req.Password != "" {
		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process password"})
			return
		}
		updates["password"] = hashedPassword
	}

	if req.Role != "" {
		// 只有管理员可以更改角色
		currentRole := middleware.GetCurrentUserRole(c)
		if currentRole != models.RoleAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can change roles"})
			return
		}
		updates["role"] = req.Role
	}

	if len(updates) > 0 {
		if err := database.DB.Model(&user).Updates(updates).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}
	}

	// 清除用户缓存
	cacheKey := "user:" + strconv.Itoa(int(user.ID))
	err := database.DeleteCache(cacheKey)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile updated successfully",
		"user":    user.ToResponse(),
	})
}

// 刷新令牌
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	user := middleware.GetCurrentUser(c)

	// 生成新的令牌
	newToken, err := utils.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate new token"})
		return
	}

	// 将新令牌存储到Redis
	cacheKey := "token:" + newToken
	err = database.SetCache(cacheKey, user.ID, time.Duration(24)*time.Hour)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Token refreshed successfully",
		"token":   newToken,
	})
}

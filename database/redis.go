package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"gin-auth-project/config"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis() {
	cfg := config.AppConfig

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	log.Println("Successfully connected to Redis")
}

// 设置缓存
func SetCache(key string, value interface{}, expiration time.Duration) error {
	ctx := context.Background()
	return RedisClient.Set(ctx, key, value, expiration).Err()
}

// 获取缓存
func GetCache(key string) (string, error) {
	ctx := context.Background()
	return RedisClient.Get(ctx, key).Result()
}

// 删除缓存
func DeleteCache(key string) error {
	ctx := context.Background()
	return RedisClient.Del(ctx, key).Err()
}

// 检查键是否存在
func ExistsCache(key string) (bool, error) {
	ctx := context.Background()
	result, err := RedisClient.Exists(ctx, key).Result()
	return result > 0, err
}

// 设置过期时间
func ExpireCache(key string, expiration time.Duration) error {
	ctx := context.Background()
	return RedisClient.Expire(ctx, key, expiration).Err()
}

package repository

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConfig 是 MongoDB 配置结构
type MongoConfig struct {
	URI      string
	Database string
}

// MongoRepository 是 MongoDB 存储的基础实现
type MongoRepository struct {
	client   *mongo.Client
	database *mongo.Database
}

// NewMongoRepository 创建一个新的 MongoDB 存储实例
func NewMongoRepository(config MongoConfig) (*MongoRepository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.URI))
	if err != nil {
		return nil, fmt.Errorf("连接 MongoDB 失败: %v", err)
	}

	// 测试连接
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("MongoDB 连接测试失败: %v", err)
	}

	return &MongoRepository{
		client:   client,
		database: client.Database(config.Database),
	}, nil
}

// Close 关闭 MongoDB 连接
func (r *MongoRepository) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return r.client.Disconnect(ctx)
}

// GetCollection 获取指定的集合
func (r *MongoRepository) GetCollection(name string) *mongo.Collection {
	return r.database.Collection(name)
} 
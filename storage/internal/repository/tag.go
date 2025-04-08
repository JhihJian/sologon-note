package repository

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Tag 是标签的数据模型
type Tag struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string            `bson:"name"`
	Description string            `bson:"description"`
	CreatedAt   time.Time         `bson:"created_at"`
	UpdatedAt   time.Time         `bson:"updated_at"`
}

// TagRepository 是标签存储的具体实现
type TagRepository struct {
	collection *mongo.Collection
}

// NewTagRepository 创建一个新的标签存储实例
func NewTagRepository(repo *MongoRepository) *TagRepository {
	return &TagRepository{
		collection: repo.GetCollection("tags"),
	}
}

// Create 创建新标签
func (r *TagRepository) Create(ctx context.Context, tag *Tag) error {
	tag.CreatedAt = time.Now()
	tag.UpdatedAt = time.Now()

	result, err := r.collection.InsertOne(ctx, tag)
	if err != nil {
		return fmt.Errorf("创建标签失败: %v", err)
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		tag.ID = oid
	}

	return nil
}

// GetByName 根据名称获取标签
func (r *TagRepository) GetByName(ctx context.Context, name string) (*Tag, error) {
	var tag Tag
	err := r.collection.FindOne(ctx, bson.M{"name": name}).Decode(&tag)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("未找到名称为 %s 的标签", name)
		}
		return nil, fmt.Errorf("获取标签失败: %v", err)
	}
	return &tag, nil
}

// Update 更新标签
func (r *TagRepository) Update(ctx context.Context, tag *Tag) error {
	tag.UpdatedAt = time.Now()

	update := bson.M{
		"$set": bson.M{
			"description": tag.Description,
			"updated_at":  tag.UpdatedAt,
		},
	}

	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": tag.ID}, update)
	if err != nil {
		return fmt.Errorf("更新标签失败: %v", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("未找到要更新的标签")
	}

	return nil
}

// Delete 删除标签
func (r *TagRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("删除标签失败: %v", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("未找到要删除的标签")
	}

	return nil
}

// List 列出所有标签
func (r *TagRepository) List(ctx context.Context) ([]*Tag, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("获取标签列表失败: %v", err)
	}
	defer cursor.Close(ctx)

	var tags []*Tag
	if err = cursor.All(ctx, &tags); err != nil {
		return nil, fmt.Errorf("解析标签列表失败: %v", err)
	}

	return tags, nil
} 
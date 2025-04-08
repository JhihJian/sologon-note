package repository

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Note 是笔记的数据模型
type Note struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string            `bson:"title"`
	Content   string            `bson:"content"`
	Tags      []string          `bson:"tags"`
	CreatedAt time.Time         `bson:"created_at"`
	UpdatedAt time.Time         `bson:"updated_at"`
}

// NoteRepository 是笔记存储的具体实现
type NoteRepository struct {
	collection *mongo.Collection
}

// NewNoteRepository 创建一个新的笔记存储实例
func NewNoteRepository(repo *MongoRepository) *NoteRepository {
	return &NoteRepository{
		collection: repo.GetCollection("notes"),
	}
}

// Create 创建新笔记
func (r *NoteRepository) Create(ctx context.Context, note *Note) error {
	note.CreatedAt = time.Now()
	note.UpdatedAt = time.Now()

	result, err := r.collection.InsertOne(ctx, note)
	if err != nil {
		return fmt.Errorf("创建笔记失败: %v", err)
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		note.ID = oid
	}

	return nil
}

// GetByTitle 根据标题获取笔记
func (r *NoteRepository) GetByTitle(ctx context.Context, title string) (*Note, error) {
	var note Note
	err := r.collection.FindOne(ctx, bson.M{"title": title}).Decode(&note)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("未找到标题为 %s 的笔记", title)
		}
		return nil, fmt.Errorf("获取笔记失败: %v", err)
	}
	return &note, nil
}

// Update 更新笔记
func (r *NoteRepository) Update(ctx context.Context, note *Note) error {
	note.UpdatedAt = time.Now()

	update := bson.M{
		"$set": bson.M{
			"content":    note.Content,
			"tags":       note.Tags,
			"updated_at": note.UpdatedAt,
		},
	}

	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": note.ID}, update)
	if err != nil {
		return fmt.Errorf("更新笔记失败: %v", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("未找到要更新的笔记")
	}

	return nil
}

// Delete 删除笔记
func (r *NoteRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("删除笔记失败: %v", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("未找到要删除的笔记")
	}

	return nil
}

// List 列出所有笔记
func (r *NoteRepository) List(ctx context.Context) ([]*Note, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("获取笔记列表失败: %v", err)
	}
	defer cursor.Close(ctx)

	var notes []*Note
	if err = cursor.All(ctx, &notes); err != nil {
		return nil, fmt.Errorf("解析笔记列表失败: %v", err)
	}

	return notes, nil
} 
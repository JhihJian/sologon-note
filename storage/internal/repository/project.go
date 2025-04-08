package repository

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Project 是项目的数据模型
type Project struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty"`
	Name        string              `bson:"name"`
	Description string              `bson:"description"`
	Tags        []string            `bson:"tags"`
	NoteIDs     []primitive.ObjectID `bson:"note_ids"`
	CreatedAt   time.Time           `bson:"created_at"`
	UpdatedAt   time.Time           `bson:"updated_at"`
}

// ProjectRepository 是项目存储的具体实现
type ProjectRepository struct {
	collection *mongo.Collection
}

// NewProjectRepository 创建一个新的项目存储实例
func NewProjectRepository(repo *MongoRepository) *ProjectRepository {
	return &ProjectRepository{
		collection: repo.GetCollection("projects"),
	}
}

// Create 创建新项目
func (r *ProjectRepository) Create(ctx context.Context, project *Project) error {
	project.CreatedAt = time.Now()
	project.UpdatedAt = time.Now()

	// 初始化空数组
	if project.NoteIDs == nil {
		project.NoteIDs = make([]primitive.ObjectID, 0)
	}
	if project.Tags == nil {
		project.Tags = make([]string, 0)
	}

	result, err := r.collection.InsertOne(ctx, project)
	if err != nil {
		return fmt.Errorf("创建项目失败: %v", err)
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		project.ID = oid
	}

	return nil
}

// GetByName 根据名称获取项目
func (r *ProjectRepository) GetByName(ctx context.Context, name string) (*Project, error) {
	var project Project
	err := r.collection.FindOne(ctx, bson.M{"name": name}).Decode(&project)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("未找到名称为 %s 的项目", name)
		}
		return nil, fmt.Errorf("获取项目失败: %v", err)
	}
	return &project, nil
}

// Update 更新项目
func (r *ProjectRepository) Update(ctx context.Context, project *Project) error {
	project.UpdatedAt = time.Now()

	update := bson.M{
		"$set": bson.M{
			"description": project.Description,
			"tags":        project.Tags,
			"note_ids":    project.NoteIDs,
			"updated_at":  project.UpdatedAt,
		},
	}

	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": project.ID}, update)
	if err != nil {
		return fmt.Errorf("更新项目失败: %v", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("未找到要更新的项目")
	}

	return nil
}

// Delete 删除项目
func (r *ProjectRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("删除项目失败: %v", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("未找到要删除的项目")
	}

	return nil
}

// List 列出所有项目
func (r *ProjectRepository) List(ctx context.Context) ([]*Project, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("获取项目列表失败: %v", err)
	}
	defer cursor.Close(ctx)

	var projects []*Project
	if err = cursor.All(ctx, &projects); err != nil {
		return nil, fmt.Errorf("解析项目列表失败: %v", err)
	}

	return projects, nil
}

// AddNote 向项目添加笔记
func (r *ProjectRepository) AddNote(ctx context.Context, projectID primitive.ObjectID, noteID primitive.ObjectID) error {
	update := bson.M{
		"$addToSet": bson.M{
			"note_ids": noteID,
		},
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": projectID}, update)
	if err != nil {
		return fmt.Errorf("向项目添加笔记失败: %v", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("未找到要更新的项目")
	}

	return nil
}

// RemoveNote 从项目中移除笔记
func (r *ProjectRepository) RemoveNote(ctx context.Context, projectID primitive.ObjectID, noteID primitive.ObjectID) error {
	update := bson.M{
		"$pull": bson.M{
			"note_ids": noteID,
		},
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": projectID}, update)
	if err != nil {
		return fmt.Errorf("从项目移除笔记失败: %v", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("未找到要更新的项目")
	}

	return nil
}

// GetNotes 获取项目中的所有笔记
func (r *ProjectRepository) GetNotes(ctx context.Context, projectID primitive.ObjectID) ([]primitive.ObjectID, error) {
	var project Project
	err := r.collection.FindOne(ctx, bson.M{"_id": projectID}).Decode(&project)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("未找到项目")
		}
		return nil, fmt.Errorf("获取项目笔记失败: %v", err)
	}

	return project.NoteIDs, nil
} 
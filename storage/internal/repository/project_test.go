package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func setupTestProjectRepository(t *testing.T) (*ProjectRepository, func()) {
	// 创建测试配置
	config := MongoConfig{
		URI:      "mongodb://localhost:27017",
		Database: "sologon_test",
	}

	// 创建 MongoDB 存储实例
	repo, err := NewMongoRepository(config)
	require.NoError(t, err, "创建 MongoDB 存储实例失败")

	// 创建项目存储实例
	projectRepo := NewProjectRepository(repo)

	// 返回清理函数
	cleanup := func() {
		// 删除测试数据库
		ctx := context.Background()
		err := repo.database.Drop(ctx)
		require.NoError(t, err, "清理测试数据库失败")
		repo.Close()
	}

	return projectRepo, cleanup
}

func TestProjectRepository_Create(t *testing.T) {
	repo, cleanup := setupTestProjectRepository(t)
	defer cleanup()

	ctx := context.Background()

	// 测试用例：创建新项目
	project := &Project{
		Name:        "测试项目",
		Description: "这是一个测试项目",
		Tags:        []string{"测试", "项目"},
	}

	err := repo.Create(ctx, project)
	require.NoError(t, err, "创建项目失败")
	assert.NotEmpty(t, project.ID, "项目 ID 不应为空")
	assert.NotZero(t, project.CreatedAt, "创建时间不应为零")
	assert.NotZero(t, project.UpdatedAt, "更新时间不应为零")

	// 验证项目是否成功创建
	retrievedProject, err := repo.GetByName(ctx, "测试项目")
	require.NoError(t, err, "获取项目失败")
	assert.Equal(t, project.Name, retrievedProject.Name)
	assert.Equal(t, project.Description, retrievedProject.Description)
	assert.Equal(t, project.Tags, retrievedProject.Tags)
}

func TestProjectRepository_GetByName(t *testing.T) {
	repo, cleanup := setupTestProjectRepository(t)
	defer cleanup()

	ctx := context.Background()

	// 准备测试数据
	project := &Project{
		Name:        "测试项目",
		Description: "这是一个测试项目",
		Tags:        []string{"测试", "项目"},
	}
	err := repo.Create(ctx, project)
	require.NoError(t, err)

	// 测试用例：获取存在的项目
	retrievedProject, err := repo.GetByName(ctx, "测试项目")
	require.NoError(t, err)
	assert.Equal(t, project.Name, retrievedProject.Name)
	assert.Equal(t, project.Description, retrievedProject.Description)
	assert.Equal(t, project.Tags, retrievedProject.Tags)

	// 测试用例：获取不存在的项目
	_, err = repo.GetByName(ctx, "不存在的项目")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "未找到名称为 不存在的项目 的项目")
}

func TestProjectRepository_Update(t *testing.T) {
	repo, cleanup := setupTestProjectRepository(t)
	defer cleanup()

	ctx := context.Background()

	// 准备测试数据
	project := &Project{
		Name:        "测试项目",
		Description: "这是一个测试项目",
		Tags:        []string{"测试", "项目"},
	}
	err := repo.Create(ctx, project)
	require.NoError(t, err)

	// 测试用例：更新项目
	project.Description = "更新后的描述"
	project.Tags = []string{"测试", "更新"}
	err = repo.Update(ctx, project)
	require.NoError(t, err)

	// 验证更新结果
	retrievedProject, err := repo.GetByName(ctx, "测试项目")
	require.NoError(t, err)
	assert.Equal(t, "更新后的描述", retrievedProject.Description)
	assert.Equal(t, []string{"测试", "更新"}, retrievedProject.Tags)

	// 测试用例：更新不存在的项目
	nonExistentProject := &Project{
		ID:          primitive.NewObjectID(),
		Name:        "不存在的项目",
		Description: "描述",
	}
	err = repo.Update(ctx, nonExistentProject)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "未找到要更新的项目")
}

func TestProjectRepository_Delete(t *testing.T) {
	repo, cleanup := setupTestProjectRepository(t)
	defer cleanup()

	ctx := context.Background()

	// 准备测试数据
	project := &Project{
		Name:        "测试项目",
		Description: "这是一个测试项目",
		Tags:        []string{"测试", "项目"},
	}
	err := repo.Create(ctx, project)
	require.NoError(t, err)

	// 测试用例：删除存在的项目
	err = repo.Delete(ctx, project.ID)
	require.NoError(t, err)

	// 验证项目是否被删除
	_, err = repo.GetByName(ctx, "测试项目")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "未找到名称为 测试项目 的项目")

	// 测试用例：删除不存在的项目
	err = repo.Delete(ctx, primitive.NewObjectID())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "未找到要删除的项目")
}

func TestProjectRepository_List(t *testing.T) {
	repo, cleanup := setupTestProjectRepository(t)
	defer cleanup()

	ctx := context.Background()

	// 准备测试数据
	projects := []*Project{
		{
			Name:        "项目1",
			Description: "描述1",
			Tags:        []string{"测试1"},
		},
		{
			Name:        "项目2",
			Description: "描述2",
			Tags:        []string{"测试2"},
		},
	}

	for _, project := range projects {
		err := repo.Create(ctx, project)
		require.NoError(t, err)
	}

	// 测试用例：获取所有项目
	retrievedProjects, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, retrievedProjects, 2)

	// 验证项目内容
	for i, project := range retrievedProjects {
		assert.Equal(t, projects[i].Name, project.Name)
		assert.Equal(t, projects[i].Description, project.Description)
		assert.Equal(t, projects[i].Tags, project.Tags)
	}
}

func TestProjectRepository_AddNote(t *testing.T) {
	repo, cleanup := setupTestProjectRepository(t)
	defer cleanup()

	ctx := context.Background()

	// 准备测试数据
	project := &Project{
		Name:        "测试项目",
		Description: "这是一个测试项目",
		Tags:        []string{"测试", "项目"},
	}
	err := repo.Create(ctx, project)
	require.NoError(t, err)

	// 测试用例：添加笔记
	noteID := primitive.NewObjectID()
	err = repo.AddNote(ctx, project.ID, noteID)
	require.NoError(t, err)

	// 验证笔记是否被添加
	notes, err := repo.GetNotes(ctx, project.ID)
	require.NoError(t, err)
	assert.Contains(t, notes, noteID)

	// 测试用例：添加到不存在的项目
	err = repo.AddNote(ctx, primitive.NewObjectID(), noteID)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "未找到要更新的项目")
}

func TestProjectRepository_RemoveNote(t *testing.T) {
	repo, cleanup := setupTestProjectRepository(t)
	defer cleanup()

	ctx := context.Background()

	// 准备测试数据
	project := &Project{
		Name:        "测试项目",
		Description: "这是一个测试项目",
		Tags:        []string{"测试", "项目"},
	}
	err := repo.Create(ctx, project)
	require.NoError(t, err)

	noteID := primitive.NewObjectID()
	err = repo.AddNote(ctx, project.ID, noteID)
	require.NoError(t, err)

	// 测试用例：移除笔记
	err = repo.RemoveNote(ctx, project.ID, noteID)
	require.NoError(t, err)

	// 验证笔记是否被移除
	notes, err := repo.GetNotes(ctx, project.ID)
	require.NoError(t, err)
	assert.NotContains(t, notes, noteID)

	// 测试用例：从不存在的项目移除笔记
	err = repo.RemoveNote(ctx, primitive.NewObjectID(), noteID)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "未找到要更新的项目")
}

func TestProjectRepository_GetNotes(t *testing.T) {
	repo, cleanup := setupTestProjectRepository(t)
	defer cleanup()

	ctx := context.Background()

	// 准备测试数据
	project := &Project{
		Name:        "测试项目",
		Description: "这是一个测试项目",
		Tags:        []string{"测试", "项目"},
	}
	err := repo.Create(ctx, project)
	require.NoError(t, err)

	noteIDs := []primitive.ObjectID{
		primitive.NewObjectID(),
		primitive.NewObjectID(),
	}

	for _, noteID := range noteIDs {
		err = repo.AddNote(ctx, project.ID, noteID)
		require.NoError(t, err)
	}

	// 测试用例：获取项目笔记
	notes, err := repo.GetNotes(ctx, project.ID)
	require.NoError(t, err)
	assert.Len(t, notes, 2)
	for _, noteID := range noteIDs {
		assert.Contains(t, notes, noteID)
	}

	// 测试用例：获取不存在项目的笔记
	_, err = repo.GetNotes(ctx, primitive.NewObjectID())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "未找到项目")
} 
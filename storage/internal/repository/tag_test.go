package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func setupTestTagRepository(t *testing.T) (*TagRepository, func()) {
	// 创建测试配置
	config := MongoConfig{
		URI:      "mongodb://localhost:27017",
		Database: "sologon_test",
	}

	// 创建 MongoDB 存储实例
	repo, err := NewMongoRepository(config)
	require.NoError(t, err, "创建 MongoDB 存储实例失败")

	// 创建标签存储实例
	tagRepo := NewTagRepository(repo)

	// 返回清理函数
	cleanup := func() {
		// 删除测试数据库
		ctx := context.Background()
		err := repo.database.Drop(ctx)
		require.NoError(t, err, "清理测试数据库失败")
		repo.Close()
	}

	return tagRepo, cleanup
}

func TestTagRepository_Create(t *testing.T) {
	repo, cleanup := setupTestTagRepository(t)
	defer cleanup()

	ctx := context.Background()

	// 测试用例：创建新标签
	tag := &Tag{
		Name:        "测试标签",
		Description: "这是一个测试标签",
	}

	err := repo.Create(ctx, tag)
	require.NoError(t, err, "创建标签失败")
	assert.NotEmpty(t, tag.ID, "标签 ID 不应为空")
	assert.NotZero(t, tag.CreatedAt, "创建时间不应为零")
	assert.NotZero(t, tag.UpdatedAt, "更新时间不应为零")

	// 验证标签是否成功创建
	retrievedTag, err := repo.GetByName(ctx, "测试标签")
	require.NoError(t, err, "获取标签失败")
	assert.Equal(t, tag.Name, retrievedTag.Name)
	assert.Equal(t, tag.Description, retrievedTag.Description)
}

func TestTagRepository_GetByName(t *testing.T) {
	repo, cleanup := setupTestTagRepository(t)
	defer cleanup()

	ctx := context.Background()

	// 准备测试数据
	tag := &Tag{
		Name:        "测试标签",
		Description: "这是一个测试标签",
	}
	err := repo.Create(ctx, tag)
	require.NoError(t, err)

	// 测试用例：获取存在的标签
	retrievedTag, err := repo.GetByName(ctx, "测试标签")
	require.NoError(t, err)
	assert.Equal(t, tag.Name, retrievedTag.Name)
	assert.Equal(t, tag.Description, retrievedTag.Description)

	// 测试用例：获取不存在的标签
	_, err = repo.GetByName(ctx, "不存在的标签")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "未找到名称为 不存在的标签 的标签")
}

func TestTagRepository_Update(t *testing.T) {
	repo, cleanup := setupTestTagRepository(t)
	defer cleanup()

	ctx := context.Background()

	// 准备测试数据
	tag := &Tag{
		Name:        "测试标签",
		Description: "这是一个测试标签",
	}
	err := repo.Create(ctx, tag)
	require.NoError(t, err)

	// 测试用例：更新标签
	tag.Description = "更新后的描述"
	err = repo.Update(ctx, tag)
	require.NoError(t, err)

	// 验证更新结果
	retrievedTag, err := repo.GetByName(ctx, "测试标签")
	require.NoError(t, err)
	assert.Equal(t, "更新后的描述", retrievedTag.Description)

	// 测试用例：更新不存在的标签
	nonExistentTag := &Tag{
		ID:          primitive.NewObjectID(),
		Name:        "不存在的标签",
		Description: "描述",
	}
	err = repo.Update(ctx, nonExistentTag)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "未找到要更新的标签")
}

func TestTagRepository_Delete(t *testing.T) {
	repo, cleanup := setupTestTagRepository(t)
	defer cleanup()

	ctx := context.Background()

	// 准备测试数据
	tag := &Tag{
		Name:        "测试标签",
		Description: "这是一个测试标签",
	}
	err := repo.Create(ctx, tag)
	require.NoError(t, err)

	// 测试用例：删除存在的标签
	err = repo.Delete(ctx, tag.ID)
	require.NoError(t, err)

	// 验证标签是否被删除
	_, err = repo.GetByName(ctx, "测试标签")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "未找到名称为 测试标签 的标签")

	// 测试用例：删除不存在的标签
	err = repo.Delete(ctx, primitive.NewObjectID())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "未找到要删除的标签")
}

func TestTagRepository_List(t *testing.T) {
	repo, cleanup := setupTestTagRepository(t)
	defer cleanup()

	ctx := context.Background()

	// 准备测试数据
	tags := []*Tag{
		{
			Name:        "标签1",
			Description: "描述1",
		},
		{
			Name:        "标签2",
			Description: "描述2",
		},
	}

	for _, tag := range tags {
		err := repo.Create(ctx, tag)
		require.NoError(t, err)
	}

	// 测试用例：获取所有标签
	retrievedTags, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, retrievedTags, 2)

	// 验证标签内容
	for i, tag := range retrievedTags {
		assert.Equal(t, tags[i].Name, tag.Name)
		assert.Equal(t, tags[i].Description, tag.Description)
	}
} 
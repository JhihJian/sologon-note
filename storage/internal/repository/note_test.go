package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func setupTestNoteRepository(t *testing.T) (*NoteRepository, func()) {
	// 创建测试配置
	config := MongoConfig{
		URI:      "mongodb://localhost:27017",
		Database: "sologon_test",
	}

	// 创建 MongoDB 存储实例
	repo, err := NewMongoRepository(config)
	require.NoError(t, err, "创建 MongoDB 存储实例失败")

	// 创建笔记存储实例
	noteRepo := NewNoteRepository(repo)

	// 返回清理函数
	cleanup := func() {
		// 删除测试数据库
		ctx := context.Background()
		err := repo.database.Drop(ctx)
		require.NoError(t, err, "清理测试数据库失败")
		repo.Close()
	}

	return noteRepo, cleanup
}

func TestNoteRepository_Create(t *testing.T) {
	repo, cleanup := setupTestNoteRepository(t)
	defer cleanup()

	ctx := context.Background()

	// 测试用例：创建新笔记
	note := &Note{
		Title:   "测试笔记",
		Content: "这是测试内容",
		Tags:    []string{"测试", "单元测试"},
	}

	err := repo.Create(ctx, note)
	require.NoError(t, err, "创建笔记失败")
	assert.NotEmpty(t, note.ID, "笔记 ID 不应为空")
	assert.NotZero(t, note.CreatedAt, "创建时间不应为零")
	assert.NotZero(t, note.UpdatedAt, "更新时间不应为零")

	// 验证笔记是否成功创建
	retrievedNote, err := repo.GetByTitle(ctx, "测试笔记")
	require.NoError(t, err, "获取笔记失败")
	assert.Equal(t, note.Title, retrievedNote.Title)
	assert.Equal(t, note.Content, retrievedNote.Content)
	assert.Equal(t, note.Tags, retrievedNote.Tags)
}

func TestNoteRepository_GetByTitle(t *testing.T) {
	repo, cleanup := setupTestNoteRepository(t)
	defer cleanup()

	ctx := context.Background()

	// 准备测试数据
	note := &Note{
		Title:   "测试笔记",
		Content: "这是测试内容",
		Tags:    []string{"测试", "单元测试"},
	}
	err := repo.Create(ctx, note)
	require.NoError(t, err)

	// 测试用例：获取存在的笔记
	retrievedNote, err := repo.GetByTitle(ctx, "测试笔记")
	require.NoError(t, err)
	assert.Equal(t, note.Title, retrievedNote.Title)
	assert.Equal(t, note.Content, retrievedNote.Content)
	assert.Equal(t, note.Tags, retrievedNote.Tags)

	// 测试用例：获取不存在的笔记
	_, err = repo.GetByTitle(ctx, "不存在的笔记")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "未找到标题为 不存在的笔记 的笔记")
}

func TestNoteRepository_Update(t *testing.T) {
	repo, cleanup := setupTestNoteRepository(t)
	defer cleanup()

	ctx := context.Background()

	// 准备测试数据
	note := &Note{
		Title:   "测试笔记",
		Content: "这是测试内容",
		Tags:    []string{"测试", "单元测试"},
	}
	err := repo.Create(ctx, note)
	require.NoError(t, err)

	// 测试用例：更新笔记
	note.Content = "更新后的内容"
	note.Tags = []string{"测试", "更新"}
	err = repo.Update(ctx, note)
	require.NoError(t, err)

	// 验证更新结果
	retrievedNote, err := repo.GetByTitle(ctx, "测试笔记")
	require.NoError(t, err)
	assert.Equal(t, "更新后的内容", retrievedNote.Content)
	assert.Equal(t, []string{"测试", "更新"}, retrievedNote.Tags)

	// 测试用例：更新不存在的笔记
	nonExistentNote := &Note{
		ID:      primitive.NewObjectID(),
		Title:   "不存在的笔记",
		Content: "内容",
	}
	err = repo.Update(ctx, nonExistentNote)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "未找到要更新的笔记")
}

func TestNoteRepository_Delete(t *testing.T) {
	repo, cleanup := setupTestNoteRepository(t)
	defer cleanup()

	ctx := context.Background()

	// 准备测试数据
	note := &Note{
		Title:   "测试笔记",
		Content: "这是测试内容",
		Tags:    []string{"测试", "单元测试"},
	}
	err := repo.Create(ctx, note)
	require.NoError(t, err)

	// 测试用例：删除存在的笔记
	err = repo.Delete(ctx, note.ID)
	require.NoError(t, err)

	// 验证笔记是否被删除
	_, err = repo.GetByTitle(ctx, "测试笔记")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "未找到标题为 测试笔记 的笔记")

	// 测试用例：删除不存在的笔记
	err = repo.Delete(ctx, primitive.NewObjectID())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "未找到要删除的笔记")
}

func TestNoteRepository_List(t *testing.T) {
	repo, cleanup := setupTestNoteRepository(t)
	defer cleanup()

	ctx := context.Background()

	// 准备测试数据
	notes := []*Note{
		{
			Title:   "笔记1",
			Content: "内容1",
			Tags:    []string{"测试"},
		},
		{
			Title:   "笔记2",
			Content: "内容2",
			Tags:    []string{"测试"},
		},
	}

	for _, note := range notes {
		err := repo.Create(ctx, note)
		require.NoError(t, err)
	}

	// 测试用例：获取所有笔记
	retrievedNotes, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, retrievedNotes, 2)

	// 验证笔记内容
	for i, note := range retrievedNotes {
		assert.Equal(t, notes[i].Title, note.Title)
		assert.Equal(t, notes[i].Content, note.Content)
		assert.Equal(t, notes[i].Tags, note.Tags)
	}
} 
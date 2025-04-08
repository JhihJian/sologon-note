package main

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/sologon/storage/internal/tools"
)

func main() {
	// 创建 MCP 服务器
	s := server.NewMCPServer(
		"Sologon Storage",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)

	// 添加笔记管理工具
	noteTool := mcp.NewTool("note",
		mcp.WithDescription("管理笔记"),
		mcp.WithString("action",
			mcp.Required(),
			mcp.Description("操作类型 (create, read, update, delete)"),
			mcp.Enum("create", "read", "update", "delete"),
		),
		mcp.WithString("title",
			mcp.Description("笔记标题"),
		),
		mcp.WithString("content",
			mcp.Description("笔记内容"),
		),
	)

	// 添加笔记工具处理器
	s.AddTool(noteTool, tools.NoteHandler)

	// 添加标签管理工具
	tagTool := mcp.NewTool("tag",
		mcp.WithDescription("管理标签"),
		mcp.WithString("action",
			mcp.Required(),
			mcp.Description("操作类型 (create, read, update, delete)"),
			mcp.Enum("create", "read", "update", "delete"),
		),
		mcp.WithString("name",
			mcp.Description("标签名称"),
		),
	)

	// 添加标签工具处理器
	s.AddTool(tagTool, tools.TagHandler)

	// 添加项目管理工具
	projectTool := mcp.NewTool("project",
		mcp.WithDescription("管理项目"),
		mcp.WithString("action",
			mcp.Required(),
			mcp.Description("操作类型 (create, read, update, delete)"),
			mcp.Enum("create", "read", "update", "delete"),
		),
		mcp.WithString("name",
			mcp.Description("项目名称"),
		),
		mcp.WithString("description",
			mcp.Description("项目描述"),
		),
	)

	// 添加项目工具处理器
	s.AddTool(projectTool, tools.ProjectHandler)

	// 启动服务器
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("服务器错误: %v\n", err)
	}
} 
package test

import (
	"context"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/sologon/storage/internal/tools"
)

func TestNoteHandler(t *testing.T) {
	tests := []struct {
		name    string
		request mcp.CallToolRequest
		want    string
		wantErr bool
	}{
		{
			name: "创建笔记",
			request: mcp.CallToolRequest{
				Params: struct {
					Name      string                 `json:"name"`
					Arguments map[string]interface{} `json:"arguments,omitempty"`
					Meta      *struct {
						ProgressToken mcp.ProgressToken `json:"progressToken,omitempty"`
					} `json:"_meta,omitempty"`
				}{
					Arguments: map[string]interface{}{
						"action":  "create",
						"title":   "测试笔记",
						"content": "测试内容",
					},
				},
			},
			want:    "创建笔记: 测试笔记, 内容: 测试内容",
			wantErr: false,
		},
		{
			name: "读取笔记",
			request: mcp.CallToolRequest{
				Params: struct {
					Name      string                 `json:"name"`
					Arguments map[string]interface{} `json:"arguments,omitempty"`
					Meta      *struct {
						ProgressToken mcp.ProgressToken `json:"progressToken,omitempty"`
					} `json:"_meta,omitempty"`
				}{
					Arguments: map[string]interface{}{
						"action": "read",
						"title":  "测试笔记",
					},
				},
			},
			want:    "读取笔记: 测试笔记",
			wantErr: false,
		},
		{
			name: "更新笔记",
			request: mcp.CallToolRequest{
				Params: struct {
					Name      string                 `json:"name"`
					Arguments map[string]interface{} `json:"arguments,omitempty"`
					Meta      *struct {
						ProgressToken mcp.ProgressToken `json:"progressToken,omitempty"`
					} `json:"_meta,omitempty"`
				}{
					Arguments: map[string]interface{}{
						"action":  "update",
						"title":   "测试笔记",
						"content": "更新内容",
					},
				},
			},
			want:    "更新笔记: 测试笔记, 新内容: 更新内容",
			wantErr: false,
		},
		{
			name: "删除笔记",
			request: mcp.CallToolRequest{
				Params: struct {
					Name      string                 `json:"name"`
					Arguments map[string]interface{} `json:"arguments,omitempty"`
					Meta      *struct {
						ProgressToken mcp.ProgressToken `json:"progressToken,omitempty"`
					} `json:"_meta,omitempty"`
				}{
					Arguments: map[string]interface{}{
						"action": "delete",
						"title":  "测试笔记",
					},
				},
			},
			want:    "删除笔记: 测试笔记",
			wantErr: false,
		},
		{
			name: "缺少 action 参数",
			request: mcp.CallToolRequest{
				Params: struct {
					Name      string                 `json:"name"`
					Arguments map[string]interface{} `json:"arguments,omitempty"`
					Meta      *struct {
						ProgressToken mcp.ProgressToken `json:"progressToken,omitempty"`
					} `json:"_meta,omitempty"`
				}{
					Arguments: map[string]interface{}{
						"title": "测试笔记",
					},
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "不支持的操作",
			request: mcp.CallToolRequest{
				Params: struct {
					Name      string                 `json:"name"`
					Arguments map[string]interface{} `json:"arguments,omitempty"`
					Meta      *struct {
						ProgressToken mcp.ProgressToken `json:"progressToken,omitempty"`
					} `json:"_meta,omitempty"`
				}{
					Arguments: map[string]interface{}{
						"action": "invalid",
						"title":  "测试笔记",
					},
				},
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tools.NoteHandler(context.Background(), tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("NoteHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.Content[0].(mcp.TextContent).Text != tt.want {
				t.Errorf("NoteHandler() = %v, want %v", got.Content[0].(mcp.TextContent).Text, tt.want)
			}
		})
	}
} 
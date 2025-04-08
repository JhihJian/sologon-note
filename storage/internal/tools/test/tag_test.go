package test

import (
	"context"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/sologon/storage/internal/tools"
)

func TestTagHandler(t *testing.T) {
	tests := []struct {
		name    string
		request mcp.CallToolRequest
		want    string
		wantErr bool
	}{
		{
			name: "创建标签",
			request: mcp.CallToolRequest{
				Params: struct {
					Name      string                 `json:"name"`
					Arguments map[string]interface{} `json:"arguments,omitempty"`
					Meta      *struct {
						ProgressToken mcp.ProgressToken `json:"progressToken,omitempty"`
					} `json:"_meta,omitempty"`
				}{
					Arguments: map[string]interface{}{
						"action": "create",
						"name":   "测试标签",
					},
				},
			},
			want:    "创建标签: 测试标签",
			wantErr: false,
		},
		{
			name: "读取标签",
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
						"name":   "测试标签",
					},
				},
			},
			want:    "读取标签: 测试标签",
			wantErr: false,
		},
		{
			name: "更新标签",
			request: mcp.CallToolRequest{
				Params: struct {
					Name      string                 `json:"name"`
					Arguments map[string]interface{} `json:"arguments,omitempty"`
					Meta      *struct {
						ProgressToken mcp.ProgressToken `json:"progressToken,omitempty"`
					} `json:"_meta,omitempty"`
				}{
					Arguments: map[string]interface{}{
						"action": "update",
						"name":   "测试标签",
					},
				},
			},
			want:    "更新标签: 测试标签",
			wantErr: false,
		},
		{
			name: "删除标签",
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
						"name":   "测试标签",
					},
				},
			},
			want:    "删除标签: 测试标签",
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
						"name": "测试标签",
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
						"name":   "测试标签",
					},
				},
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tools.TagHandler(context.Background(), tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("TagHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.Content[0].(mcp.TextContent).Text != tt.want {
				t.Errorf("TagHandler() = %v, want %v", got.Content[0].(mcp.TextContent).Text, tt.want)
			}
		})
	}
} 
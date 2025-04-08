package mcp

// Request 表示 MCP 协议的请求结构
type Request struct {
	// 工具名称
	Tool string `json:"tool"`
	// 工具版本
	Version string `json:"version"`
	// 请求参数
	Params map[string]interface{} `json:"params"`
}

// Response 表示 MCP 协议的响应结构
type Response struct {
	// 是否成功
	Success bool `json:"success"`
	// 错误信息
	Error string `json:"error,omitempty"`
	// 响应数据
	Data interface{} `json:"data,omitempty"`
}

// Tool 表示 MCP 工具接口
type Tool interface {
	// 获取工具名称
	Name() string
	// 获取工具版本
	Version() string
	// 执行工具
	Execute(params map[string]interface{}) (interface{}, error)
} 
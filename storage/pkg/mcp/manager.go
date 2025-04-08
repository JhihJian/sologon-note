package mcp

import (
	"fmt"
	"sync"
)

// Manager 管理 MCP 工具
type Manager struct {
	tools map[string]Tool
	mu    sync.RWMutex
}

// NewManager 创建新的工具管理器
func NewManager() *Manager {
	return &Manager{
		tools: make(map[string]Tool),
	}
}

// Register 注册新工具
func (m *Manager) Register(tool Tool) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	name := tool.Name()
	if _, exists := m.tools[name]; exists {
		return fmt.Errorf("tool %s already registered", name)
	}

	m.tools[name] = tool
	return nil
}

// GetTool 获取指定名称的工具
func (m *Manager) GetTool(name string) (Tool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	tool, exists := m.tools[name]
	if !exists {
		return nil, fmt.Errorf("tool %s not found", name)
	}

	return tool, nil
}

// Execute 执行指定工具
func (m *Manager) Execute(req *Request) (*Response, error) {
	tool, err := m.GetTool(req.Tool)
	if err != nil {
		return &Response{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	data, err := tool.Execute(req.Params)
	if err != nil {
		return &Response{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &Response{
		Success: true,
		Data:    data,
	}, nil
} 
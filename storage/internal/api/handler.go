package api

import (
	"encoding/json"
	"net/http"

	"github.com/sologon/storage/pkg/mcp"
)

// Handler 处理 HTTP 请求
type Handler struct {
	manager *mcp.Manager
}

// NewHandler 创建新的处理器
func NewHandler(manager *mcp.Manager) *Handler {
	return &Handler{
		manager: manager,
	}
}

// HandleRequest 处理 MCP 请求
func (h *Handler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req mcp.Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.manager.Execute(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
} 
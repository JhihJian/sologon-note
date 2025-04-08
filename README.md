# Sologon - 智能知识入库助手

![Sologon 宣传图](assert/images/sologon.svg)

Sologon 是一款智能知识入库助手，帮助用户以极低的成本记录、管理和同步笔记、灵感、任务等碎片信息。

## 系统架构

![项目结构图](assert/images/project-structure.svg)

### 前端架构 (React)
- **组件化设计**：采用 React 组件化开发，实现高内聚低耦合
- **状态管理**：使用 React Hooks 管理组件状态
- **路由管理**：使用 React Router 实现页面路由
- **UI 框架**：基于 Material-UI 构建现代化界面
- **API 集成**：通过 Axios 与后端服务通信

### 后端架构 (Go)
- **模块化设计**：采用 Go 标准项目布局
- **API 设计**：RESTful API 设计规范
- **数据库**：使用 SQLite 存储数据
- **认证授权**：基于 JWT 的认证机制
- **文件存储**：本地文件系统存储

### Chrome 扩展
- **内容抓取**：智能识别网页内容
- **快速记录**：一键保存到知识库
- **标签管理**：自动生成标签建议
- **项目关联**：智能关联相关项目

## 主要功能

- 多端快速记录（Chrome插件、Android App、网页端）
- 四大记录入口（我干了、我想到了、提醒我、做记录）
- 智能内容结构化（自动标题、分类、标签）
- 项目化管理系统
- GitHub 自动同步

## 技术栈

- 前端：React + TypeScript
- 后端：Python + FastAPI
- 数据库：SQLite
- 部署：Docker

## 开发环境设置

### 后端设置

1. 创建虚拟环境：
```bash
cd backend
python3 -m venv venv
source venv/bin/activate
```

2. 安装依赖：
```bash
pip install -r requirements.txt
```

3. 运行后端服务：
```bash
uvicorn app.main:app --reload
```

### 前端设置

1. 安装依赖：
```bash
cd frontend
npm install
```

2. 运行开发服务器：
```bash
npm start
```

## API 文档

启动后端服务后，访问 http://localhost:8000/docs 查看完整的 API 文档。

## 项目结构

```
sologon/
├── backend/
│   ├── app/
│   │   ├── api/
│   │   ├── core/
│   │   ├── db/
│   │   ├── models/
│   │   ├── schemas/
│   │   └── services/
│   └── requirements.txt
├── frontend/
│   ├── public/
│   ├── src/
│   └── package.json
└── README.md
```

## 贡献指南

1. Fork 项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

## 许可证

MIT License 
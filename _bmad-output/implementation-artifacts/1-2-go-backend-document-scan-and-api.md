# Story 1.2: Go 后端文档扫描与 API

Status: ready-for-dev

## Story

As a 浏览者，
I want 应用启动时自动扫描文档目录并提供 API 访问，
so that 我可以通过浏览器获取文档数据。

## Acceptance Criteria

1. Go 服务启动时自动扫描 `_bmad-output/` 目录下所有 `.md` 文件（FR15）
2. 解析每个 Markdown 文件的 YAML frontmatter 元数据（FR16）
3. 根据目录结构自动推断文档所属的 BMAD 阶段（FR18）
4. 所有解析数据缓存到内存
5. `GET /api/documents/:path` 返回指定文档的 JSON 数据（path、title、content、frontmatter、phase）
6. 文档不存在时 `GET /api/documents/:path` 返回 404 `{"error": "document not found"}`
7. `GET /api/documents`（无 path 参数）返回所有文档列表（path、title、phase，不含 content）

## Tasks / Subtasks

- [ ] Task 1: 定义数据模型 (AC: #5, #7)
  - [ ] 创建 `server/model/document.go`，定义 Document struct（Path、Title、Content、Frontmatter、Phase）
  - [ ] JSON tag 使用 camelCase（如 `json:"path"`、`json:"frontmatter"`）
  - [ ] 创建 DocumentSummary struct（Path、Title、Phase，不含 Content），用于列表 API

- [ ] Task 2: 实现 Markdown 解析器 (AC: #2)
  - [ ] 创建 `server/parser/markdown_parser.go`
  - [ ] 实现 ParseMarkdown 函数：读取 .md 文件，分离 YAML frontmatter 和 Markdown 内容
  - [ ] frontmatter 解析为 `map[string]interface{}`
  - [ ] 从 frontmatter 或 Markdown 内容提取标题（优先 frontmatter 的 title 字段，否则取第一个 `# ` 标题）
  - [ ] 需要引入 YAML 解析库（`gopkg.in/yaml.v3`）

- [ ] Task 3: 实现文件系统扫描器 (AC: #1, #3, #4)
  - [ ] 创建 `server/parser/scanner.go`
  - [ ] 实现 ScanDocuments 函数：递归扫描指定目录下所有 `.md` 文件
  - [ ] 对每个文件调用 ParseMarkdown 解析
  - [ ] 根据目录结构推断 BMAD 阶段（如 `brainstorming/` → 分析阶段，`planning-artifacts/` → 规划阶段）
  - [ ] 返回 `[]Document` 缓存在内存中
  - [ ] 支持通过命令行参数指定扫描目录路径（默认 `../_bmad-output`）

- [ ] Task 4: 实现文档 API handler (AC: #5, #6, #7)
  - [ ] 创建 `server/handler/documents.go`
  - [ ] 实现 `GET /api/documents` handler：返回所有文档的 DocumentSummary 列表
  - [ ] 实现 `GET /api/documents/:path` handler：根据 path 参数返回完整 Document
  - [ ] 文档不存在时返回 404 `{"error": "document not found"}`
  - [ ] Content-Type 设置为 `application/json`
  - [ ] 注意：Go 标准库 `http.HandleFunc` 不直接支持路径参数，需用路径前缀匹配 + 手动提取

- [ ] Task 5: 集成到 main.go (AC: #1, #4)
  - [ ] 修改 `server/main.go`：启动时调用 ScanDocuments 扫描文档目录
  - [ ] 将扫描结果传递给 handler
  - [ ] 添加命令行参数解析（`-port`、`-dir`）
  - [ ] 保留 `/api/health` 端点
  - [ ] 日志输出扫描到的文件数量

- [ ] Task 6: 验证 (AC: #1-#7)
  - [ ] Go 编译成功
  - [ ] 启动后日志显示扫描到的文件数量
  - [ ] `curl /api/documents` 返回文档列表
  - [ ] `curl /api/documents/planning-artifacts/prd.md` 返回完整文档（含 content 和 frontmatter）
  - [ ] `curl /api/documents/nonexistent.md` 返回 404

## Dev Notes

### 前序故事 1.1 已建立的基础
- Go 项目在 `bmad-viewer/server/`，module 名 `bmad-viewer/server`
- `main.go` 已有最小 HTTP server + `/api/health` 端点
- 空目录 `handler/`、`parser/`、`model/` 已创建（有 .gitkeep）
- Go 标准库 HTTP server，无第三方路由框架

### 架构约束
- Go 文件名用 `snake_case`（如 `markdown_parser.go`、`documents.go`）
- API 响应直接返回数据，不加包装层
- JSON 字段用 `camelCase`（Go struct tag）
- 错误响应格式：`{"error": "message"}`
- 不引入架构文档未列出的第三方依赖（YAML 解析库 `gopkg.in/yaml.v3` 是必需的例外）

### 数据模型参考（来自架构文档）
```go
// model/document.go
type Document struct {
    Path        string                 `json:"path"`
    Title       string                 `json:"title"`
    Content     string                 `json:"content"`
    Frontmatter map[string]interface{} `json:"frontmatter"`
    Phase       string                 `json:"phase"`
}

type DocumentSummary struct {
    Path  string `json:"path"`
    Title string `json:"title"`
    Phase string `json:"phase"`
}
```

### Frontmatter 解析说明
Markdown 文件以 `---` 开头和结尾的 YAML 块为 frontmatter：
```markdown
---
key: value
list:
  - item1
  - item2
---

# Document Title

Content here...
```

### 阶段推断逻辑
根据文件路径推断 BMAD 阶段：
- `brainstorming/` → "analysis"（分析）
- `planning-artifacts/` → "planning"（规划）
- `implementation-artifacts/` → "implementation"（实施）
- 其他 → "other"

### Go 标准库路由注意事项
Go 1.22+ 的 `http.HandleFunc` 支持路径模式匹配：
```go
// Go 1.22+ 支持
http.HandleFunc("GET /api/documents/{path...}", handler)
```
如果 Go 版本支持，使用此方式。否则用 `strings.TrimPrefix` 手动提取路径参数。

### 文件结构
本故事创建/修改的文件：
```
server/
├── main.go              (modify: 添加扫描和路由)
├── go.mod               (modify: 添加 yaml.v3 依赖)
├── go.sum               (new: 依赖校验)
├── model/
│   └── document.go      (new: 数据模型)
├── parser/
│   ├── markdown_parser.go (new: Markdown 解析)
│   └── scanner.go       (new: 文件系统扫描)
└── handler/
    └── documents.go     (new: 文档 API)
```

### References
- [Source: architecture.md#Data Architecture] 数据模型和缓存策略
- [Source: architecture.md#API & Communication Patterns] API 端点设计
- [Source: architecture.md#Implementation Patterns] 命名规范
- [Source: epics.md#Story 1.2] 原始故事定义和验收标准

## Dev Agent Record

### Agent Model Used

### Debug Log References

### Completion Notes List

### File List

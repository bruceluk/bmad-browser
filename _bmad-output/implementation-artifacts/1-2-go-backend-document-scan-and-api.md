# Story 1.2: Go 后端文档扫描与 API

Status: review

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

- [x] Task 1: 定义数据模型 (AC: #5, #7)
  - [x] 创建 `server/model/document.go`，定义 Document struct（Path、Title、Content、Frontmatter、Phase）
  - [x] JSON tag 使用 camelCase（如 `json:"path"`、`json:"frontmatter"`）
  - [x] 创建 DocumentSummary struct（Path、Title、Phase，不含 Content），用于列表 API

- [x] Task 2: 实现 Markdown 解析器 (AC: #2)
  - [x] 创建 `server/parser/markdown_parser.go`
  - [x] 实现 ParseMarkdown 函数：读取 .md 文件，分离 YAML frontmatter 和 Markdown 内容
  - [x] frontmatter 解析为 `map[string]interface{}`
  - [x] 从 frontmatter 或 Markdown 内容提取标题（优先 frontmatter 的 title 字段，否则取第一个 `# ` 标题）
  - [x] 引入 YAML 解析库（`gopkg.in/yaml.v3` v3.0.1）

- [x] Task 3: 实现文件系统扫描器 (AC: #1, #3, #4)
  - [x] 创建 `server/parser/scanner.go`
  - [x] 实现 ScanDocuments 函数：递归扫描指定目录下所有 `.md` 文件
  - [x] 对每个文件调用 ParseMarkdown 解析
  - [x] 根据目录结构推断 BMAD 阶段（brainstorming→analysis, planning→planning, implementation→implementation）
  - [x] 返回 `[]Document` 缓存在内存中
  - [x] 支持通过命令行参数 `-dir` 指定扫描目录路径（默认 `../_bmad-output`）

- [x] Task 4: 实现文档 API handler (AC: #5, #6, #7)
  - [x] 创建 `server/handler/documents.go`
  - [x] 实现 `GET /api/documents` handler：返回所有文档的 DocumentSummary 列表
  - [x] 实现 `GET /api/documents/:path` handler：用路径前缀匹配 + strings.TrimPrefix 提取路径参数
  - [x] 文档不存在时返回 404 `{"error": "document not found"}`
  - [x] Content-Type 设置为 `application/json`

- [x] Task 5: 集成到 main.go (AC: #1, #4)
  - [x] 修改 `server/main.go`：启动时调用 ScanDocuments
  - [x] 将扫描结果传递给 DocumentHandler
  - [x] 添加命令行参数（`-port`、`-dir`）
  - [x] 保留 `/api/health` 端点
  - [x] 日志输出扫描到的文件数量（"Scanned 10 documents"）

- [x] Task 6: 验证 (AC: #1-#7)
  - [x] Go 编译成功
  - [x] 启动后日志显示 "Scanned 10 documents from ../../_bmad-output"
  - [x] `curl /api/documents` 返回 10 个文档的列表（JSON 数组）
  - [x] `curl /api/documents/planning-artifacts/prd.md` 返回完整文档（title: "产品需求文档 - BMAD Viewer", 5720 chars content）
  - [x] `curl /api/documents/nonexistent.md` 返回 404 `{"error":"document not found"}`

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

Claude Opus 4.6 (1M context)

### Debug Log References

- gopkg.in/yaml.v3 v3.0.1 安装成功
- Go 标准库路由使用 HandleFunc 路径前缀匹配 + TrimPrefix 提取路径参数

### Completion Notes List

- ✅ Document + DocumentSummary 数据模型定义完成，含 ToSummary 转换方法
- ✅ Markdown 解析器实现：frontmatter 分离、YAML 解析、标题提取（frontmatter 优先，降级到 # 标题）
- ✅ 文件扫描器实现：递归扫描 .md 文件、阶段推断、内存缓存
- ✅ 文档 API handler：列表（无 content）和详情（含 content）两个端点
- ✅ main.go 集成：启动时扫描、命令行参数（-port、-dir）、日志输出
- ✅ 端到端验证通过：10 个文档扫描成功，列表 API、详情 API、404 均正常
- ✅ 删除不再需要的 .gitkeep 占位文件

### File List

- bmad-viewer/server/model/document.go (new)
- bmad-viewer/server/parser/markdown_parser.go (new)
- bmad-viewer/server/parser/scanner.go (new)
- bmad-viewer/server/handler/documents.go (new)
- bmad-viewer/server/main.go (modified)
- bmad-viewer/server/go.mod (modified)
- bmad-viewer/server/go.sum (new)
- bmad-viewer/server/handler/.gitkeep (deleted)
- bmad-viewer/server/parser/.gitkeep (deleted)
- bmad-viewer/server/model/.gitkeep (deleted)

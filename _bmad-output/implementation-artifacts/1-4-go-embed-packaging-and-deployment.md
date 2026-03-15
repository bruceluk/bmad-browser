# Story 1.4: Go embed 打包与一行部署

Status: done

## Story

As a 部署者（Lu），
I want 通过 `go run` 一行命令启动完整应用，
so that 我可以快速部署给团队使用。

## Acceptance Criteria

1. `make build` 执行前端构建（Vite）后通过 `//go:embed` 将 `web/dist/` 嵌入 Go 二进制，生成单一可执行文件
2. 执行 `./bmad-viewer`（或 `go run ./server`）可启动完整应用（FR13）
3. 同一内网用户可以通过浏览器访问前端页面和 API（FR14）
4. 前端页面正常加载，API 正常响应
5. 支持通过命令行参数或环境变量指定端口号和文档目录路径
6. 启动时间 < 5 秒（NFR5）

## Tasks / Subtasks

- [x] Task 1: 创建 embed.go 嵌入前端资源 (AC: #1)
  - [x] 创建 `server/embed.go`：`//go:embed all:dist`（构建时复制 web/dist → server/dist）
  - [x] 注意：go:embed 不支持 `..` 路径，改为构建时复制方案

- [x] Task 2: 创建静态资源 handler (AC: #3, #4)
  - [x] 创建 `server/handler/static.go`
  - [x] fs.Sub 去掉 "dist" 前缀，http.FileServer 提供静态文件
  - [x] SPA 回退：fs.Stat 检查文件是否存在，不存在则重写为 /（index.html）
  - [x] API 路由保护：/api/ 开头的请求直接 404（不走静态文件）

- [x] Task 3: 集成到 main.go (AC: #2, #5)
  - [x] 注册 `http.HandleFunc("/", handler.NewStaticHandler(webDistFS))`
  - [x] API 路由注册在前，优先匹配
  - [x] 命令行参数 -port、-dir 保持不变

- [x] Task 4: 更新 Makefile (AC: #1)
  - [x] build: 前端构建 → 复制 web/dist 到 server/dist → Go 编译
  - [x] 添加 `make run` 目标
  - [x] clean 包含 server/dist 清理

- [x] Task 5: 端到端验证 (AC: #1-#6)
  - [x] `make build` 成功，生成 8.7MB 单一可执行文件（Mach-O arm64）
  - [x] 启动后扫描到 12 个文档
  - [x] /api/health 返回 {"status":"ok"}
  - [x] /api/documents 返回 12 个文档列表
  - [x] / 返回前端 HTML
  - [x] /doc/planning-artifacts/prd.md SPA 回退正常（返回 index.html）
  - [x] 静态资源（CSS）正常返回 200

## Dev Notes

### 前序故事已建立的基础
**Story 1.1:** 项目结构、Makefile（dev/build/clean）、Vite + Go 开发环境
**Story 1.2:** Go 后端 API（/api/health, /api/documents, /api/documents/:path）、命令行参数（-port, -dir）
**Story 1.3:** Vue 前端完整功能（文档列表、文档详情、Markdown 渲染）、Vite 构建输出到 web/dist/

### Go embed 关键要点

**embed.go 位置和路径：**
```go
// server/embed.go
package main

import "embed"

//go:embed all:../web/dist
var webDistFS embed.FS
```

注意：`//go:embed` 路径相对于 Go 源文件所在目录（server/），所以用 `../web/dist`。
`all:` 前缀确保包含以 `.` 和 `_` 开头的文件。

**SPA 回退逻辑：**
Go embed 的文件系统嵌入后路径为 `web/dist/...`，需要用 `io/fs.Sub` 去掉前缀：
```go
import "io/fs"

sub, _ := fs.Sub(webDistFS, "web/dist")
fileServer := http.FileServer(http.FS(sub))
```

SPA 回退：对于非 API、非静态文件的请求，返回 index.html：
```go
func spaHandler(fsHandler http.Handler, sub fs.FS) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // 尝试提供静态文件
        path := r.URL.Path
        _, err := fs.Stat(sub, strings.TrimPrefix(path, "/"))
        if err != nil {
            // 文件不存在，返回 index.html（SPA 回退）
            r.URL.Path = "/"
        }
        fsHandler.ServeHTTP(w, r)
    }
}
```

### 路由优先级
main.go 中路由注册顺序很重要：
1. `/api/health` — 精确匹配
2. `/api/documents/` — 前缀匹配（文档详情）
3. `/api/documents` — 精确匹配（文档列表）
4. `/` — 默认路由（静态资源 + SPA 回退）

Go 标准库 `http.HandleFunc` 按最长前缀匹配，API 路由会优先于默认 `/` 路由。

### Makefile 构建流程
```makefile
build:
	cd web && npm run build
	cd server && go build -o ../bmad-viewer .

run: build
	./bmad-viewer
```

### 架构约束
- Go 文件名用 `snake_case`（`embed.go`、`static.go`）
- 单一可执行文件，零外部依赖
- HTTP only（内网），无 HTTPS

### 文件结构
本故事创建/修改的文件：
```
server/
├── embed.go           (new: //go:embed 前端资源)
├── main.go            (modify: 注册静态资源 handler)
├── handler/
│   └── static.go      (new: 静态文件服务 + SPA 回退)
Makefile               (modify: 更新 build 目标)
```

### References
- [Source: architecture.md#Infrastructure & Deployment] 单一可执行文件部署
- [Source: architecture.md#Starter Template Evaluation] Go embed 构建流程
- [Source: epics.md#Story 1.4] 原始故事定义

## Dev Agent Record

### Agent Model Used

Claude Opus 4.6 (1M context)

### Debug Log References

- go:embed 不支持 `..` 相对路径，改为 Makefile 构建时 `cp -r web/dist server/dist`
- fs.Sub 路径从 "web/dist" 改为 "dist"（匹配 embed 路径）

### Completion Notes List

- ✅ embed.go：`//go:embed all:dist` 嵌入前端构建产物
- ✅ static.go：静态文件服务 + SPA 回退（fs.Stat + 重写 / ）+ API 路由保护
- ✅ main.go：注册静态 handler 为默认路由，API 路由优先
- ✅ Makefile：build 目标包含前端构建→复制→Go 编译完整流程
- ✅ 8.7MB 单一可执行文件，端到端验证全部通过
- ✅ SPA 回退正常，Vue Router history 模式路径可直接访问

### File List

- bmad-viewer/server/embed.go (new)
- bmad-viewer/server/handler/static.go (new)
- bmad-viewer/server/main.go (modified)
- bmad-viewer/Makefile (modified)
- bmad-viewer/server/dist/ (new, build artifact, gitignored)

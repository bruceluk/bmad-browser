# Story 1.1: 项目初始化与开发环境搭建

Status: review

## Story

As a 开发者（Lu），
I want 初始化 monorepo 项目结构（web/ + server/），配置开发环境，
so that 前后端可以独立开发并通过 Vite 代理联调。

## Acceptance Criteria

1. 克隆仓库后执行 `npm install`（web/）和 `go mod download`（server/）可成功安装依赖
2. `cd web && npm run dev` 可启动 Vite 开发服务器
3. `cd server && go run .` 可启动 Go 后端（返回简单健康检查响应）
4. Vite 代理 `/api/*` 请求到 Go 后端（vite.config.ts 配置）
5. 项目结构符合架构文档定义
6. Tailwind CSS 4 + `@tailwindcss/typography` 已配置并生效
7. Makefile 包含 `make dev` 和 `make build` 目标

## Tasks / Subtasks

- [x] Task 1: 创建项目根目录和基础文件 (AC: #5)
  - [x] 创建 `bmad-viewer/` 根目录
  - [x] 创建 `.gitignore`（Go + Node 合并规则：node_modules/、dist/、*.exe）
  - [x] 创建 `Makefile`（dev 和 build 目标，详见下方 Dev Notes）

- [x] Task 2: 初始化 Vue 3 前端 (AC: #1, #2, #5)
  - [x] 手动创建 Vue 3 + TypeScript + Router 项目结构（create-vue CLI 在非 TTY 环境不可用，改为手动配置）
  - [x] 清理不需要的示例文件（手动创建故无需清理）
  - [x] 创建空的目录结构：`src/components/`、`src/views/`、`src/api/`、`src/types/`

- [x] Task 3: 配置 Tailwind CSS (AC: #6)
  - [x] 安装 `tailwindcss` 和 `@tailwindcss/typography`
  - [x] 在 `src/style.css` 中使用 Tailwind CSS 4 的 `@import "tailwindcss"` 和 `@plugin` 指令
  - [x] 添加深色主题 CSS 变量
  - [x] 验证 Tailwind 类在组件中生效（前端构建成功，HomeView 使用 Tailwind 类）

- [x] Task 4: 配置 Vite API 代理 (AC: #4)
  - [x] 在 `vite.config.ts` 中添加 proxy 配置：`/api` → `http://localhost:8080`

- [x] Task 5: 初始化 Go 后端 (AC: #1, #3, #5)
  - [x] 创建 `server/` 目录
  - [x] 运行 `go mod init bmad-viewer/server`
  - [x] 创建 `server/main.go`：最小 HTTP server，监听 8080 端口
  - [x] 添加健康检查端点 `GET /api/health` 返回 `{"status": "ok"}`
  - [x] 创建空的目录结构：`handler/`、`parser/`、`model/`（使用 .gitkeep 占位）

- [x] Task 6: 创建 Makefile (AC: #7)
  - [x] `make dev`：并行启动前后端开发服务器
  - [x] `make build`：前端构建 + Go 编译

- [x] Task 7: 验证端到端联调 (AC: #2, #3, #4)
  - [x] Go 后端编译成功，`/api/health` 返回 `{"status":"ok"}`
  - [x] 前端 TypeScript 类型检查通过，Vite 构建成功
  - [x] npm install 安装依赖成功（56 packages）

## Dev Notes

### 技术栈版本要求
- Go 1.22+
- Node.js 18+
- Vue 3（最新稳定版）
- TypeScript 严格模式
- Vite 6.4.1
- Tailwind CSS 4
- `@tailwindcss/typography` 插件

### 实现说明
- create-vue CLI 在非 TTY 环境下无法交互，改为手动创建项目结构
- Tailwind CSS 4 使用新的 `@import "tailwindcss"` 语法替代旧版 `@tailwind` 指令
- Tailwind CSS 4 使用 `@plugin` 指令加载插件，不再需要 `tailwind.config.js`
- 空目录使用 `.gitkeep` 文件占位以确保 git 跟踪

### References

- [Source: architecture.md#Project Structure] 完整目录结构定义
- [Source: architecture.md#Starter Template Evaluation] 初始化命令
- [Source: architecture.md#Implementation Patterns] 命名规范和编码约束
- [Source: ux-design-specification.md#Visual Design Foundation] 色彩系统和字体系统

## Dev Agent Record

### Agent Model Used

Claude Opus 4.6 (1M context)

### Debug Log References

- create-vue CLI 在非 TTY 环境报错 ERR_PARSE_ARGS_UNKNOWN_OPTION，改为手动创建项目结构
- Tailwind CSS 4 不使用 tailwind.config.js，改用 CSS 内的 @import 和 @plugin 指令

### Completion Notes List

- ✅ 项目根目录创建完成，包含 .gitignore 和 Makefile
- ✅ Vue 3 前端手动初始化：TypeScript 严格模式、Vue Router、Vite 6
- ✅ Tailwind CSS 4 配置完成，使用新版 @import/@plugin 语法
- ✅ 深色主题 CSS 变量定义完成（12 个变量）
- ✅ 中文字体栈和等宽代码字体配置完成
- ✅ Vite API 代理配置完成（/api → localhost:8080）
- ✅ Go 后端初始化：标准库 HTTP server + /api/health 端点
- ✅ 前端 TypeScript 类型检查通过，Vite 生产构建成功
- ✅ Go 后端编译成功，健康检查 API 返回正常
- ✅ npm install 安装 56 个包，0 漏洞

### File List

- bmad-viewer/.gitignore (new)
- bmad-viewer/Makefile (new)
- bmad-viewer/web/package.json (new)
- bmad-viewer/web/tsconfig.json (new)
- bmad-viewer/web/vite.config.ts (new)
- bmad-viewer/web/index.html (new)
- bmad-viewer/web/env.d.ts (new)
- bmad-viewer/web/src/main.ts (new)
- bmad-viewer/web/src/App.vue (new)
- bmad-viewer/web/src/style.css (new)
- bmad-viewer/web/src/vite-env.d.ts (new)
- bmad-viewer/web/src/router/index.ts (new)
- bmad-viewer/web/src/views/HomeView.vue (new)
- bmad-viewer/server/go.mod (new)
- bmad-viewer/server/main.go (new)
- bmad-viewer/server/handler/.gitkeep (new)
- bmad-viewer/server/parser/.gitkeep (new)
- bmad-viewer/server/model/.gitkeep (new)

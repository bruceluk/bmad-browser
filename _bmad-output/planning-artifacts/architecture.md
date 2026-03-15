---
stepsCompleted: ['step-01-init', 'step-02-context', 'step-03-starter', 'step-04-decisions', 'step-05-patterns']
inputDocuments: ['_bmad-output/planning-artifacts/prd.md', '_bmad-output/planning-artifacts/ux-design-specification.md', '_bmad-output/planning-artifacts/product-brief-bamd-2026-03-13.md']
workflowType: 'architecture'
project_name: 'BMAD Viewer'
user_name: 'Lu'
date: '2026-03-15'
---

# Architecture Decision Document

_This document builds collaboratively through step-by-step discovery. Sections are appended as we work through each architectural decision together._

## Project Context Analysis

### Requirements Overview

**Functional Requirements:**
- 18 个功能需求，分为 5 类：
  - 文档浏览（FR1-FR5）：扫描目录、按阶段分类、渲染 Markdown
  - 命令与产出映射（FR6-FR9）：解析 CSV、展示命令/代理/产出关系
  - 导航与信息架构（FR10-FR12）：首页概览、导航菜单、位置感知
  - 部署与运行（FR13-FR15）：`go run` 启动、内网访问、自动扫描
  - 数据解析（FR16-FR18）：Markdown frontmatter、CSV 解析、阶段推断
- 架构含义：所有需求围绕"读取文件系统 → 解析数据 → 结构化展示"，无写入、无状态变更

**Non-Functional Requirements:**
- NFR1: 首页加载 < 2秒
- NFR2: 文档切换 < 500ms
- NFR3: Markdown 渲染 < 1秒
- NFR4: 支持 < 50 并发用户
- NFR5: 启动时间 < 5秒
- 架构含义：性能要求宽松，内存缓存即可满足，无需复杂优化

**UX 架构需求：**
- 流程图优先布局：上方固定流程图区 + 下方文档区
- 角色入口导航：数据需按角色（开发者/PM/测试）组织，而非按文件系统目录
- 6 个自建 Vue 组件：RoleCard、RoleTab、FlowNode、FlowArrow、DocMeta、DocRenderer
- 深色主题 Tailwind CSS，仅桌面端，仅 Chrome

**Scale & Complexity:**
- Primary domain: 全栈 Web（Go 后端 + Vue 3 前端）
- Complexity level: 低
- Estimated architectural components: ~8（Go HTTP server、文件扫描器、CSV 解析器、Markdown 解析器、内存缓存、Vue SPA、流程图组件、文档渲染组件）

### Technical Constraints & Dependencies

1. **零外部依赖**：无数据库、无消息队列、无第三方服务
2. **单一可执行文件**：Go embed 嵌入前端资源
3. **文件系统即数据源**：`_bmad-output/` 目录 + `bmad-help.csv`
4. **仅 Chrome**：无需浏览器兼容
5. **仅 HTTP**：内网部署，无 HTTPS
6. **只读**：无用户输入、无认证、无会话
7. **一人开发**：Lu 全栈开发，架构必须简单直接

### Cross-Cutting Concerns Identified

1. **角色-流程映射**：CSV 中的工作流数据需转换为按角色分组的流程图结构，前后端都需理解这个映射
2. **Markdown 渲染**：后端解析 frontmatter 提取元数据，前端渲染 Markdown 内容为 HTML
3. **数据一致性**：CSV 中的 outputs 字段需与文件系统中的实际文档对应

## Starter Template Evaluation

### Primary Technology Domain

全栈 Web 应用（Go 后端 + Vue 3 前端），技术栈在 PRD 和 UX 设计中已明确确定。

### Starter Options Considered

**前端（Vue 3）：**
- `create-vue`（Vue 官方脚手架）——自动配置 Vue Router + TypeScript + Vite，社区标准

**后端（Go）：**
- 无需脚手架，`go mod init` + 手动组织即可——Go 项目结构简单，标准库足够

### Selected Starter

**前端：** `create-vue` 初始化 Vue 3 + TypeScript + Vue Router + Vite
**后端：** `go mod init` 手动组织

### Initialization Commands

```bash
# 项目根目录
mkdir bmad-viewer && cd bmad-viewer

# 前端初始化
npm create vue@latest web -- --typescript --router --no-pinia --no-vitest --no-e2e --no-eslint

# 后端初始化
mkdir -p server
cd server && go mod init bmad-viewer/server

# Tailwind CSS
cd ../web && npm install tailwindcss @tailwindcss/typography
```

### Project Structure

```
bmad-viewer/
├── web/                    # Vue 3 前端
│   ├── src/
│   │   ├── components/     # Vue 组件（RoleCard, FlowNode, DocRenderer 等）
│   │   ├── views/          # 页面视图（HomeView, FlowView）
│   │   ├── router/         # Vue Router 配置
│   │   ├── types/          # TypeScript 类型定义
│   │   ├── api/            # 后端 API 调用封装
│   │   ├── App.vue
│   │   └── main.ts
│   ├── index.html
│   ├── tailwind.config.js
│   ├── vite.config.ts
│   └── package.json
├── server/                 # Go 后端
│   ├── main.go             # 入口，HTTP server 启动
│   ├── handler/            # HTTP 路由处理
│   ├── parser/             # Markdown 和 CSV 解析
│   ├── model/              # 数据结构定义
│   ├── embed.go            # 前端资源嵌入
│   └── go.mod
└── Makefile                # 构建脚本（前端构建 + Go embed 打包）
```

### Architectural Decisions Provided by Starter

**Language & Runtime:**
- 前端：TypeScript（严格模式），Vue 3 Composition API + `<script setup>`
- 后端：Go 1.22+，标准库为主

**Styling Solution:**
- Tailwind CSS 4 + `@tailwindcss/typography` 插件
- 深色主题通过 CSS 变量定义

**Build Tooling:**
- 前端：Vite（开发服务器 + 生产构建）
- 后端：`go build`，通过 `//go:embed` 嵌入前端构建产物
- Makefile 统一构建流程：`make build` = 前端构建 + Go 编译 = 单一可执行文件

**Code Organization:**
- 前端按职责分层：components / views / api / types
- 后端按职责分层：handler / parser / model
- 前后端通过 REST API 通信

**Development Experience:**
- 开发时前后端分别启动：Vite dev server（前端）+ `go run`（后端）
- Vite 代理 API 请求到 Go 后端，避免跨域
- 生产时合并为单一可执行文件

## Core Architectural Decisions

### Decision Priority Analysis

**Critical Decisions (Block Implementation):**
1. 数据架构：文件系统 + 启动时内存缓存
2. API 设计：3 个 REST 端点
3. 前端状态管理：Vue 响应式 API

**Deferred Decisions (Post-MVP):**
- 文件变更监听（热更新缓存）——MVP 重启即可刷新
- API 分页——文档数量少，不需要

### Data Architecture

**数据源：** 文件系统（`_bmad-output/` 目录 + `bmad-help.csv`）

**缓存策略：** 启动时一次性扫描
- Go 服务启动时扫描 `_bmad-output/` 目录，解析所有 Markdown 文件的 frontmatter
- 同时解析 `bmad-help.csv`，构建工作流数据结构
- 所有数据缓存到内存中的 Go struct
- 运行期间不再读取文件系统，所有 API 请求从内存返回
- 文件变更需重启服务生效（MVP 足够）

**数据模型：**
```go
// 工作流步骤
type WorkflowStep struct {
    Name        string   // 步骤名称
    Code        string   // 步骤代码
    Command     string   // BMAD 命令
    AgentName   string   // 代理名称
    AgentIcon   string   // 代理图标
    Phase       string   // 所属阶段
    Required    bool     // 是否必需
    Description string   // 描述
    Outputs     []string // 产出文件
    Duration    string   // 预期时间
}

// 角色流程
type RoleFlow struct {
    Role      string         // 角色名称（开发者/PM/测试）
    Steps     []WorkflowStep // 该角色的流程步骤
    Upstream  []WorkflowStep // 上游步骤（其他角色）
    Downstream []WorkflowStep // 下游步骤（其他角色）
}

// 文档
type Document struct {
    Path        string            // 文件路径
    Title       string            // 文档标题
    Content     string            // Markdown 原始内容
    Frontmatter map[string]interface{} // YAML frontmatter
    Phase       string            // 所属 BMAD 阶段
}
```

### Authentication & Security

**不适用。** 只读、内网、无用户认证、无会话管理。

### API & Communication Patterns

**API 风格：** REST，JSON 响应

**端点设计：**

| 端点 | 方法 | 说明 | 响应 |
|------|------|------|------|
| `/api/roles` | GET | 返回角色列表和对应的流程步骤 | `RoleFlow[]` |
| `/api/documents/:path` | GET | 返回指定文档的内容和元数据 | `Document` |
| `/api/workflows` | GET | 返回 CSV 解析后的完整工作流数据 | `WorkflowStep[]` |

**错误处理：**
- 404：文档不存在
- 500：服务器内部错误
- 无需 401/403（无认证）

**CORS：** 开发环境需要（Vite dev server 跨域请求 Go 后端），生产环境不需要（同源）

### Frontend Architecture

**状态管理：** Vue 响应式 API（ref/reactive），无 Pinia
- `currentRole: ref<string>` — 当前选中角色
- `currentNode: ref<string>` — 当前选中流程节点
- `visitedNodes: ref<Set<string>>` — 已浏览节点集合
- `roles: ref<RoleFlow[]>` — 从 API 获取的角色流程数据

**路由：** Vue Router
- `/` — 首页（角色选择）
- `/flow/:role` — 流程图视图（带角色参数）
- `/flow/:role/:step` — 流程图视图 + 选中节点

**Markdown 渲染：** 前端使用 markdown-it 库将 Markdown 转为 HTML，配合 Tailwind Typography 样式

### Infrastructure & Deployment

**部署方式：** 单一可执行文件
- `make build` → 前端构建（Vite）→ Go embed 嵌入 → `go build` → 单一二进制文件
- 部署：拷贝二进制文件到服务器 → `./bmad-viewer` 启动
- 或直接在源码目录 `go run ./server` 启动（需先构建前端）

**环境配置：**
- 命令行参数或环境变量指定：端口号（默认 8080）、`_bmad-output/` 目录路径、`bmad-help.csv` 路径
- 无需配置文件

**监控与日志：**
- Go 标准库 `log` 包，输出到 stdout
- 记录：启动信息、扫描文件数量、API 请求日志
- 无需专业监控工具

### Decision Impact Analysis

**Implementation Sequence:**
1. Go 数据模型定义（model/）
2. CSV 解析器 + Markdown 解析器（parser/）
3. 启动时扫描和内存缓存
4. REST API 端点（handler/）
5. Vue 前端组件和路由
6. Go embed 集成和构建流程

**Cross-Component Dependencies:**
- 前端 TypeScript 类型需与 Go 数据模型对应（手动保持一致）
- CSV 解析逻辑决定了角色-流程映射的数据结构，前后端共享理解
- Markdown 渲染在前端完成，后端只提供原始内容

## Implementation Patterns & Consistency Rules

### Naming Patterns

**Go 代码（server/）：**
- 变量/函数：`camelCase`（未导出）、`PascalCase`（导出）
- 文件名：`snake_case.go`（如 `csv_parser.go`、`role_handler.go`）
- 包名：小写单词（如 `handler`、`parser`、`model`）

**Vue/TypeScript 代码（web/）：**
- 组件文件：`PascalCase.vue`（如 `FlowNode.vue`、`RoleCard.vue`）
- TypeScript 文件：`camelCase.ts`（如 `apiClient.ts`）
- 类型/接口：`PascalCase`（如 `RoleFlow`、`WorkflowStep`）
- 变量/函数：`camelCase`（如 `currentRole`、`fetchRoles`）
- CSS 类名：Tailwind utility classes，无自定义类名

**API：**
- 端点路径：`/api/kebab-case`（如 `/api/roles`、`/api/documents`）
- JSON 字段：`camelCase`（Go struct 通过 `json:"camelCase"` tag 转换）

### Format Patterns

**API 响应格式：**
- 成功时直接返回数据，不包装：
```json
// 数组响应
[{"role": "developer", "steps": [...]}]

// 对象响应
{"path": "...", "title": "...", "content": "..."}
```
- 错误时返回统一结构：
```json
{"error": "document not found"}
```
- HTTP 状态码：200（成功）、404（资源不存在）、500（服务器错误）

**日期格式：** ISO 8601 字符串（如 `"2026-03-15"`）

### Structure Patterns

**组件组织：** 按类型（已在项目结构中确定）
```
web/src/components/   # 所有 Vue 组件平铺
web/src/views/        # 页面级视图
web/src/api/          # API 调用封装
web/src/types/        # TypeScript 类型定义
```

**后端组织：** 按职责
```
server/handler/       # HTTP 路由处理函数
server/parser/        # 数据解析（CSV、Markdown）
server/model/         # 数据结构定义
```

**测试位置：** 与源文件同目录
- Go：`handler/role_handler_test.go`
- Vue：`components/__tests__/FlowNode.test.ts`（如需要）

### Process Patterns

**错误处理：**
- Go：函数返回 `(result, error)`，handler 层统一转为 HTTP 错误响应
- Vue：API 调用使用 `try/catch`，错误信息输出到 `console.error`
- 不做全局错误弹窗（只读应用，错误场景极少）

**加载状态：**
- 鼠标光标 `cursor: wait`，无其他加载指示
- 状态变量命名：`isLoading`（布尔值）

### Enforcement Guidelines

**所有 AI 代理必须遵守：**
1. Go 文件用 `snake_case`，Vue 组件用 `PascalCase`
2. API 响应直接返回数据，不加包装层
3. JSON 字段用 `camelCase`
4. 不引入未在架构文档中列出的第三方依赖
5. 不添加数据库、认证、中间件等未规划的功能
6. 前端不写 `<style>` 块，所有样式用 Tailwind utility classes

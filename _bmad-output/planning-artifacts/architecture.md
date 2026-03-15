---
stepsCompleted: ['step-01-init', 'step-02-context', 'step-03-starter']
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

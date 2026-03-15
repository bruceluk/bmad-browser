---
stepsCompleted: ['step-01-validate-prerequisites', 'step-02-design-epics', 'step-03-create-stories', 'step-04-final-validation']
workflow_completed: true
inputDocuments: ['_bmad-output/planning-artifacts/prd.md', '_bmad-output/planning-artifacts/architecture.md', '_bmad-output/planning-artifacts/ux-design-specification.md']
---

# BMAD Viewer - Epic Breakdown

## Overview

This document provides the complete epic and story breakdown for BMAD Viewer, decomposing the requirements from the PRD, UX Design, and Architecture requirements into implementable stories.

## Requirements Inventory

### Functional Requirements

- FR1: 浏览者可以查看 `_bmad-output/` 目录下所有 Markdown 文档的渲染内容
- FR2: 浏览者可以按 BMAD 阶段分类浏览文档
- FR3: 浏览者可以通过阶段列表导航在不同阶段间切换
- FR4: 浏览者可以查看单个文档的完整渲染内容
- FR5: 浏览者可以在文档间切换，无需重新加载页面
- FR6: 浏览者可以查看所有 BMAD 工作流命令的列表
- FR7: 浏览者可以查看每个命令对应的 AI 代理角色名称和职责
- FR8: 浏览者可以查看每个命令产出的文件列表
- FR9: 浏览者可以从命令映射页面导航到对应的产出文档
- FR10: 浏览者可以在首页看到 BMAD 项目的整体结构概览
- FR11: 浏览者可以通过导航菜单访问所有主要功能区域
- FR12: 浏览者可以看到当前所在位置
- FR13: 部署者可以通过 `go run` 单一命令启动应用
- FR14: 应用启动后，同一内网的用户可以通过浏览器访问
- FR15: 应用可以自动扫描指定目录下的文档，无需手动配置文件列表
- FR16: 系统可以解析 Markdown 文件并提取 YAML frontmatter 元数据
- FR17: 系统可以解析 `bmad-help.csv` 并提取工作流、命令、代理角色等字段
- FR18: 系统可以根据目录结构自动推断文档所属的 BMAD 阶段

### NonFunctional Requirements

- NFR1: 首页加载时间 < 2秒（内网环境）
- NFR2: 文档切换响应时间 < 500ms
- NFR3: Markdown 文档渲染时间 < 1秒
- NFR4: 系统支持 < 50 并发用户
- NFR5: 应用启动时间 < 5秒

### Additional Requirements

- 项目初始化：`create-vue`（前端）+ `go mod init`（后端），monorepo 结构 web/ + server/
- Go embed 嵌入前端构建产物，单一可执行文件部署
- 3 个 REST API 端点：`/api/roles`、`/api/documents/:path`、`/api/workflows`
- 启动时一次性扫描文件系统，内存缓存所有数据
- Makefile 统一构建流程
- Vite 开发代理配置（开发环境 API 代理到 Go 后端）

### UX Design Requirements

- UX-DR1: 实现深色主题色彩系统（背景色 `#1a1a2e`、表面色 `#16213e`、三色角色标识：天蓝/草绿/淡紫）
- UX-DR2: 实现首页三个角色卡片组件（RoleCard），包含角色图标、名称、说明、步骤数、总预期时间，hover 上浮+发光效果
- UX-DR3: 实现流程图视图顶部角色 Tab 组件（RoleTab），点击切换角色流程图，激活态使用角色色
- UX-DR4: 实现流程图节点组件（FlowNode），显示步骤名称/命令/代理图标/预期时间，支持默认/hover/激活/已浏览/上下游五种状态
- UX-DR5: 实现流程图箭头组件（FlowArrow），连接节点，角色色/灰色两种状态
- UX-DR6: 实现文档元信息组件（DocMeta），水平展示代理徽章、命令代码、时间徽章、难度徽章
- UX-DR7: 实现 Markdown 文档渲染组件（DocRenderer），基于 `@tailwindcss/typography` prose 样式，深色主题适配，支持标题/代码块/表格/列表
- UX-DR8: 实现流程图优先布局（上方固定流程图区 + 下方弹性文档区），流程图始终可见
- UX-DR9: 实现上下游延伸交互——点击虚线节点自动切换角色 Tab 并定位到对应节点
- UX-DR10: 实现中文字体栈（PingFang SC / Microsoft YaHei）+ 等宽代码字体（JetBrains Mono），完整字号层级（28/22/18/16/14/12px）

### FR Coverage Map

| FR | Epic | 说明 |
|----|------|------|
| FR1 | Epic 1 | 查看所有 Markdown 文档 |
| FR2 | Epic 2 | 按阶段分类浏览 |
| FR3 | Epic 2 | 阶段列表导航 |
| FR4 | Epic 1 | 查看单个文档完整内容 |
| FR5 | Epic 1 | 文档间切换无刷新 |
| FR6 | Epic 3 | 查看工作流命令列表 |
| FR7 | Epic 3 | 查看命令对应的代理角色 |
| FR8 | Epic 3 | 查看命令产出文件 |
| FR9 | Epic 3 | 从命令映射导航到产出文档 |
| FR10 | Epic 2 | 首页结构概览 |
| FR11 | Epic 2 | 导航菜单 |
| FR12 | Epic 2 | 位置感知 |
| FR13 | Epic 1 | `go run` 启动 |
| FR14 | Epic 1 | 内网访问 |
| FR15 | Epic 1 | 自动扫描文档 |
| FR16 | Epic 1 | Markdown frontmatter 解析 |
| FR17 | Epic 2 | CSV 解析 |
| FR18 | Epic 2 | 阶段推断 |

NFR1-NFR5 通过架构决策（内存缓存 + SPA + Go 标准库）整体满足，不单独分配到某个 Epic。

## Epic List

### Epic 1: 项目基础与文档浏览
用户可以通过 `go run` 启动应用，打开浏览器看到 BMAD 项目的文档列表，并阅读任意 Markdown 文档的渲染内容。
**FRs covered:** FR1, FR4, FR5, FR13, FR14, FR15, FR16
**Additional:** 项目初始化（create-vue + go mod init）、Go embed、Makefile、Vite 代理配置
**UX-DRs:** UX-DR1（深色主题基础）、UX-DR7（DocRenderer）、UX-DR10（字体系统）

### Epic 2: 角色导航与流程图
用户打开首页看到三个角色入口，选择角色后看到该角色的线性工作流程图，每个节点显示步骤名称、命令、代理角色和预期时间。
**FRs covered:** FR2, FR3, FR10, FR11, FR12, FR17, FR18
**UX-DRs:** UX-DR2（RoleCard）、UX-DR3（RoleTab）、UX-DR4（FlowNode）、UX-DR5（FlowArrow）、UX-DR8（流程图优先布局）

### Epic 3: 流程图与文档联动
用户点击流程图节点，下方展示对应的产出文档和元信息。用户可以点击上下游虚线节点延伸到其他角色的流程。
**FRs covered:** FR6, FR7, FR8, FR9
**UX-DRs:** UX-DR6（DocMeta）、UX-DR9（上下游延伸交互）

## Epic 1: 项目基础与文档浏览

用户可以通过 `go run` 启动应用，打开浏览器看到 BMAD 项目的文档列表，并阅读任意 Markdown 文档的渲染内容。

### Story 1.1: 项目初始化与开发环境搭建

As a 开发者（Lu），
I want 初始化 monorepo 项目结构（web/ + server/），配置开发环境，
So that 前后端可以独立开发并通过 Vite 代理联调。

**Acceptance Criteria:**

**Given** 一台全新的开发机器
**When** 克隆仓库并执行 `npm install`（web/）和 `go mod download`（server/）
**Then** 前后端依赖安装成功
**And** `cd web && npm run dev` 启动 Vite 开发服务器
**And** `cd server && go run .` 启动 Go 后端
**And** Vite 代理 `/api/*` 请求到 Go 后端（vite.config.ts 配置）
**And** 项目结构符合架构文档定义（web/src/components、views、api、types + server/handler、parser、model）
**And** Tailwind CSS 4 + `@tailwindcss/typography` 已配置
**And** Makefile 包含 `make dev`（并行启动前后端）和 `make build`（构建单一可执行文件）

### Story 1.2: Go 后端文档扫描与 API

As a 浏览者，
I want 应用启动时自动扫描文档目录并提供 API 访问，
So that 我可以通过浏览器获取文档数据。

**Acceptance Criteria:**

**Given** `_bmad-output/` 目录下有 Markdown 文件
**When** Go 服务启动
**Then** 自动扫描 `_bmad-output/` 目录下所有 `.md` 文件
**And** 解析每个文件的 YAML frontmatter 元数据（FR16）
**And** 根据目录结构推断文档所属的 BMAD 阶段（FR18）
**And** 所有数据缓存到内存

**Given** 服务已启动
**When** 请求 `GET /api/documents/:path`
**Then** 返回指定文档的 JSON 数据（path、title、content、frontmatter、phase）
**And** 文档不存在时返回 404 `{"error": "document not found"}`

**Given** 服务已启动
**When** 请求 `GET /api/documents`（无 path 参数）
**Then** 返回所有文档的列表（path、title、phase，不含 content）

### Story 1.3: 前端文档浏览与渲染

As a 浏览者，
I want 在浏览器中查看文档列表并阅读渲染后的 Markdown 内容，
So that 我可以浏览 BMAD 项目的产出物。

**Acceptance Criteria:**

**Given** 用户打开浏览器访问应用
**When** 页面加载完成
**Then** 显示文档列表（FR1）
**And** 页面使用深色主题（UX-DR1：背景色 `#1a1a2e`、文字色 `#e0e0e0`）
**And** 使用中文字体栈和等宽代码字体（UX-DR10）

**Given** 文档列表已展示
**When** 用户点击某个文档
**Then** 展示该文档的完整渲染内容（FR4）
**And** Markdown 渲染为 HTML，支持标题、列表、代码块、表格（UX-DR7）
**And** 使用 Tailwind Typography prose 样式，深色主题适配
**And** 切换文档无需整页刷新（FR5，SPA 路由切换）

### Story 1.4: Go embed 打包与一行部署

As a 部署者（Lu），
I want 通过 `go run` 一行命令启动完整应用，
So that 我可以快速部署给团队使用。

**Acceptance Criteria:**

**Given** 已执行 `make build`
**When** 前端构建产物（web/dist/）通过 `//go:embed` 嵌入 Go 二进制
**Then** 生成单一可执行文件

**Given** 单一可执行文件已生成
**When** 执行 `./bmad-viewer`（或 `go run ./server`）
**Then** HTTP 服务在默认端口 8080 启动（FR13）
**And** 同一内网用户可以通过浏览器访问（FR14）
**And** 前端页面正常加载，API 正常响应
**And** 支持通过命令行参数或环境变量指定端口号和文档目录路径
**And** 启动时间 < 5 秒（NFR5）

## Epic 2: 角色导航与流程图

用户打开首页看到三个角色入口，选择角色后看到该角色的线性工作流程图，每个节点显示步骤名称、命令、代理角色和预期时间。

### Story 2.1: CSV 解析与角色流程 API

As a 浏览者，
I want 应用能解析 BMAD 工作流数据并按角色分组返回，
So that 前端可以展示每个角色的工作流程图。

**Acceptance Criteria:**

**Given** `bmad-help.csv` 文件存在
**When** Go 服务启动
**Then** 解析 CSV 文件，提取工作流名称、命令、代理角色、阶段、描述等字段（FR17）
**And** 将工作流数据按角色（开发者/产品经理/测试人员）分组，构建 RoleFlow 结构
**And** 每个角色的流程包含核心步骤、上游步骤和下游步骤
**And** 数据缓存到内存

**Given** 服务已启动
**When** 请求 `GET /api/roles`
**Then** 返回角色列表，每个角色包含：角色名称、流程步骤（含名称、命令、代理、预期时间）、上游步骤、下游步骤

**Given** 服务已启动
**When** 请求 `GET /api/workflows`
**Then** 返回 CSV 解析后的完整工作流数据列表

### Story 2.2: 首页角色选择

As a 浏览者，
I want 打开首页看到三个角色入口和简短介绍，
So that 我能一眼找到与自己相关的内容。

**Acceptance Criteria:**

**Given** 用户打开应用首页
**When** 页面加载完成
**Then** 显示标题"BMAD Viewer"和介绍文字（FR10）
**And** 显示三个角色卡片：开发者（天蓝 `#4fc3f7`）、产品经理（草绿 `#81c784`）、测试人员（淡紫 `#b39ddb`）（UX-DR2）
**And** 每个卡片包含角色图标、角色名称、一句话说明、步骤数量、总预期时间

**Given** 用户将鼠标悬停在角色卡片上
**When** hover 状态触发
**Then** 卡片上浮 4px 并显示角色色边框发光效果（150ms ease 过渡）

**Given** 用户点击某个角色卡片
**When** 点击事件触发
**Then** 路由跳转到 `/flow/:role`，进入该角色的流程图视图

### Story 2.3: 流程图视图与节点展示

As a 浏览者，
I want 看到所选角色的线性工作流程图，每个节点显示步骤信息，
So that 我能了解"我的工作有哪几步、每步用什么命令"。

**Acceptance Criteria:**

**Given** 用户从首页选择了一个角色
**When** 进入流程图视图 `/flow/:role`
**Then** 上方显示角色 Tab 切换栏，当前角色高亮（UX-DR3）
**And** 中间显示该角色的线性流程图（UX-DR8：上方固定流程图区）
**And** 流程图节点水平排列，箭头连接（UX-DR5）
**And** 每个节点显示：步骤名称、BMAD 命令（等宽字体）、代理角色图标、预期时间（UX-DR4）
**And** 上游步骤以虚线边框 + 40% 透明度显示，标注"上游"
**And** 下游步骤以虚线边框 + 40% 透明度显示，标注"下游"
**And** 默认不选中任何节点，下方文档区显示"点击上方流程节点查看详情"

**Given** 流程图视图已展示
**When** 用户点击角色 Tab 切换到另一个角色
**Then** 流程图节点即时更新为新角色的流程（FR3, FR11）
**And** 当前 Tab 使用角色色背景高亮

**Given** 用户在流程图视图中
**When** 用户将鼠标悬停在流程节点上
**Then** 节点边框变为角色色，上浮 2px（150ms ease 过渡）

**Given** 用户在流程图视图中
**When** 用户查看浏览器地址栏或面包屑
**Then** 能看到当前位置信息（FR12）
**And** 左上角 BMAD Viewer 标题可点击返回首页

## Epic 3: 流程图与文档联动

用户点击流程图节点，下方展示对应的产出文档和元信息。用户可以点击上下游虚线节点延伸到其他角色的流程。

### Story 3.1: 节点点击与文档展示

As a 浏览者，
I want 点击流程图节点后在下方看到该步骤的产出文档和元信息，
So that 我能了解每个环节具体做了什么、产出了什么。

**Acceptance Criteria:**

**Given** 用户在流程图视图中
**When** 用户点击一个流程节点
**Then** 该节点高亮为角色色边框 + 8% 透明度背景（激活态）
**And** 下方文档区展示该步骤的元信息：代理角色徽章、BMAD 命令（等宽字体代码样式）、预期时间徽章、难度徽章（UX-DR6）
**And** 元信息下方展示对应的产出文档渲染内容（Markdown → HTML）
**And** 文档区自动滚动到顶部
**And** 路由更新为 `/flow/:role/:step`

**Given** 用户已点击过某个节点
**When** 用户点击另一个节点
**Then** 之前的节点变为已浏览状态（角色色 40% 透明度）
**And** 新节点变为激活态
**And** 文档区内容替换为新节点的文档

**Given** 用户点击一个流程节点
**When** 该节点没有对应的产出文档
**Then** 文档区显示"该步骤暂无产出文档"

**Given** 用户点击流程节点查看文档
**When** 数据请求中
**Then** 鼠标光标变为 `cursor: wait`

### Story 3.2: 命令与代理角色信息展示

As a 浏览者，
I want 在流程图和文档详情中看到每个步骤的命令、代理角色和产出文件信息，
So that 我知道每一步用什么命令、由哪个 AI 代理负责、产出什么文件。

**Acceptance Criteria:**

**Given** 用户在流程图视图中查看节点
**When** 流程图节点已渲染
**Then** 每个节点显示 BMAD 命令名称（FR6）
**And** 每个节点显示对应的 AI 代理角色图标和名称（FR7）

**Given** 用户点击某个流程节点
**When** 文档元信息区（DocMeta）展示
**Then** 显示该命令产出的文件列表（FR8）
**And** 产出文件可点击，跳转到对应的产出文档查看（FR9）

### Story 3.3: 上下游延伸导航

As a 浏览者，
I want 点击流程图中的上下游虚线节点后自动切换到对应角色的流程，
So that 我能自然地了解上下游同事的工作流程。

**Acceptance Criteria:**

**Given** 用户在开发者流程图视图中
**When** 用户点击上游虚线节点（如 PM 的"创建 PRD"）
**Then** 角色 Tab 自动切换到产品经理（UX-DR9）
**And** 流程图更新为产品经理的流程
**And** 对应节点自动激活并高亮
**And** 下方文档区展示该节点的产出文档

**Given** 用户在产品经理流程图视图中
**When** 用户点击下游虚线节点（如开发者的"故事开发"）
**Then** 角色 Tab 自动切换到开发者
**And** 流程图更新为开发者的流程
**And** 对应节点自动激活

**Given** 用户通过上下游延伸导航切换了角色
**When** 流程图更新完成
**Then** 路由更新为 `/flow/:newRole/:step`
**And** 之前角色的已浏览状态保留（刷新页面后重置）

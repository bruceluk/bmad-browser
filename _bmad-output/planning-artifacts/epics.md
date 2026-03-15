---
stepsCompleted: ['step-01-validate-prerequisites', 'step-02-design-epics']
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

{{requirements_coverage_map}}

## Epic List

{{epics_list}}

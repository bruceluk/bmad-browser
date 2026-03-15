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

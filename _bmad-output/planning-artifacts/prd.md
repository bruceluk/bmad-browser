---
stepsCompleted: ['step-01-init', 'step-02-discovery', 'step-02b-vision', 'step-02c-executive-summary', 'step-03-success', 'step-04-journeys', 'step-05-domain', 'step-06-innovation', 'step-07-project-type', 'step-08-scoping', 'step-09-functional', 'step-10-nonfunctional', 'step-11-polish', 'step-12-complete']
workflow_completed: true
inputDocuments: ['_bmad-output/planning-artifacts/product-brief-bamd-2026-03-13.md', '_bmad-output/brainstorming/brainstorming-session-2026-03-13-1200.md']
workflowType: 'prd'
documentCounts:
  briefs: 1
  research: 0
  brainstorming: 1
  projectDocs: 0
  projectContext: 0
classification:
  projectType: 'web_app'
  domain: 'general'
  complexity: 'low'
  projectContext: 'greenfield'
---

# 产品需求文档 - BMAD Viewer

**作者:** Lu
**日期:** 2026-03-13

## Executive Summary

BMAD Viewer 是一个面向团队的只读 Web 应用（Vue 3 + Golang），将 BMAD 方法的完整项目产出物以结构化、可交互的方式展示。目标用户是需要学习人+AI 协作工作模式的团队成员——包括开发者、产品经理、设计师和测试人员。核心问题：团队缺乏系统化的方式学习 AI 协作工作流，市面课程与实际工作脱节，各人各用 AI 缺乏统一方法，没有覆盖产品全生命周期的真实参考案例。

### What Makes This Special

**元学习闭环**——项目本身就是用 BMAD 方法构建的，构建过程即学习内容，"做项目的过程就是教材"。与通用 AI 课程不同，这里展示的是一个同事真正走过的从需求到上线的完整过程，而非理论幻灯片。零门槛体验：无需注册、无需数据库、`go run` 一行启动，团队成员即刻浏览。方法论可视化：不只展示"产出了什么"，更展示"为什么这样做"和"怎么做的"。

## Project Classification

- **项目类型：** Web 应用（SPA，浏览器访问）
- **领域：** 通用软件（内部团队学习工具）
- **复杂度：** 低（标准 Web 技术栈，无特殊合规要求）
- **项目背景：** 全新项目（greenfield）

## Success Criteria

### User Success

- 团队成员浏览 BMAD Viewer 后，能清晰描述人+AI 协作在需求、设计、开发、测试各环节的具体运作方式
- 成员产生"原来 AI 可以这样参与工作"的认知突破，从"只会用 AI 聊天"进阶到理解系统化协作流程
- 成员看到完整的命令→代理角色→产出物映射，知道"下一步该做什么、用什么命令"

### Business Success

- **1个月内：** 全体团队成员（100%）完成 BMAD Viewer 浏览
- **3个月内：** ≥50% 成员尝试使用过 BMAD 命令
- **3个月内：** ≥1 个真实项目使用 BMAD 方法完整交付
- **核心目标：** BMAD 方法成为团队日常工作流程的一部分

### Technical Success

- `go run` 一行启动，零外部依赖，零数据库
- 正确解析并渲染 `_bmad-output/` 下所有 Markdown 文档
- 正确解析 `bmad-help.csv` 并展示命令与产出映射
- 性能达标（详见 Non-Functional Requirements）

### Measurable Outcomes

| 指标 | 衡量方式 | 目标 |
|------|----------|------|
| 团队认知覆盖 | 浏览过 BMAD Viewer 的成员比例 | 100% |
| 实践转化率 | 尝试使用过 BMAD 命令的成员比例 | ≥50% |
| 项目落地 | 使用 BMAD 方法完成的真实项目数 | ≥1个（3个月内） |
| 部署可用性 | `go run` 一行启动即可用 | 100% |

## Product Scope

### MVP Strategy

**MVP 方式：** 问题解决型——用最少的功能让团队成员能浏览 BMAD 项目产出物，验证"看真实案例比听理论更有效"这一核心假设。

**资源：** Lu 一人开发，前后端全栈。

### MVP Feature Set (Phase 1)

**支持的核心旅程：**
- 旅程一（部署）：`go run` 启动，内网可访问
- 旅程二~四（浏览）：按阶段导航、查看文档、查看命令映射

**必须具备的能力：**
1. Go 后端扫描 `_bmad-output/` 目录，解析 Markdown 文件并通过 API 返回
2. Go 后端解析 `bmad-help.csv`，返回命令、代理角色、产出文件的映射数据
3. Vue 前端渲染 BMAD 阶段列表导航
4. Vue 前端渲染 Markdown 文档内容
5. Vue 前端展示命令与产出映射页面
6. Go embed 嵌入前端资源，单一可执行文件部署

**明确排除（MVP 不做）：**
- 用户登录/认证
- 评论和反馈功能
- 文档编辑功能
- 数据库
- 交互式流程图
- Git 演进时间线

### Phase 2: Growth

- 交互式 BMAD 工作流程图导航（mermaid 或 d3）
- 自适应双模式布局（全景概览↔文档详情）

### Phase 3: Expansion

- 逆向追溯导航
- Git 演进时间线与版本对比
- Markdown 渲染后的 diff 视图
- 评论反馈、多项目支持、学习路径引导

### Risk Mitigation

**技术风险：** 低。Vue 3 + Go 均为成熟技术栈，文件系统读取和 Markdown 解析都有现成库支持。最大技术挑战是 Go embed 嵌入前端资源的构建流程，可通过早期验证解决。

**市场风险：** 低。用户是自己的团队，需求明确，无需市场验证。核心风险是团队成员"看了但没行动"——通过 Lu 的会议引导来弥补。

**资源风险：** 一人开发，时间是主要约束。MVP 功能已精简到最小，优先保证核心浏览体验可用，避免过度设计。

## User Journeys

### 旅程一：Lu 部署并推广 BMAD Viewer

**人物：** Lu，技术负责人，已走通 BMAD 全流程，希望带动团队转型。

**开场：** Lu 完成了 BMAD Viewer 的开发，准备让团队看到成果。他登录团队内网服务器，将项目代码拉取到服务器上。

**行动：** 在服务器上执行 `go run`，服务启动成功，确认内网可访问。Lu 打开浏览器验证页面正常加载——文档列表完整、命令映射正确、Markdown 渲染清晰。

**关键时刻：** Lu 将链接发送到团队群组，附一句话："这是我用 BMAD 方法做的一个完整项目，大家可以先自己看看，下周我们开会讨论。"

**结果：** 团队成员收到链接，随时可以打开浏览。Lu 完成了从"自己学会"到"分享给团队"的第一步。

---

### 旅程二：开发者小王自由探索

**人物：** 小王，后端开发，用过 Cursor 辅助写代码，但从没在需求和架构阶段用过 AI。

**开场：** 小王在群里看到 Lu 发的链接，午休时随手点开。首页展示 BMAD 的阶段列表和文档结构。

**探索：** 小王先点了自己熟悉的区域——故事实现相关的文档，看到开发故事的详细规格。然后好奇地往上游点——产品简报、PRD，发现这些文档都是通过特定的 BMAD 命令和 AI 代理角色协作产出的。

**认知突破：** 在命令与产出映射页面，小王看到 `bmad-bmm-create-prd` 命令对应 📋 John（产品经理代理），产出了完整的 PRD。他意识到：AI 不只是帮写代码，还能参与需求分析、架构设计等上游环节。

**结果：** 小王带着几个问题找 Lu："这个 PRD 是 AI 自动生成的吗？我们以后写需求也可以这样？" 学习兴趣被激发。

---

### 旅程三：产品经理小李深度阅读

**人物：** 小李，产品经理，平时只用 AI 聊天问问题，没有系统化的 AI 协作经验。

**开场：** 小李打开链接，直奔自己关心的——PRD 和产品简报。

**深度阅读：** 小李仔细阅读了产品简报的完整内容，注意到它的结构：执行摘要→核心愿景→目标用户→成功指标→MVP 范围。每个部分都条理清晰、信息密度高。然后查看命令映射，发现产品简报是通过 `bmad-bmm-create-product-brief` 命令，由 AI 代理引导一步步完成的。

**认知突破：** 小李发现这份产品简报的质量不亚于自己写过的最好的简报，而且结构更规范、覆盖更全面。她意识到 AI 代理不是替代产品经理，而是像一个经验丰富的搭档一样引导思考。

**结果：** 小李在团队会议上主动提问："我下个项目能不能试试用这个方法写 PRD？" 转化发生了。

---

### 旅程四：设计师/测试人员快速浏览

**人物：** 小赵，UI 设计师，对 AI 好奇但不知道跟自己的工作有什么关系。

**开场：** 小赵在会议前花了 10 分钟快速浏览 BMAD Viewer。

**浏览：** 点击各个阶段快速扫一遍，在 UX 设计相关文档处停留，看到设计规范是如何通过 AI 辅助生成的。虽然这个项目的设计部分不复杂，但小赵看到了 AI 在整个工作流中的角色分布。

**初步认识：** 小赵对 AI 在设计领域的具体应用还有疑问，但至少理解了 BMAD 方法的全貌——"原来一个项目从头到尾可以这样做"。

**结果：** 会议上小赵问 Lu："设计环节具体怎么用 AI？有没有更详细的例子？" 好奇心被激活，等待下一步引导。

### Journey Requirements Summary

| 能力 | 对应旅程 | 优先级 |
|------|----------|--------|
| 文档结构化展示与 Markdown 渲染 | 所有旅程 | MVP |
| BMAD 阶段列表导航 | 所有旅程 | MVP |
| 命令与产出映射展示 | 旅程二、三 | MVP |
| 代理角色与职责展示 | 旅程二、三 | MVP |
| `go run` 一行部署 | 旅程一 | MVP |
| 内网服务器可访问 | 旅程一 | MVP |
| 快速页面加载（< 2秒） | 旅程四（快速浏览） | MVP |

## Web App Specific Requirements

### Project-Type Overview

单页应用（SPA），Vue 3 前端 + Golang 后端，仅支持 Chrome 浏览器，部署于团队内网服务器。只读浏览，无实时交互需求，无 SEO 需求，无无障碍性要求。

### Technical Architecture Considerations

**前端架构：**
- Vue 3 SPA，使用 Vue Router 实现客户端路由
- Markdown 渲染组件（客户端渲染）
- 通过 REST API 从后端获取文档数据

**后端架构：**
- Golang HTTP 服务器，提供 REST API
- 直接读取文件系统（`_bmad-output/` 目录 + `bmad-help.csv`）
- 内存缓存已解析的文档数据
- 同时托管 Vue 前端静态资源

**浏览器支持：**
- 仅 Chrome（最新稳定版）
- 可自由使用现代 CSS 和 JS 特性，无需兼容性处理

### Implementation Considerations

- 前后端可合并为单一可执行文件（Go embed 嵌入前端资源），实现 `go run` 一行启动
- 无需 HTTPS（内网环境），HTTP 即可
- 无需用户认证、会话管理
- 无需数据库驱动或 ORM

## Functional Requirements

### 文档浏览

- FR1: 浏览者可以查看 `_bmad-output/` 目录下所有 Markdown 文档的渲染内容
- FR2: 浏览者可以按 BMAD 阶段（分析、规划、架构、实施）分类浏览文档
- FR3: 浏览者可以通过阶段列表导航在不同阶段间切换
- FR4: 浏览者可以查看单个文档的完整渲染内容
- FR5: 浏览者可以在文档间切换，无需重新加载页面（响应时间见 NFR2）

### 命令与产出映射

- FR6: 浏览者可以查看所有 BMAD 工作流命令的列表
- FR7: 浏览者可以查看每个命令对应的 AI 代理角色名称和职责
- FR8: 浏览者可以查看每个命令产出的文件列表
- FR9: 浏览者可以从命令映射页面导航到对应的产出文档

### 导航与信息架构

- FR10: 浏览者可以在首页看到 BMAD 项目的整体结构概览
- FR11: 浏览者可以通过导航菜单访问所有主要功能区域（文档浏览、命令映射）
- FR12: 浏览者可以看到当前所在位置（面包屑或高亮导航项）

### 部署与运行

- FR13: 部署者可以通过 `go run` 单一命令启动应用
- FR14: 应用启动后，同一内网的用户可以通过浏览器访问
- FR15: 应用可以自动扫描指定目录下的文档，无需手动配置文件列表

### 数据解析

- FR16: 系统可以解析 Markdown 文件并提取 YAML frontmatter 元数据
- FR17: 系统可以解析 `bmad-help.csv` 并提取工作流、命令、代理角色等字段
- FR18: 系统可以根据目录结构自动推断文档所属的 BMAD 阶段

## Non-Functional Requirements

### Performance

- NFR1: 首页加载时间 < 2秒（内网环境）
- NFR2: 文档切换响应时间 < 500ms（非首次访问时）
- NFR3: Markdown 文档渲染时间 < 1秒（单个文档）
- NFR4: 系统支持 < 50 并发用户，响应时间不超过单用户时的 2 倍
- NFR5: 应用启动时间 < 5秒（从启动到可接受请求）

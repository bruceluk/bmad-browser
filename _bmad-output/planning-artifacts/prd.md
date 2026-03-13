---
stepsCompleted: ['step-01-init', 'step-02-discovery', 'step-02b-vision', 'step-02c-executive-summary', 'step-03-success', 'step-04-journeys']
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
- 页面加载时间 < 2秒（本地/局域网环境）
- 正确解析并渲染 `_bmad-output/` 下所有 Markdown 文档
- 正确解析 `bmad-help.csv` 并展示命令与产出映射

### Measurable Outcomes

| 指标 | 衡量方式 | 目标 |
|------|----------|------|
| 团队认知覆盖 | 浏览过 BMAD Viewer 的成员比例 | 100% |
| 实践转化率 | 尝试使用过 BMAD 命令的成员比例 | ≥50% |
| 项目落地 | 使用 BMAD 方法完成的真实项目数 | ≥1个（3个月内） |
| 部署可用性 | `go run` 一行启动即可用 | 100% |

## Product Scope

### MVP - Minimum Viable Product

1. **文档浏览器：** Go 后端扫描 `_bmad-output/` 目录，解析 Markdown 文件，Vue 前端以结构化方式渲染展示
2. **零数据库架构：** Go 直接读取文件系统，内存缓存，`go run` 一行启动
3. **命令与产出映射：** 解析 `bmad-help.csv`，展示每个工作流的命令、代理角色和对应产出文件
4. **基本导航：** BMAD 阶段列表导航，点击阶段查看该阶段的产出文档

### Growth Features (Post-MVP)

- 交互式 BMAD 工作流程图导航，点击阶段节点直达产出文档
- 自适应双模式布局（全景概览↔文档详情自然切换）

### Vision (Future)

- 逆向追溯导航——从当前进度往回追溯每个阶段的产出和决策
- Git 演进时间线——通过 Git 历史展示文档版本变化
- Markdown 渲染后的 diff 对比视图
- 评论反馈、多项目支持、学习路径引导

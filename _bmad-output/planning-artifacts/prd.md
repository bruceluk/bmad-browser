---
stepsCompleted: ['step-01-init', 'step-02-discovery', 'step-02b-vision', 'step-02c-executive-summary', 'step-03-success']
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

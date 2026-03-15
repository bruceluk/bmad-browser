---
stepsCompleted: ['step-01-init', 'step-02-discovery', 'step-03-core-experience']
inputDocuments: ['_bmad-output/planning-artifacts/product-brief-bamd-2026-03-13.md', '_bmad-output/planning-artifacts/prd.md']
---

# UX Design Specification - BMAD Viewer

**Author:** Lu
**Date:** 2026-03-13

---

<!-- UX design content will be appended sequentially through collaborative workflow steps -->

## Executive Summary

### Project Vision

BMAD Viewer 是一个面向软件开发团队的只读 Web 应用，将 BMAD 方法的完整项目产出物以结构化方式展示，让团队成员通过浏览真实项目的全过程来学习人+AI 协作的工作模式。项目本身即用 BMAD 方法构建，形成"元学习闭环"。

### Target Users

所有用户均为软件开发团队成员，具备基本技术素养：
- **团队推动者（Lu）**：部署并推广，引导团队学习
- **开发者**：熟悉 AI 编程助手，需了解 AI 在上游环节的作用
- **产品经理**：需了解 AI 如何融入产品工作流
- **设计师/测试人员**：需了解 AI 在各自领域的应用场景

**用户特征：**
- 技术背景：全员软件开发相关，无需降低技术门槛
- 使用场景：桌面端深度阅读（30 分钟以上），不考虑移动端
- 界面偏好：无特定限制，以好用为准，开发者友好的 UI 模式可接受

### Key Design Challenges

1. **深度阅读舒适度**：连续阅读 30 分钟以上的长文档，需解决阅读疲劳、定位和导航问题（目录锚点、阅读进度感知）
2. **信息架构认知负担**：BMAD 工作流涉及多阶段、多代理角色、多命令，新用户首次接触易信息过载，需渐进式呈现
3. **文档间关联发现**：用户浏览文档时需自然地发现上下文关系（命令来源、下一步推荐），跨文档跳转需流畅

### Design Opportunities

1. **工作流即导航**：利用 BMAD 阶段的自然顺序作为信息架构，用户跟着工作流走就能理解全貌
2. **代理角色作为记忆锚点**：用 AI 代理角色的视觉标识（图标+名称）帮助用户快速定位上下文
3. **开发者友好的阅读体验**：面向开发者可大胆使用代码风格 UI 模式（等宽字体、语法高亮、侧边栏树状结构）

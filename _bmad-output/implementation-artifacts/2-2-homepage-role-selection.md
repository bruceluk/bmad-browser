# Story 2.2: 首页角色选择

Status: ready-for-dev

## Story

As a 浏览者，
I want 打开首页看到三个角色入口和简短介绍，
so that 我能一眼找到与自己相关的内容。

## Acceptance Criteria

1. 首页显示标题"BMAD Viewer"和介绍文字（FR10）
2. 显示三个角色卡片：开发者（天蓝 `#4fc3f7`）、产品经理（草绿 `#81c784`）、测试人员（淡紫 `#b39ddb`）（UX-DR2）
3. 每个卡片包含角色图标、角色名称、一句话说明、步骤数量、总预期时间
4. 卡片 hover 时上浮 4px 并显示角色色边框发光效果（150ms ease）
5. 点击角色卡片，路由跳转到 `/flow/:role`（流程图视图）

## Tasks / Subtasks

- [ ] Task 1: 创建 RoleCard 组件 (AC: #2, #3, #4)
  - [ ] 创建 `web/src/components/RoleCard.vue`
  - [ ] Props: role (string), roleColor (string), label (string), description (string), stepCount (number), totalDuration (string)
  - [ ] 卡片布局：角色图标（emoji）、角色名称、一句话说明、步骤数和时间
  - [ ] 默认态：`var(--border)` 边框，`var(--surface)` 背景
  - [ ] Hover 态：上浮 4px（translateY(-4px)）+ 角色色边框 + box-shadow 发光
  - [ ] 过渡：150ms ease
  - [ ] 点击 emit click 事件

- [ ] Task 2: 改造 HomeView 为角色选择页 (AC: #1, #2, #5)
  - [ ] 重写 `web/src/views/HomeView.vue`
  - [ ] 页面加载时调用 `fetchRoles()` 获取角色数据
  - [ ] 显示标题"BMAD Viewer"和介绍文字
  - [ ] 渲染三个 RoleCard 组件，水平排列居中
  - [ ] 点击卡片路由跳转到 `/flow/:role`
  - [ ] 计算每个角色的步骤数量和总预期时间（从 RoleFlow.steps 计算）

- [ ] Task 3: 添加流程图路由 (AC: #5)
  - [ ] 修改 `web/src/router/index.ts`
  - [ ] 添加 `/flow/:role` 路由（FlowView 组件，下一个故事实现，此处先用占位组件）
  - [ ] 创建占位 `web/src/views/FlowView.vue`（显示角色名称，等待 Story 2.3 填充）

- [ ] Task 4: 验证 (AC: #1-#5)
  - [ ] TypeScript 类型检查通过
  - [ ] Vite 构建成功
  - [ ] 首页显示三个角色卡片，数据从 API 获取
  - [ ] Hover 效果正常
  - [ ] 点击跳转到 /flow/:role

## Dev Notes

### 已建立的代码库
**Story 2.1 完成：**
- `/api/roles` API 已实现，返回 3 个 RoleFlow（含 role、roleColor、label、steps、upstream、downstream）
- `fetchRoles()` 前端 API 客户端已实现
- WorkflowStep + RoleFlow TypeScript 类型已定义

**当前 HomeView：**
- 目前是文档列表页（Story 1.3 实现），需要完全重写为角色选择页
- 文档列表功能将在后续故事中集成到流程图视图的文档区

### 角色卡片设计参考（来自 UX 设计规范）

| 角色 | 图标 | 颜色 | 说明 |
|------|------|------|------|
| 开发者 | 💻 | #4fc3f7 | 从需求到代码：了解故事怎么来的、怎么开发、怎么审查 |
| 产品经理 | 📋 | #81c784 | 从想法到需求：了解产品简报、PRD、史诗故事怎么产出 |
| 测试人员 | 🧪 | #b39ddb | 从开发到验收：了解自动化测试、QA 流程怎么跑 |

### 步骤数量和时间计算
从 RoleFlow.steps 数组计算：
- stepCount = steps.length
- totalDuration：解析每个 step 的 duration 字段（如 "~30min"），提取数字求和，格式化为 "约X小时Y分钟"

### 架构约束
- Vue 组件文件名用 `PascalCase`
- 不写 `<style>` 块，所有样式用 Tailwind utility classes
- CSS 变量用于颜色（已在 style.css 中定义）
- 角色色需要通过内联 style 绑定（Tailwind 不支持动态 CSS 变量值）

### 布局参考
- 三个卡片水平排列，间距 24px（gap-6）
- 居中显示，最大宽度适配三卡片
- 卡片宽度约 300px
- 圆角 16px（rounded-2xl）
- 内边距 28px

### References
- [Source: ux-design-specification.md#Component Strategy] RoleCard 组件规格
- [Source: ux-design-specification.md#User Journey Flows] 首页角色选择流程
- [Source: ux-design-specification.md#Interaction Feedback Patterns] hover 效果
- [Source: epics.md#Story 2.2] 原始故事定义

## Dev Agent Record

### Agent Model Used

### Debug Log References

### Completion Notes List

### File List

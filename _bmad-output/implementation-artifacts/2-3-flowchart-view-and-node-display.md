# Story 2.3: 流程图视图与节点展示

Status: ready-for-dev

## Story

As a 浏览者，
I want 看到所选角色的线性工作流程图，每个节点显示步骤信息，
so that 我能了解"我的工作有哪几步、每步用什么命令"。

## Acceptance Criteria

1. 进入流程图视图 `/flow/:role`，上方显示角色 Tab 切换栏，当前角色高亮（UX-DR3）
2. 中间显示该角色的线性流程图（UX-DR8：上方固定流程图区）
3. 流程图节点水平排列，箭头连接（UX-DR5）
4. 每个节点显示：步骤名称、BMAD 命令（等宽字体）、代理角色图标、预期时间（UX-DR4）
5. 上游步骤以虚线边框 + 40% 透明度显示
6. 下游步骤以虚线边框 + 40% 透明度显示
7. 默认不选中任何节点，下方文档区显示"点击上方流程节点查看详情"
8. 点击角色 Tab 切换到另一个角色时，流程图节点即时更新（FR3, FR11）
9. 节点 hover 时边框变为角色色，上浮 2px（150ms ease）
10. 左上角 BMAD Viewer 标题可点击返回首页

## Tasks / Subtasks

- [ ] Task 1: 创建 RoleTab 组件 (AC: #1, #8)
  - [ ] 创建 `web/src/components/RoleTab.vue`
  - [ ] Props: roles (RoleFlow[]), currentRole (string)
  - [ ] 每个 Tab 显示角色图标 + 角色名称
  - [ ] 当前角色 Tab 使用角色色背景高亮
  - [ ] 其他 Tab 灰色边框，hover 时边框变角色色
  - [ ] 点击 emit select 事件（传递 role 值）

- [ ] Task 2: 创建 FlowNode 组件 (AC: #4, #5, #6, #9)
  - [ ] 创建 `web/src/components/FlowNode.vue`
  - [ ] Props: step (WorkflowStep), roleColor (string), isUpstream (boolean), isDownstream (boolean)
  - [ ] 节点内容：步骤名称、命令（等宽字体）、代理图标、预期时间
  - [ ] 默认态：var(--border) 边框
  - [ ] Hover 态：角色色边框 + 上浮 2px（150ms ease）
  - [ ] 上下游态：虚线边框 + 整体 40% 透明度
  - [ ] 点击 emit click 事件

- [ ] Task 3: 创建 FlowArrow 组件 (AC: #3)
  - [ ] 创建 `web/src/components/FlowArrow.vue`
  - [ ] Props: color (string，角色色或灰色)
  - [ ] 显示箭头符号 →

- [ ] Task 4: 重写 FlowView 为完整流程图视图 (AC: #1-#10)
  - [ ] 重写 `web/src/views/FlowView.vue`
  - [ ] 页面加载时调用 fetchRoles() 获取所有角色数据
  - [ ] 顶部：BMAD Viewer 标题（可点击返回首页）+ RoleTab 组件
  - [ ] 中间（固定区域）：线性流程图（上游节点 → 核心节点 → 下游节点，箭头连接）
  - [ ] 下方：空状态提示"点击上方流程节点查看详情"
  - [ ] 切换角色 Tab 时更新流程图
  - [ ] 通过路由参数 :role 初始化当前角色
  - [ ] 支持 /flow/:role/:step 路由（为 Epic 3 准备，此处仅解析但不激活节点）

- [ ] Task 5: 更新路由配置 (AC: #10)
  - [ ] 修改 router：添加 `/flow/:role/:step?` 可选 step 参数

- [ ] Task 6: 验证 (AC: #1-#10)
  - [ ] TypeScript 类型检查通过
  - [ ] Vite 构建成功
  - [ ] 从首页点击角色卡片进入流程图视图
  - [ ] 流程图正确显示核心节点 + 上下游虚线节点
  - [ ] Tab 切换正常
  - [ ] Hover 效果正常
  - [ ] 返回首页链接正常

## Dev Notes

### 已建立的代码库
**Story 2.2 完成：**
- HomeView 重写为角色选择页，调用 fetchRoles()，三个 RoleCard
- FlowView 目前是占位组件，需要完全重写
- RoleCard 组件模式可参考（props、hover 效果实现方式）
- /flow/:role 路由已注册

**可用 API 和类型：**
- `fetchRoles(): Promise<RoleFlow[]>` — 返回 3 角色，每个含 steps/upstream/downstream
- RoleFlow: { role, roleColor, label, steps, upstream, downstream }
- WorkflowStep: { name, code, command, agentName, agentIcon, phase, required, description, outputs, duration, sequence }

### 布局结构（来自 UX 设计规范：流程图优先）
```
┌─────────────────────────────────────────┐
│ BMAD Viewer (可点击返回)                   │
│ [开发者] [产品经理] [测试人员]  ← RoleTab   │
├─────────────────────────────────────────┤
│                                         │
│  [上游]→[步骤1]→[步骤2]→...→[下游]       │  ← 固定流程图区
│                                         │
├─────────────────────────────────────────┤
│                                         │
│     点击上方流程节点查看详情               │  ← 文档区（Epic 3 实现联动）
│                                         │
└─────────────────────────────────────────┘
```

### FlowNode 节点设计（来自 UX 设计规范）
- 宽度：min-width 140px
- 圆角：12px（rounded-xl）
- 内边距：16px 20px
- 节点内容纵向排列：图标 → 名称 → 命令（等宽 10px）→ 时间
- 上下游节点：border-dashed + opacity-40

### RoleTab 设计
- 水平排列，间距 8px
- 每个 Tab：padding 8px 20px，圆角 20px（rounded-full 药丸形）
- 激活态：角色色背景 15% 透明度 + 角色色边框 + 角色色文字
- 非激活态：var(--border) 边框

### 角色图标映射（复用 HomeView 的映射）
```typescript
const roleIcons: Record<string, string> = {
  developer: '💻',
  pm: '📋',
  qa: '🧪',
}
```

### 架构约束
- Vue 组件文件名 PascalCase
- 不写 `<style>` 块
- 动态颜色用内联 style（Tailwind 不支持动态值）
- 流程图区域使用 CSS Flexbox 水平排列

### References
- [Source: ux-design-specification.md#Design Direction Decision] 流程图优先布局
- [Source: ux-design-specification.md#Component Strategy] RoleTab、FlowNode、FlowArrow 规格
- [Source: ux-design-specification.md#Interaction Feedback Patterns] hover 效果
- [Source: ux-design-specification.md#State Patterns] 空状态提示
- [Source: epics.md#Story 2.3] 原始故事定义

## Dev Agent Record

### Agent Model Used

### Debug Log References

### Completion Notes List

### File List

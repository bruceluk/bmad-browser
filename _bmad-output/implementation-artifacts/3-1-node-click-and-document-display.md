# Story 3.1: 节点点击与文档展示

Status: ready-for-dev

## Story

As a 浏览者，
I want 点击流程图节点后在下方看到该步骤的产出文档和元信息，
so that 我能了解每个环节具体做了什么、产出了什么。

## Acceptance Criteria

1. 点击流程节点后，该节点高亮为角色色边框 + 8% 透明度背景（激活态）
2. 下方文档区展示该步骤的元信息：代理角色徽章、BMAD 命令（等宽字体代码样式）、预期时间徽章、难度徽章（UX-DR6）
3. 元信息下方展示对应的产出文档渲染内容（Markdown → HTML）
4. 文档区自动滚动到顶部
5. 路由更新为 `/flow/:role/:step`
6. 点击另一个节点时，之前节点变为已浏览状态（角色色 40% 透明度），新节点激活
7. 文档区内容替换为新节点的文档
8. 节点没有对应产出文档时，文档区显示"该步骤暂无产出文档"
9. 数据请求中鼠标光标变为 cursor: wait

## Tasks / Subtasks

- [ ] Task 1: 创建 DocMeta 组件 (AC: #2)
  - [ ] 创建 `web/src/components/DocMeta.vue`
  - [ ] Props: step (WorkflowStep), roleColor (string)
  - [ ] 水平排列展示：代理角色徽章（图标+名称，角色色背景）、命令代码（等宽字体）、时间徽章（warning 色）、难度徽章（success 色）
  - [ ] flex wrap 布局

- [ ] Task 2: 添加节点状态管理到 FlowView (AC: #1, #5, #6, #9)
  - [ ] 在 FlowView 中添加状态：activeStep (ref), visitedSteps (ref<Set>), docLoading (ref)
  - [ ] 点击核心节点时：设置 activeStep、添加到 visitedSteps、更新路由为 /flow/:role/:step
  - [ ] 传递 isActive 和 isVisited props 给 FlowNode
  - [ ] FlowNode 需添加 isVisited prop：角色色 40% 透明度边框

- [ ] Task 3: 实现文档加载和展示 (AC: #3, #4, #7, #8, #9)
  - [ ] 在 FlowView 中：点击节点后根据 step.outputs 查找对应文档
  - [ ] 调用 fetchDocument(path) 加载文档内容
  - [ ] 下方文档区：展示 DocMeta + DocRenderer
  - [ ] 无文档时显示"该步骤暂无产出文档"
  - [ ] 加载时 cursor: wait
  - [ ] 文档切换时自动滚动到文档区顶部

- [ ] Task 4: 文档匹配逻辑 (AC: #3, #8)
  - [ ] step.outputs 包含产出文件描述（如 "brainstorming session"、"prd"）
  - [ ] 需要将 outputs 模糊匹配到实际文档路径
  - [ ] 先调用 fetchDocuments() 获取文档列表，然后在本地匹配
  - [ ] 匹配策略：outputs 关键词与文档 path 或 title 做模糊匹配
  - [ ] 如果匹配到多个文档，显示第一个

- [ ] Task 5: 更新 FlowNode 支持 isVisited (AC: #6)
  - [ ] 修改 FlowNode.vue：添加 isVisited prop
  - [ ] 已浏览态：角色色边框 40% 透明度（非虚线）

- [ ] Task 6: 验证 (AC: #1-#9)
  - [ ] TypeScript 类型检查通过
  - [ ] Vite 构建成功
  - [ ] 点击节点后文档区显示元信息和文档内容
  - [ ] 节点激活/已浏览状态正确
  - [ ] 无文档节点显示提示
  - [ ] 路由正确更新

## Dev Notes

### 已建立的代码库
**Epic 2 完成后：**
- FlowView：fetchRoles() → RoleTab + FlowNode + FlowArrow，下方空状态占位
- FlowNode：支持 isActive/isUpstream/isDownstream props，emit click
- DocRenderer：markdown-it 渲染，prose prose-invert
- fetchDocument(path)、fetchDocuments() API 客户端已实现
- 路由已支持 `/flow/:role/:step?` 可选参数

### 文档匹配挑战
step.outputs 是描述性文本（如 "brainstorming session"、"prd"），不是文件路径。需要模糊匹配到实际文档。

**匹配策略建议：**
```typescript
function findDocumentForStep(step: WorkflowStep, docs: DocumentSummary[]): string | null {
  for (const output of step.outputs) {
    const keywords = output.toLowerCase().split(/\s+/)
    const match = docs.find(doc => {
      const target = (doc.path + ' ' + doc.title).toLowerCase()
      return keywords.some(kw => target.includes(kw))
    })
    if (match) return match.path
  }
  return null
}
```

### DocMeta 设计参考（来自 UX 设计规范）
水平排列徽章：
- 代理角色：角色色背景 15% + 角色色文字，如 "💻 Amelia"
- 命令：等宽字体代码样式，如 `bmad-dev-story`
- 时间：warning 色背景 10% + warning 色文字，如 "⏱ ~45min"
- 难度：success 色背景 10% + success 色文字，如 "中等"

### 文档区滚动
切换文档时需要滚动到文档区顶部。使用 `ref` 引用文档区 DOM 元素：
```typescript
const docArea = ref<HTMLElement>()
// 切换后
nextTick(() => docArea.value?.scrollIntoView({ behavior: 'smooth' }))
```

### 架构约束
- Vue 组件 PascalCase
- 不写 `<style>` 块
- 动态颜色用内联 style

### References
- [Source: ux-design-specification.md#Component Strategy] DocMeta 组件规格
- [Source: ux-design-specification.md#UX Consistency Patterns] 节点选择、已浏览标记、加载状态
- [Source: epics.md#Story 3.1] 原始故事定义

## Dev Agent Record

### Agent Model Used

### Debug Log References

### Completion Notes List

### File List

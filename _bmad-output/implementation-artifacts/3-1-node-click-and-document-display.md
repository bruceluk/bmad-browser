# Story 3.1: 节点点击与文档展示

Status: done

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

- [x] Task 1: 创建 DocMeta 组件 (AC: #2)
  - [x] DocMeta.vue: 代理徽章(角色色)、命令(等宽)、时间(warning)、必需(success)
  - [x] flex wrap 布局

- [x] Task 2: 节点状态管理 (AC: #1, #5, #6, #9)
  - [x] activeStep, visitedSteps(Set), docLoading 状态
  - [x] selectStep: 设置激活、添加已浏览、更新路由
  - [x] 传递 isActive/isVisited props

- [x] Task 3: 文档加载展示 (AC: #3, #4, #7, #8, #9)
  - [x] 点击节点 → findDocumentForStep → fetchDocument
  - [x] DocMeta + DocRenderer 展示
  - [x] 无文档显示"该步骤暂无产出文档"
  - [x] cursor: wait, scrollIntoView

- [x] Task 4: 文档匹配逻辑 (AC: #3, #8)
  - [x] fetchDocuments() 预加载文档列表
  - [x] outputs 关键词模糊匹配 path+title（过滤长度>2 的关键词）

- [x] Task 5: FlowNode isVisited (AC: #6)
  - [x] 添加 isVisited prop
  - [x] 已浏览态：角色色 + '66' 透明度边框

- [x] Task 6: 验证 (AC: #1-#9)
  - [x] TypeScript 类型检查通过
  - [x] Vite 构建成功（581ms，8 个输出文件）

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

Claude Opus 4.6 (1M context)

### Debug Log References

- DocRenderer 被 Vite 自动拆分为独立 chunk（103KB，含 markdown-it）
- Promise.all 并行加载 roles + documents 提升启动速度

### Completion Notes List

- ✅ DocMeta 组件：4 个徽章（代理、命令、时间、必需），flex wrap
- ✅ FlowView 完整重写：activeStep/visitedSteps 状态管理 + 文档加载 + 匹配
- ✅ FlowNode 添加 isVisited：角色色 66 透明度边框
- ✅ 文档匹配：outputs 关键词模糊匹配 path+title
- ✅ 空状态/无文档/加载中三种文档区状态
- ✅ scrollIntoView 自动滚动到文档区

### File List

- bmad-viewer/web/src/components/DocMeta.vue (new)
- bmad-viewer/web/src/components/FlowNode.vue (modified - added isVisited)
- bmad-viewer/web/src/views/FlowView.vue (modified - complete rewrite with doc loading)

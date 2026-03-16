# Story 3.2: 命令与代理角色信息展示

Status: ready-for-dev

## Story

As a 浏览者，
I want 在流程图和文档详情中看到每个步骤的命令、代理角色和产出文件信息，
so that 我知道每一步用什么命令、由哪个 AI 代理负责、产出什么文件。

## Acceptance Criteria

1. 每个流程图节点显示 BMAD 命令名称（FR6）
2. 每个节点显示对应的 AI 代理角色图标和名称（FR7）
3. 点击节点后，DocMeta 展示该命令产出的文件列表（FR8）
4. 产出文件可点击，跳转到对应的产出文档查看（FR9）

## Tasks / Subtasks

- [ ] Task 1: 增强 DocMeta 展示产出文件列表 (AC: #3, #4)
  - [ ] 修改 DocMeta.vue：在现有徽章下方添加产出文件列表
  - [ ] 从 step.outputs 获取产出描述列表
  - [ ] 使用 findDocumentForStep 同样的匹配逻辑找到对应文档路径
  - [ ] 每个产出文件显示为可点击链接
  - [ ] 点击产出文件 emit navigate 事件（传递文档路径）

- [ ] Task 2: FlowView 处理产出文件导航 (AC: #4)
  - [ ] 在 FlowView 中监听 DocMeta 的 navigate 事件
  - [ ] 导航到 /doc/:path 查看产出文档（复用 DocView）

- [ ] Task 3: 验证 FlowNode 已展示命令和代理 (AC: #1, #2)
  - [ ] 确认 FlowNode 已显示 step.command（等宽字体）— Story 2.3 已实现
  - [ ] 确认 FlowNode 已显示 step.agentIcon — Story 2.3 已实现
  - [ ] 此为验证任务，无需代码修改（除非发现缺失）

- [ ] Task 4: 验证 (AC: #1-#4)
  - [ ] TypeScript 类型检查通过
  - [ ] Vite 构建成功
  - [ ] 点击节点后 DocMeta 显示产出文件列表
  - [ ] 产出文件可点击跳转到文档详情

## Dev Notes

### 已建立的代码库
**Story 3.1 完成：**
- DocMeta.vue 已展示代理徽章、命令、时间、必需标签
- FlowView 已有 selectStep → findDocumentForStep → fetchDocument 完整链路
- FlowNode 已显示 command（等宽字体）和 agentIcon
- allDocs（DocumentSummary[]）已在 FlowView 中预加载
- findDocumentForStep 已实现模糊匹配

### DocMeta 增强方案
在现有徽章行下方添加产出文件区域：
```vue
<!-- Outputs -->
<div v-if="outputs.length > 0" class="mt-3">
  <span class="text-xs opacity-50">产出文件：</span>
  <span v-for="..." class="text-sm cursor-pointer" @click="...">
    {{ output.title || output.path }}
  </span>
</div>
```

### 需要传递 allDocs 给 DocMeta
DocMeta 需要 allDocs 列表来匹配 outputs 到实际文档路径。两种方案：
1. 将 allDocs 作为 prop 传递给 DocMeta
2. 在 FlowView 中预计算 matched outputs，传递给 DocMeta

推荐方案 2：FlowView 预计算，DocMeta 只接收已匹配的结果。

新增类型：
```typescript
interface MatchedOutput {
  label: string    // 产出描述（来自 step.outputs）
  docPath: string | null  // 匹配到的文档路径（null 表示无对应文档）
}
```

### 架构约束
- Vue 组件 PascalCase
- 不写 `<style>` 块
- 动态颜色用内联 style

### References
- [Source: epics.md#Story 3.2] 原始故事定义
- [Source: ux-design-specification.md#Component Strategy] DocMeta 组件规格

## Dev Agent Record

### Agent Model Used

### Debug Log References

### Completion Notes List

### File List

# Story 3.3: 上下游延伸导航

Status: done

## Story

As a 浏览者，
I want 点击流程图中的上下游虚线节点后自动切换到对应角色的流程，
so that 我能自然地了解上下游同事的工作流程。

## Acceptance Criteria

1. 在开发者流程图中点击上游虚线节点（如 PM 的"创建 PRD"），角色 Tab 自动切换到产品经理（UX-DR9）
2. 流程图更新为产品经理的流程，对应节点自动激活并高亮
3. 下方文档区展示该节点的产出文档
4. 在产品经理流程图中点击下游虚线节点（如开发者的"故事开发"），角色 Tab 自动切换到开发者，对应节点自动激活
5. 通过上下游延伸导航切换角色后，路由更新为 `/flow/:newRole/:step`
6. 之前角色的已浏览状态保留（刷新页面后重置）

## Tasks / Subtasks

- [x] Task 1: 上下游节点点击跳转 (AC: #1, #2, #3, #4)
  - [x] upstream/downstream FlowNode 添加 @click="selectContextStep(step)"
  - [x] selectContextStep: findRoleForStep → 切换 currentRole → nextTick → selectStep

- [x] Task 2: 角色查找逻辑 (AC: #1, #4)
  - [x] findRoleForStep: 遍历 roles，查找 steps 中包含相同 code 的角色

- [x] Task 3: 已浏览状态跨角色保持 (AC: #6)
  - [x] visitedSteps 是页面级 Set，switchRole 不清空，selectContextStep 也不清空 ✅

- [x] Task 4: 路由更新 (AC: #5)
  - [x] selectStep 内 router.replace 自动更新为 /flow/:newRole/:step ✅

- [x] Task 5: 验证 (AC: #1-#6)
  - [x] TypeScript 类型检查通过
  - [x] Vite 构建成功（593ms）

## Dev Notes

### 已建立的代码库
**Story 3.1+3.2 完成后：**
- FlowView 已有：selectStep（激活节点+加载文档）、switchRole（切换角色）、visitedSteps（已浏览集合）
- 上下游 FlowNode 当前没有 @click 处理（Story 2.3 只给核心节点绑定了 click）
- 上下游节点已渲染，有 isUpstream/isDownstream 标记

### 实现方案
点击上下游节点时，需要：
1. 确定该节点属于哪个角色（通过 findRoleForStep）
2. 切换到该角色（更新 currentRole + 路由）
3. 激活该节点并加载文档（调用 selectStep）

关键：上下游节点的 step 对象来自 currentRoleFlow.upstream/downstream，这些 step 实际上是其他角色 steps 的子集。可以用 step.code 在目标角色的 steps 中匹配。

```typescript
async function selectUpstreamDownstreamStep(step: WorkflowStep) {
  const targetRole = findRoleForStep(step)
  if (!targetRole) return

  currentRole.value = targetRole.role
  router.replace({ name: 'flow', params: { role: targetRole.role, step: step.code } })

  // Wait for currentRoleFlow to update, then select step
  await nextTick()
  const targetStep = targetRole.steps.find(s => s.code === step.code)
  if (targetStep) selectStep(targetStep)
}

function findRoleForStep(step: WorkflowStep): RoleFlow | null {
  return roles.value.find(r => r.steps.some(s => s.code === step.code)) || null
}
```

### 注意事项
- switchRole 当前清空 activeStep 和 currentDoc，但上下游跳转需要保留激活状态
- 不要调用 switchRole，直接修改 currentRole 并调用 selectStep
- visitedSteps 在 switchRole 中不清空（当前已经是这样）

### References
- [Source: ux-design-specification.md#UX Consistency Patterns] 上下游延伸交互
- [Source: ux-design-specification.md#User Journey Flows] 上下游延伸模式
- [Source: epics.md#Story 3.3] 原始故事定义

## Dev Agent Record

### Agent Model Used

Claude Opus 4.6 (1M context)

### Debug Log References

- selectContextStep 使用 nextTick 等待 currentRoleFlow computed 更新后再 selectStep
- 不调用 switchRole（会清空 activeStep），直接修改 currentRole

### Completion Notes List

- ✅ selectContextStep：上下游节点跳转（查找目标角色 → 切换 → 激活节点 → 加载文档）
- ✅ findRoleForStep：通过 step.code 在 roles.steps 中匹配
- ✅ upstream/downstream FlowNode 绑定 @click
- ✅ visitedSteps 跨角色保持（页面级 Set，不随角色切换清空）
- ✅ 路由自动更新为 /flow/:newRole/:step

### File List

- bmad-viewer/web/src/views/FlowView.vue (modified - added selectContextStep, findRoleForStep, upstream/downstream @click)

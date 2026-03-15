# Story 2.1: CSV 解析与角色流程 API

Status: review

## Story

As a 浏览者，
I want 应用能解析 BMAD 工作流数据并按角色分组返回，
so that 前端可以展示每个角色的工作流程图。

## Acceptance Criteria

1. Go 服务启动时解析 `bmad-help.csv`，提取工作流名称、命令、代理角色、阶段、描述等字段（FR17）
2. 将工作流数据按角色（开发者/产品经理/测试人员）分组，构建 RoleFlow 结构
3. 每个角色的流程包含核心步骤、上游步骤和下游步骤
4. 数据缓存到内存
5. `GET /api/roles` 返回角色列表，每个角色包含流程步骤（名称、命令、代理、预期时间）、上游步骤、下游步骤
6. `GET /api/workflows` 返回 CSV 解析后的完整工作流数据列表

## Tasks / Subtasks

- [x] Task 1: 定义工作流数据模型 (AC: #2, #5)
  - [x] 创建 `server/model/workflow.go`：WorkflowStep（含 Sequence 字段）+ RoleFlow（含 Label 字段）
  - [x] JSON tag 使用 camelCase

- [x] Task 2: 实现 CSV 解析器 (AC: #1)
  - [x] 创建 `server/parser/csv_parser.go`：ParseCSV 函数
  - [x] Go 标准库 encoding/csv + LazyQuotes（CSV 含裸引号）
  - [x] 仅解析 module=bmm 的行，跳过 header
  - [x] 命令行参数 -csv 指定路径

- [x] Task 3: 实现角色-流程映射逻辑 (AC: #2, #3)
  - [x] BuildRoleFlows：按 isPMStep/isDevStep/isQAStep 分类
  - [x] 开发者：3-solutioning + 4-implementation（排除 QA/SM/Retrospective）
  - [x] PM：1-analysis + 2-planning
  - [x] QA：name 或 agent 含 "qa"
  - [x] 上下游：lastN/firstN 取相邻角色的首尾步骤
  - [x] 按 sequence 排序

- [x] Task 4: 实现工作流 API handler (AC: #5, #6)
  - [x] `server/handler/workflows.go`：HandleRoles + HandleWorkflows

- [x] Task 5: 集成到 main.go (AC: #1, #4)
  - [x] 添加 -csv 参数，启动时 ParseCSV + BuildRoleFlows
  - [x] 注册 /api/roles 和 /api/workflows
  - [x] 日志：31 BMM workflow steps, 3 role flows

- [x] Task 6: 更新前端类型定义 (AC: #5)
  - [x] types/index.ts：WorkflowStep + RoleFlow 接口
  - [x] api/client.ts：fetchRoles() + fetchWorkflows()

- [x] Task 7: 验证 (AC: #1-#6)
  - [x] Go 编译成功
  - [x] /api/workflows 返回 31 个工作流
  - [x] /api/roles 返回 3 角色：开发者(7步+2上游+1下游)、PM(9步+0+2)、QA(1步+2+0)
  - [x] 前端 TypeScript 类型检查通过

## Dev Notes

### 已建立的代码库
**Epic 1 完成后的项目状态：**
- Go 后端：main.go + handler/documents.go + parser/scanner.go + parser/markdown_parser.go + model/document.go + handler/static.go + embed.go
- 前端：Vue 3 + TypeScript + Tailwind + markdown-it，HomeView + DocView + DocRenderer
- 构建：`make build` 生成单一可执行文件
- 命令行参数：-port, -dir

### CSV 文件结构
`bmad-help.csv` 列定义（按索引）：
```
0: module        # 模块名（bmm, wds, tea, gds, cis, core）
1: phase         # 阶段（1-analysis, 2-planning, 3-solutioning, 4-implementation, anytime）
2: name          # 工作流名称
3: code          # 代码缩写
4: sequence      # 排序序号
5: workflow-file # 工作流文件路径
6: command       # BMAD 命令名
7: required      # 是否必需（true/false）
8: agent-name    # 代理内部名
9: agent-command # 代理命令
10: agent-display-name  # 代理显示名
11: agent-title         # 代理标题（含 emoji）
12: options             # 选项
13: description         # 描述
14: output-location     # 输出位置
15: outputs             # 产出物
```

### 角色-流程映射规则

**仅解析 module=bmm 的工作流**（BMAD Viewer 展示 BMM 模块的工作流）。

**三个角色及其核心阶段：**

| 角色 | 角色色 | 核心阶段 | 说明 |
|------|--------|----------|------|
| 开发者 | #4fc3f7 | 3-solutioning, 4-implementation | 架构设计、故事开发、代码审查 |
| 产品经理 | #81c784 | 1-analysis, 2-planning | 产品简报、PRD、UX 设计、史诗故事 |
| 测试人员 | #b39ddb | 4-implementation（QA 相关） | QA 自动化测试 |

**核心步骤筛选：**
- 开发者：phase 包含 "3-solutioning" 或 "4-implementation"（排除 QA/SM 相关）
  - 包含：Create Architecture, Create Epics and Stories, Dev Story, Code Review
  - 排除：Sprint Planning, Sprint Status, Create Story, QA Automation, Retrospective（这些是 SM/QA 角色）
- 产品经理：phase 包含 "1-analysis" 或 "2-planning"
  - 包含：Brainstorm, Create Brief, Create PRD, Validate PRD, Create UX
- 测试人员：name 包含 "QA" 或 agent-name 包含 "qa"
  - 包含：QA Automation Test

**上下游逻辑：**
- 开发者上游：产品经理的最后几步（Create PRD、Create UX）
- 开发者下游：测试人员的步骤（QA Automation）
- 产品经理下游：开发者的前几步（Create Architecture、Create Epics）
- 测试人员上游：开发者的后几步（Dev Story、Code Review）

**简化实现建议：** 上下游可以硬编码关键步骤名称映射，不需要复杂的自动推断。

### 数据模型参考（来自架构文档）
```go
type WorkflowStep struct {
    Name        string   `json:"name"`
    Code        string   `json:"code"`
    Command     string   `json:"command"`
    AgentName   string   `json:"agentName"`
    AgentIcon   string   `json:"agentIcon"`
    Phase       string   `json:"phase"`
    Required    bool     `json:"required"`
    Description string   `json:"description"`
    Outputs     []string `json:"outputs"`
    Duration    string   `json:"duration"`
}

type RoleFlow struct {
    Role       string         `json:"role"`
    RoleColor  string         `json:"roleColor"`
    Steps      []WorkflowStep `json:"steps"`
    Upstream   []WorkflowStep `json:"upstream"`
    Downstream []WorkflowStep `json:"downstream"`
}
```

### 架构约束
- Go 文件名用 `snake_case`
- JSON 字段用 `camelCase`
- API 响应直接返回数据，不加包装层
- 使用 Go 标准库 `encoding/csv`，不引入第三方 CSV 库

### 文件结构
```
server/
├── main.go                 (modify: 添加 CSV 解析和路由)
├── model/
│   └── workflow.go         (new: WorkflowStep, RoleFlow)
├── parser/
│   └── csv_parser.go       (new: CSV 解析 + 角色映射)
├── handler/
│   └── workflows.go        (new: /api/roles, /api/workflows)
web/src/
├── types/index.ts          (modify: 添加 WorkflowStep, RoleFlow)
├── api/client.ts           (modify: 添加 fetchRoles, fetchWorkflows)
```

### References
- [Source: architecture.md#Data Architecture] WorkflowStep, RoleFlow 数据模型
- [Source: architecture.md#API & Communication Patterns] /api/roles, /api/workflows 端点
- [Source: ux-design-specification.md#Core User Experience] 三个角色入口和角色色
- [Source: epics.md#Story 2.1] 原始故事定义

## Dev Agent Record

### Agent Model Used

Claude Opus 4.6 (1M context)

### Debug Log References

- CSV LazyQuotes 修复：bmad-help.csv 第 59 行含裸引号，标准解析失败，启用 LazyQuotes
- estimateDuration 基于步骤名称硬编码预估时间（CSV 无此字段）

### Completion Notes List

- ✅ WorkflowStep + RoleFlow 数据模型（含 Sequence 排序字段和 Label 中文标签）
- ✅ CSV 解析器：31 个 BMM 工作流步骤，LazyQuotes 处理特殊字符
- ✅ 角色映射：3 角色分类逻辑 + 上下游关系 + 按 sequence 排序
- ✅ API handler：/api/roles 和 /api/workflows
- ✅ 前端类型和 API 客户端同步更新
- ✅ 端到端验证通过

### File List

- bmad-viewer/server/model/workflow.go (new)
- bmad-viewer/server/parser/csv_parser.go (new)
- bmad-viewer/server/handler/workflows.go (new)
- bmad-viewer/server/main.go (modified)
- bmad-viewer/web/src/types/index.ts (modified)
- bmad-viewer/web/src/api/client.ts (modified)

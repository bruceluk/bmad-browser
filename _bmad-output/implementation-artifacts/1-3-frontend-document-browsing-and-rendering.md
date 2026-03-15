# Story 1.3: 前端文档浏览与渲染

Status: ready-for-dev

## Story

As a 浏览者，
I want 在浏览器中查看文档列表并阅读渲染后的 Markdown 内容，
so that 我可以浏览 BMAD 项目的产出物。

## Acceptance Criteria

1. 用户打开浏览器访问应用，页面加载后显示文档列表（FR1）
2. 页面使用深色主题（UX-DR1：背景色 `#1a1a2e`、文字色 `#e0e0e0`）
3. 使用中文字体栈和等宽代码字体（UX-DR10）
4. 用户点击某个文档，展示该文档的完整渲染内容（FR4）
5. Markdown 渲染为 HTML，支持标题、列表、代码块、表格（UX-DR7）
6. 使用 Tailwind Typography prose 样式，深色主题适配
7. 切换文档无需整页刷新（FR5，SPA 路由切换）

## Tasks / Subtasks

- [ ] Task 1: 定义 TypeScript 类型 (AC: #1)
  - [ ] 创建 `web/src/types/index.ts`，定义与 Go 后端对应的类型：
    - `Document`（path, title, content, frontmatter, phase）
    - `DocumentSummary`（path, title, phase）

- [ ] Task 2: 创建 API 客户端 (AC: #1)
  - [ ] 创建 `web/src/api/client.ts`
  - [ ] 实现 `fetchDocuments(): Promise<DocumentSummary[]>` — 调用 `GET /api/documents`
  - [ ] 实现 `fetchDocument(path: string): Promise<Document>` — 调用 `GET /api/documents/:path`
  - [ ] 使用 fetch API，错误时 throw Error

- [ ] Task 3: 安装 markdown-it (AC: #5)
  - [ ] `npm install markdown-it`
  - [ ] `npm install -D @types/markdown-it`

- [ ] Task 4: 实现 DocRenderer 组件 (AC: #4, #5, #6)
  - [ ] 创建 `web/src/components/DocRenderer.vue`
  - [ ] Props: `content: string`（Markdown 原始内容）
  - [ ] 使用 markdown-it 将 Markdown 转为 HTML
  - [ ] 使用 Tailwind Typography `prose` 类渲染，深色主题适配（`prose-invert`）
  - [ ] 内容区最大宽度 800px，居中，左右内边距 48px
  - [ ] 支持标题（h1-h3）、列表、代码块、表格、引用块

- [ ] Task 5: 改造 HomeView 为文档列表页 (AC: #1, #2, #3, #7)
  - [ ] 修改 `web/src/views/HomeView.vue`
  - [ ] 页面加载时调用 `fetchDocuments()` 获取文档列表
  - [ ] 渲染文档列表：每项显示标题和阶段（phase）
  - [ ] 使用深色主题样式（CSS 变量已在 style.css 中定义）
  - [ ] 点击文档项，路由跳转到文档详情页

- [ ] Task 6: 创建文档详情视图 (AC: #4, #5, #6, #7)
  - [ ] 创建 `web/src/views/DocView.vue`
  - [ ] 路由参数接收文档 path（使用 catch-all 路由 `path(.*)`）
  - [ ] 页面加载时调用 `fetchDocument(path)` 获取文档
  - [ ] 使用 DocRenderer 组件渲染文档内容
  - [ ] 显示文档标题和返回列表的链接
  - [ ] 加载时鼠标光标变为 `cursor: wait`

- [ ] Task 7: 更新路由配置 (AC: #7)
  - [ ] 修改 `web/src/router/index.ts`
  - [ ] 添加文档详情路由：`/doc/:path(.*)`（catch-all 匹配子路径）
  - [ ] 懒加载 DocView 组件

- [ ] Task 8: 验证 (AC: #1-#7)
  - [ ] 前端 TypeScript 类型检查通过
  - [ ] Vite 构建成功
  - [ ] 启动 Go 后端 + Vite 前端，文档列表正常显示
  - [ ] 点击文档，Markdown 渲染正确（标题、列表、代码块）
  - [ ] 切换文档无整页刷新

## Dev Notes

### 前序故事已建立的基础
**Story 1.1:**
- Vue 3 + TypeScript + Vue Router + Vite 已初始化
- Tailwind CSS 4 + Typography 插件已配置
- 深色主题 CSS 变量已在 `style.css` 中定义（12 个变量）
- 中文字体栈 + 等宽代码字体已配置
- Vite API 代理配置（`/api` → `localhost:8080`）
- 空目录 `components/`、`views/`、`api/`、`types/` 已创建
- HomeView.vue 当前只有占位内容

**Story 1.2:**
- Go 后端 API 已实现：
  - `GET /api/documents` — 返回 DocumentSummary 数组
  - `GET /api/documents/:path` — 返回完整 Document（含 content）
- API 响应格式为 camelCase JSON

### 架构约束
- Vue 组件文件名用 `PascalCase`（如 `DocRenderer.vue`、`DocView.vue`）
- TypeScript 文件名用 `camelCase`（如 `client.ts`、`index.ts`）
- 不写 `<style>` 块，所有样式用 Tailwind utility classes
- 状态管理用 Vue ref/reactive，不用 Pinia
- 不引入架构文档未列出的依赖（markdown-it 是本故事需要的）

### Markdown 渲染注意事项
- markdown-it 默认不启用 HTML 标签渲染，保持默认（安全）
- Tailwind Typography 的 `prose` 类提供完整的排版样式
- 深色主题使用 `prose-invert` 修饰符
- 代码块需要等宽字体（已在 style.css 中全局设置）

### Vue Router catch-all 路由
文档路径包含 `/`（如 `planning-artifacts/prd.md`），需要 catch-all：
```typescript
{
  path: '/doc/:path(.*)',
  name: 'doc',
  component: () => import('@/views/DocView.vue'),
}
```

### API 客户端模式
```typescript
// api/client.ts
const API_BASE = '/api'

export async function fetchDocuments(): Promise<DocumentSummary[]> {
  const res = await fetch(`${API_BASE}/documents`)
  if (!res.ok) throw new Error(`Failed to fetch documents: ${res.status}`)
  return res.json()
}

export async function fetchDocument(path: string): Promise<Document> {
  const res = await fetch(`${API_BASE}/documents/${path}`)
  if (!res.ok) throw new Error(`Failed to fetch document: ${res.status}`)
  return res.json()
}
```

### 布局参考（来自 UX 设计规范）
- 内容区最大宽度：800px，居中
- 内容区左右内边距：48px
- 正文字号：16px / 1.75 行高
- 标题字号：h1=28px, h2=22px, h3=18px

### References
- [Source: architecture.md#Frontend Architecture] 状态管理和路由设计
- [Source: architecture.md#Implementation Patterns] 命名规范
- [Source: ux-design-specification.md#Component Strategy] DocRenderer 组件规格
- [Source: ux-design-specification.md#Visual Design Foundation] 排版和间距系统
- [Source: epics.md#Story 1.3] 原始故事定义

## Dev Agent Record

### Agent Model Used

### Debug Log References

### Completion Notes List

### File List

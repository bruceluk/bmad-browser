# Story 1.3: 前端文档浏览与渲染

Status: review

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

- [x] Task 1: 定义 TypeScript 类型 (AC: #1)
  - [x] 创建 `web/src/types/index.ts`：Document 和 DocumentSummary 接口

- [x] Task 2: 创建 API 客户端 (AC: #1)
  - [x] 创建 `web/src/api/client.ts`：fetchDocuments() + fetchDocument(path)
  - [x] 使用 fetch API，错误时 throw Error

- [x] Task 3: 安装 markdown-it (AC: #5)
  - [x] `npm install markdown-it @types/markdown-it`（8+3 packages）

- [x] Task 4: 实现 DocRenderer 组件 (AC: #4, #5, #6)
  - [x] 创建 `web/src/components/DocRenderer.vue`
  - [x] Props: content string，使用 markdown-it 渲染
  - [x] prose prose-invert 深色主题，max-w-[800px] 居中，px-12 内边距

- [x] Task 5: 改造 HomeView 为文档列表页 (AC: #1, #2, #3, #7)
  - [x] 页面加载时调用 fetchDocuments()
  - [x] 渲染文档列表（标题 + 阶段标签），点击跳转到详情页
  - [x] 深色主题样式，hover 边框变色

- [x] Task 6: 创建文档详情视图 (AC: #4, #5, #6, #7)
  - [x] 创建 `web/src/views/DocView.vue`
  - [x] catch-all 路由参数接收 path，加载文档内容
  - [x] DocRenderer 渲染，返回列表链接，cursor: wait 加载状态

- [x] Task 7: 更新路由配置 (AC: #7)
  - [x] 添加 `/doc/:path(.*)` catch-all 路由，懒加载 DocView

- [x] Task 8: 验证 (AC: #1-#7)
  - [x] TypeScript 类型检查通过
  - [x] Vite 构建成功（523ms，6 个输出文件）
  - [x] DocView chunk 104KB（含 markdown-it）

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

Claude Opus 4.6 (1M context)

### Debug Log References

- markdown-it + @types/markdown-it 安装成功（11 packages total）
- Tailwind Typography prose-invert 用于深色主题适配

### Completion Notes List

- ✅ TypeScript 类型定义：Document + DocumentSummary 接口
- ✅ API 客户端：fetchDocuments + fetchDocument，fetch API + error throw
- ✅ DocRenderer 组件：markdown-it 渲染 + prose prose-invert + 800px 居中
- ✅ HomeView 改造为文档列表页：加载文档列表、标题+阶段标签、点击跳转
- ✅ DocView 文档详情页：catch-all 路由、文档加载、DocRenderer 渲染、返回链接
- ✅ Vue Router 更新：添加 /doc/:path(.*) 路由
- ✅ TypeScript 类型检查通过，Vite 构建成功

### File List

- bmad-viewer/web/src/types/index.ts (new)
- bmad-viewer/web/src/api/client.ts (new)
- bmad-viewer/web/src/components/DocRenderer.vue (new)
- bmad-viewer/web/src/views/HomeView.vue (modified)
- bmad-viewer/web/src/views/DocView.vue (new)
- bmad-viewer/web/src/router/index.ts (modified)
- bmad-viewer/web/package.json (modified - added markdown-it)
- bmad-viewer/web/package-lock.json (modified)

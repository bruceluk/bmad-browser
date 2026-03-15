# Story 1.1: 项目初始化与开发环境搭建

Status: ready-for-dev

## Story

As a 开发者（Lu），
I want 初始化 monorepo 项目结构（web/ + server/），配置开发环境，
so that 前后端可以独立开发并通过 Vite 代理联调。

## Acceptance Criteria

1. 克隆仓库后执行 `npm install`（web/）和 `go mod download`（server/）可成功安装依赖
2. `cd web && npm run dev` 可启动 Vite 开发服务器
3. `cd server && go run .` 可启动 Go 后端（返回简单健康检查响应）
4. Vite 代理 `/api/*` 请求到 Go 后端（vite.config.ts 配置）
5. 项目结构符合架构文档定义
6. Tailwind CSS 4 + `@tailwindcss/typography` 已配置并生效
7. Makefile 包含 `make dev` 和 `make build` 目标

## Tasks / Subtasks

- [ ] Task 1: 创建项目根目录和基础文件 (AC: #5)
  - [ ] 创建 `bmad-viewer/` 根目录
  - [ ] 创建 `.gitignore`（Go + Node 合并规则：node_modules/、dist/、*.exe）
  - [ ] 创建 `Makefile`（dev 和 build 目标，详见下方 Dev Notes）

- [ ] Task 2: 初始化 Vue 3 前端 (AC: #1, #2, #5)
  - [ ] 运行 `npm create vue@latest web -- --typescript --router --no-pinia --no-vitest --no-e2e --no-eslint`
  - [ ] 清理脚手架生成的示例文件（删除 HelloWorld.vue、TheWelcome.vue 等）
  - [ ] 创建空的目录结构：`src/components/`、`src/views/`、`src/api/`、`src/types/`

- [ ] Task 3: 配置 Tailwind CSS (AC: #6)
  - [ ] 安装 `tailwindcss` 和 `@tailwindcss/typography`
  - [ ] 创建 `tailwind.config.js`，配置 content 路径和 typography 插件
  - [ ] 在 `src/style.css` 中添加 `@tailwind` 指令
  - [ ] 添加深色主题 CSS 变量（详见下方色彩系统）
  - [ ] 验证 Tailwind 类在组件中生效

- [ ] Task 4: 配置 Vite API 代理 (AC: #4)
  - [ ] 在 `vite.config.ts` 中添加 proxy 配置：`/api` → `http://localhost:8080`

- [ ] Task 5: 初始化 Go 后端 (AC: #1, #3, #5)
  - [ ] 创建 `server/` 目录
  - [ ] 运行 `go mod init bmad-viewer/server`
  - [ ] 创建 `server/main.go`：最小 HTTP server，监听 8080 端口
  - [ ] 添加健康检查端点 `GET /api/health` 返回 `{"status": "ok"}`
  - [ ] 创建空的目录结构：`handler/`、`parser/`、`model/`

- [ ] Task 6: 创建 Makefile (AC: #7)
  - [ ] `make dev`：并行启动前后端开发服务器
  - [ ] `make build`：前端构建 + Go 编译（Go embed 在 Story 1.4 实现，此处先用占位）

- [ ] Task 7: 验证端到端联调 (AC: #2, #3, #4)
  - [ ] 启动 Go 后端，确认 `/api/health` 返回正常
  - [ ] 启动 Vite 前端，确认页面加载
  - [ ] 在前端调用 `/api/health`，确认代理正常工作

## Dev Notes

### 技术栈版本要求
- Go 1.22+
- Node.js 18+（支持 npm create vue@latest）
- Vue 3（最新稳定版，通过 create-vue 安装）
- TypeScript 严格模式
- Vite（create-vue 默认）
- Tailwind CSS 4
- `@tailwindcss/typography` 插件

### 色彩系统 CSS 变量
在 `src/style.css` 中定义（来源：UX 设计规范）：
```css
:root {
  --bg: #1a1a2e;
  --surface: #16213e;
  --border: #2a2a4a;
  --text: #e0e0e0;
  --text-secondary: #a0a0b0;
  --dev: #4fc3f7;
  --pm: #81c784;
  --qa: #b39ddb;
  --success: #66bb6a;
  --warning: #ffa726;
  --error: #ef5350;
  --link: #64b5f6;
}
```

### 字体系统
```css
body {
  font-family: "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", sans-serif;
}
code, pre {
  font-family: "JetBrains Mono", "Fira Code", monospace;
}
```

### Makefile 内容参考
```makefile
.PHONY: dev build clean

dev:
	@echo "Starting dev servers..."
	cd web && npm run dev &
	cd server && go run . &
	wait

build:
	cd web && npm run build
	cd server && go build -o ../bmad-viewer .

clean:
	rm -rf web/dist bmad-viewer
```

### Vite 代理配置
```typescript
// vite.config.ts
export default defineConfig({
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
    },
  },
})
```

### Go 最小 HTTP Server
```go
// server/main.go
package main

import (
    "encoding/json"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
    })

    log.Println("BMAD Viewer server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### 架构约束（必须遵守）
- Go 文件名用 `snake_case`
- Vue 组件文件名用 `PascalCase`
- TypeScript 文件名用 `camelCase`
- 不引入架构文档未列出的第三方依赖
- 前端不写 `<style>` 块，所有样式用 Tailwind utility classes
- API 响应直接返回数据，不加包装层
- JSON 字段用 `camelCase`

### Project Structure Notes

最终项目结构（本故事需创建）：
```
bmad-viewer/
├── .gitignore
├── Makefile
├── web/
│   ├── index.html
│   ├── package.json
│   ├── tsconfig.json
│   ├── vite.config.ts
│   ├── tailwind.config.js
│   └── src/
│       ├── main.ts
│       ├── App.vue
│       ├── style.css
│       ├── router/
│       │   └── index.ts
│       ├── views/         (空目录，后续故事填充)
│       ├── components/    (空目录，后续故事填充)
│       ├── api/           (空目录，后续故事填充)
│       └── types/         (空目录，后续故事填充)
└── server/
    ├── go.mod
    ├── main.go
    ├── handler/           (空目录，Story 1.2 填充)
    ├── parser/            (空目录，Story 1.2 填充)
    └── model/             (空目录，Story 1.2 填充)
```

### References

- [Source: architecture.md#Project Structure] 完整目录结构定义
- [Source: architecture.md#Starter Template Evaluation] 初始化命令
- [Source: architecture.md#Implementation Patterns] 命名规范和编码约束
- [Source: ux-design-specification.md#Visual Design Foundation] 色彩系统和字体系统
- [Source: ux-design-specification.md#Design System Foundation] Tailwind CSS 4 选型

## Dev Agent Record

### Agent Model Used

### Debug Log References

### Completion Notes List

### File List

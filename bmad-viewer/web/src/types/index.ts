export interface Document {
  path: string
  title: string
  content: string
  frontmatter: Record<string, unknown>
  phase: string
}

export interface DocumentSummary {
  path: string
  title: string
  phase: string
}

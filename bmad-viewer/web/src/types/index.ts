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

export interface WorkflowStep {
  name: string
  code: string
  command: string
  agentName: string
  agentIcon: string
  phase: string
  required: boolean
  description: string
  outputs: string[]
  duration: string
  sequence: number
}

export interface RoleFlow {
  role: string
  roleColor: string
  label: string
  steps: WorkflowStep[]
  upstream: WorkflowStep[]
  downstream: WorkflowStep[]
}

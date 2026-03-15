import type { Document, DocumentSummary } from '@/types'

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

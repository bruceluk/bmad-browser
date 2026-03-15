<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { fetchDocuments } from '@/api/client'
import type { DocumentSummary } from '@/types'

const router = useRouter()
const docs = ref<DocumentSummary[]>([])
const isLoading = ref(true)

onMounted(async () => {
  try {
    docs.value = await fetchDocuments()
  } catch (e) {
    console.error(e)
  } finally {
    isLoading.value = false
  }
})

function goToDoc(path: string) {
  router.push({ name: 'doc', params: { path } })
}

function phaseLabel(phase: string): string {
  const labels: Record<string, string> = {
    analysis: '分析',
    planning: '规划',
    implementation: '实施',
    other: '其他',
  }
  return labels[phase] || phase
}
</script>

<template>
  <div class="min-h-screen" :class="{ 'cursor-wait': isLoading }">
    <div class="max-w-[800px] mx-auto px-12 py-12">
      <h1 class="text-3xl font-bold mb-2">BMAD Viewer</h1>
      <p class="mb-8 opacity-60">浏览 BMAD 项目的产出文档</p>

      <div v-if="docs.length === 0 && !isLoading" class="opacity-40 text-center py-12">
        暂无文档
      </div>

      <ul class="space-y-2">
        <li
          v-for="doc in docs"
          :key="doc.path"
          class="flex items-center justify-between px-4 py-3 rounded-xl cursor-pointer transition-all duration-150"
          style="border: 1px solid var(--border);"
          @mouseenter="($event.currentTarget as HTMLElement).style.borderColor = 'var(--link)'"
          @mouseleave="($event.currentTarget as HTMLElement).style.borderColor = 'var(--border)'"
          @click="goToDoc(doc.path)"
        >
          <span class="font-medium">{{ doc.title || doc.path }}</span>
          <span
            class="text-xs px-2 py-0.5 rounded opacity-60"
            style="background-color: var(--surface);"
          >
            {{ phaseLabel(doc.phase) }}
          </span>
        </li>
      </ul>
    </div>
  </div>
</template>

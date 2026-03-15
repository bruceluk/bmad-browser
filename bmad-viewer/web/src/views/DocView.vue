<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { fetchDocument } from '@/api/client'
import DocRenderer from '@/components/DocRenderer.vue'
import type { Document } from '@/types'

const route = useRoute()
const router = useRouter()
const doc = ref<Document | null>(null)
const isLoading = ref(true)
const error = ref('')

async function loadDoc(path: string) {
  isLoading.value = true
  error.value = ''
  try {
    doc.value = await fetchDocument(path)
  } catch (e) {
    error.value = '文档加载失败'
    console.error(e)
  } finally {
    isLoading.value = false
  }
}

onMounted(() => {
  const path = route.params.path as string
  if (path) loadDoc(path)
})

watch(() => route.params.path, (newPath) => {
  if (newPath) loadDoc(newPath as string)
})
</script>

<template>
  <div class="min-h-screen" :class="{ 'cursor-wait': isLoading }">
    <div class="max-w-[800px] mx-auto px-12 py-12">
      <!-- Back link -->
      <button
        class="mb-6 text-sm opacity-60 hover:opacity-100 transition-opacity duration-150 cursor-pointer"
        @click="router.push({ name: 'home' })"
      >
        ← 返回文档列表
      </button>

      <!-- Error -->
      <div v-if="error" class="text-center py-12 opacity-40">
        {{ error }}
      </div>

      <!-- Document -->
      <div v-else-if="doc">
        <h1 class="text-3xl font-bold mb-8">{{ doc.title }}</h1>
        <DocRenderer :content="doc.content" />
      </div>
    </div>
  </div>
</template>

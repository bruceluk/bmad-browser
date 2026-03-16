<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()
const role = route.params.role as string
</script>

<template>
  <div class="min-h-screen">
    <div class="max-w-[800px] mx-auto px-12 py-12">
      <button
        class="mb-6 text-sm opacity-60 hover:opacity-100 transition-opacity duration-150 cursor-pointer"
        @click="router.push({ name: 'home' })"
      >
        ← 返回首页
      </button>
      <h1 class="text-3xl font-bold mb-4">流程图视图</h1>
      <p class="opacity-60">角色：{{ role }}（Story 2.3 将实现完整流程图）</p>
    </div>
  </div>
</template>

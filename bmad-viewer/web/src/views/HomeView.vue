<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { fetchRoles } from '@/api/client'
import type { RoleFlow } from '@/types'
import RoleCard from '@/components/RoleCard.vue'

const router = useRouter()
const roles = ref<RoleFlow[]>([])
const isLoading = ref(true)

const roleIcons: Record<string, string> = {
  developer: '💻',
  pm: '📋',
  qa: '🧪',
}

const roleDescriptions: Record<string, string> = {
  developer: '从需求到代码：了解故事怎么来的、怎么开发、怎么审查',
  pm: '从想法到需求：了解产品简报、PRD、史诗故事怎么产出',
  qa: '从开发到验收：了解自动化测试、QA 流程怎么跑',
}

onMounted(async () => {
  try {
    roles.value = await fetchRoles()
  } catch (e) {
    console.error(e)
  } finally {
    isLoading.value = false
  }
})

function calcTotalDuration(role: RoleFlow): string {
  let totalMin = 0
  for (const step of role.steps) {
    const match = step.duration.match(/(\d+)/)
    if (match) totalMin += parseInt(match[1])
  }
  if (totalMin >= 60) {
    const hours = Math.floor(totalMin / 60)
    const mins = totalMin % 60
    return mins > 0 ? `约${hours}小时${mins}分钟` : `约${hours}小时`
  }
  return `约${totalMin}分钟`
}

function goToFlow(role: string) {
  router.push({ name: 'flow', params: { role } })
}
</script>

<template>
  <div class="min-h-screen flex flex-col items-center justify-center" :class="{ 'cursor-wait': isLoading }">
    <div class="text-center mb-12">
      <h1 class="text-4xl font-bold mb-3">BMAD Viewer</h1>
      <p class="text-lg opacity-60">这是一个用 BMAD 方法从零构建的真实项目。<br>选择你的角色，了解 AI 如何参与你的工作流程。</p>
    </div>

    <div class="flex gap-6 flex-wrap justify-center">
      <RoleCard
        v-for="role in roles"
        :key="role.role"
        :role="role.role"
        :role-color="role.roleColor"
        :icon="roleIcons[role.role] || '📄'"
        :label="role.label"
        :description="roleDescriptions[role.role] || role.label"
        :step-count="role.steps.length"
        :total-duration="calcTotalDuration(role)"
        @click="goToFlow(role.role)"
      />
    </div>
  </div>
</template>

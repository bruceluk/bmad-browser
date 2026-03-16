<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { fetchRoles } from '@/api/client'
import type { RoleFlow } from '@/types'
import RoleTab from '@/components/RoleTab.vue'
import FlowNode from '@/components/FlowNode.vue'
import FlowArrow from '@/components/FlowArrow.vue'

const route = useRoute()
const router = useRouter()
const roles = ref<RoleFlow[]>([])
const currentRole = ref(route.params.role as string)
const isLoading = ref(true)

const currentRoleFlow = computed(() =>
  roles.value.find(r => r.role === currentRole.value)
)

onMounted(async () => {
  try {
    roles.value = await fetchRoles()
  } catch (e) {
    console.error(e)
  } finally {
    isLoading.value = false
  }
})

watch(() => route.params.role, (newRole) => {
  if (newRole) currentRole.value = newRole as string
})

function switchRole(role: string) {
  currentRole.value = role
  router.replace({ name: 'flow', params: { role } })
}

function goHome() {
  router.push({ name: 'home' })
}
</script>

<template>
  <div class="min-h-screen flex flex-col" :class="{ 'cursor-wait': isLoading }">
    <!-- Header + Tabs -->
    <div class="px-8 py-4" style="border-bottom: 1px solid var(--border);">
      <div class="flex items-center justify-between mb-3">
        <button
          class="text-lg font-bold cursor-pointer hover:opacity-80 transition-opacity duration-150"
          @click="goHome"
        >
          BMAD Viewer
        </button>
      </div>
      <RoleTab
        v-if="roles.length > 0"
        :roles="roles"
        :current-role="currentRole"
        @select="switchRole"
      />
    </div>

    <!-- Flow Chart Area -->
    <div
      v-if="currentRoleFlow"
      class="px-8 py-6 flex items-center justify-center overflow-x-auto"
      style="border-bottom: 1px solid var(--border);"
    >
      <div class="flex items-center gap-0">
        <!-- Upstream nodes -->
        <template v-if="currentRoleFlow.upstream.length > 0">
          <span class="text-[11px] mr-2 writing-vertical" style="color: var(--text-secondary); writing-mode: vertical-lr; letter-spacing: 2px;">上游</span>
          <template v-for="(step, i) in currentRoleFlow.upstream" :key="'up-' + step.code">
            <FlowNode
              :step="step"
              :role-color="currentRoleFlow.roleColor"
              :is-upstream="true"
            />
            <FlowArrow v-if="i < currentRoleFlow.upstream.length - 1" />
          </template>
          <FlowArrow />
        </template>

        <!-- Core steps -->
        <template v-for="(step, i) in currentRoleFlow.steps" :key="step.code">
          <FlowNode
            :step="step"
            :role-color="currentRoleFlow.roleColor"
          />
          <FlowArrow
            v-if="i < currentRoleFlow.steps.length - 1"
            :color="currentRoleFlow.roleColor"
          />
        </template>

        <!-- Downstream nodes -->
        <template v-if="currentRoleFlow.downstream.length > 0">
          <FlowArrow />
          <template v-for="(step, i) in currentRoleFlow.downstream" :key="'down-' + step.code">
            <FlowNode
              :step="step"
              :role-color="currentRoleFlow.roleColor"
              :is-downstream="true"
            />
            <FlowArrow v-if="i < currentRoleFlow.downstream.length - 1" />
          </template>
          <span class="text-[11px] ml-2" style="color: var(--text-secondary); writing-mode: vertical-lr; letter-spacing: 2px;">下游</span>
        </template>
      </div>
    </div>

    <!-- Document Area (placeholder for Epic 3) -->
    <div class="flex-1 flex items-center justify-center">
      <p class="opacity-40">点击上方流程节点查看详情</p>
    </div>
  </div>
</template>

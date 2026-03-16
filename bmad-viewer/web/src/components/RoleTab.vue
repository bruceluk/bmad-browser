<script setup lang="ts">
import type { RoleFlow } from '@/types'

defineProps<{
  roles: RoleFlow[]
  currentRole: string
}>()

defineEmits<{
  select: [role: string]
}>()

const roleIcons: Record<string, string> = {
  developer: '💻',
  pm: '📋',
  qa: '🧪',
}
</script>

<template>
  <div class="flex gap-2 justify-center">
    <button
      v-for="role in roles"
      :key="role.role"
      class="px-5 py-2 rounded-full text-sm cursor-pointer transition-all duration-150"
      :style="currentRole === role.role
        ? { backgroundColor: role.roleColor + '26', borderColor: role.roleColor, color: role.roleColor, border: '1px solid' }
        : { border: '1px solid var(--border)', color: 'var(--text-secondary)' }"
      @mouseenter="(e) => {
        if (currentRole !== role.role) {
          (e.currentTarget as HTMLElement).style.borderColor = role.roleColor
        }
      }"
      @mouseleave="(e) => {
        if (currentRole !== role.role) {
          (e.currentTarget as HTMLElement).style.borderColor = 'var(--border)'
        }
      }"
      @click="$emit('select', role.role)"
    >
      {{ roleIcons[role.role] || '📄' }} {{ role.label }}
    </button>
  </div>
</template>

<script setup lang="ts">
import type { WorkflowStep } from '@/types'

const props = defineProps<{
  step: WorkflowStep
  roleColor: string
  isUpstream?: boolean
  isDownstream?: boolean
  isActive?: boolean
}>()

defineEmits<{
  click: []
}>()

const isContext = props.isUpstream || props.isDownstream
</script>

<template>
  <div
    class="min-w-[140px] flex flex-col items-center gap-2 px-5 py-4 rounded-xl cursor-pointer transition-all duration-150"
    :class="{ 'opacity-40': isContext }"
    :style="{
      border: isContext ? '2px dashed var(--border)' : isActive ? `2px solid ${roleColor}` : '2px solid var(--border)',
      backgroundColor: isActive ? roleColor + '14' : 'var(--surface)',
    }"
    @mouseenter="(e) => {
      if (!isContext) {
        const el = e.currentTarget as HTMLElement
        el.style.borderColor = roleColor
        el.style.transform = 'translateY(-2px)'
      }
    }"
    @mouseleave="(e) => {
      if (!isContext) {
        const el = e.currentTarget as HTMLElement
        el.style.borderColor = isActive ? roleColor : 'var(--border)'
        el.style.transform = ''
      }
    }"
    @click="$emit('click')"
  >
    <span class="text-xl">{{ step.agentIcon }}</span>
    <span class="text-[13px] font-semibold text-center">{{ step.name }}</span>
    <span
      class="text-[10px] text-center"
      style="font-family: 'JetBrains Mono', 'Fira Code', monospace; color: var(--text-secondary);"
    >
      {{ step.command }}
    </span>
    <span
      class="text-[11px] px-2 py-0.5 rounded-full"
      :style="{ backgroundColor: 'rgba(255,167,38,0.1)', color: 'var(--warning)' }"
    >
      {{ step.duration }}
    </span>
  </div>
</template>

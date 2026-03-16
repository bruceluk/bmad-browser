<script setup lang="ts">
import type { WorkflowStep } from '@/types'

defineProps<{
  step: WorkflowStep
  roleColor: string
}>()
</script>

<template>
  <div class="flex flex-wrap gap-3 mb-6">
    <!-- Agent badge -->
    <span
      class="text-sm px-3 py-1 rounded-md"
      :style="{ backgroundColor: roleColor + '26', color: roleColor }"
    >
      {{ step.agentIcon }} {{ step.agentName }}
    </span>

    <!-- Command -->
    <code
      class="text-sm px-3 py-1 rounded-md"
      style="background-color: var(--surface); color: var(--dev); font-family: 'JetBrains Mono', 'Fira Code', monospace;"
    >
      {{ step.command }}
    </code>

    <!-- Duration -->
    <span
      class="text-sm px-3 py-1 rounded-md"
      style="background-color: rgba(255,167,38,0.1); color: var(--warning);"
    >
      ⏱ {{ step.duration }}
    </span>

    <!-- Required badge -->
    <span
      v-if="step.required"
      class="text-sm px-3 py-1 rounded-md"
      style="background-color: rgba(102,187,106,0.1); color: var(--success);"
    >
      必需
    </span>
  </div>
</template>

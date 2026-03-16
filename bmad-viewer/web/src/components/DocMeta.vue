<script setup lang="ts">
import type { WorkflowStep } from '@/types'

export interface MatchedOutput {
  label: string
  docPath: string | null
}

defineProps<{
  step: WorkflowStep
  roleColor: string
  matchedOutputs?: MatchedOutput[]
}>()

defineEmits<{
  navigate: [path: string]
}>()
</script>

<template>
  <div class="mb-6">
    <!-- Badges row -->
    <div class="flex flex-wrap gap-3 mb-3">
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

    <!-- Output files -->
    <div v-if="matchedOutputs && matchedOutputs.length > 0" class="flex flex-wrap items-center gap-2">
      <span class="text-xs opacity-50">产出文件：</span>
      <template v-for="(output, i) in matchedOutputs" :key="i">
        <span
          v-if="output.docPath"
          class="text-xs px-2 py-0.5 rounded cursor-pointer transition-opacity duration-150 hover:opacity-100 opacity-70"
          style="background-color: var(--surface); color: var(--link);"
          @click="$emit('navigate', output.docPath!)"
        >
          📄 {{ output.label }}
        </span>
        <span
          v-else
          class="text-xs px-2 py-0.5 rounded opacity-40"
          style="background-color: var(--surface);"
        >
          {{ output.label }}
        </span>
      </template>
    </div>
  </div>
</template>
